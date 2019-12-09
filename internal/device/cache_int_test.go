package device

import (
	"context"
	"testing"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/google/uuid"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/protobuf/proto"
	"github.com/mediocregopher/radix/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nassor/studies-redis-stream-distributed-local-caching/api"
)

type fakeDevice struct {
	ID   string `faker:"uuid_hyphenated"`
	Data string `faker:"len=2150"`
}

func newFakeDevice(t *testing.T) api.Device {
	fd := fakeDevice{}
	err := faker.FakeData(&fd)
	require.NoError(t, err)
	d := api.Device{
		Id:   fd.ID,
		Data: fd.Data,
	}
	return d
}

func newPool(t *testing.T) (*radix.Pool, string) {
	pool, err := radix.NewPool("tcp", "127.0.0.1:6379", 10)
	require.NoError(t, err)
	err = pool.Do(radix.Cmd(nil, "FLUSHALL"))
	require.NoError(t, err)
	groupID := uuid.New().String()
	err = pool.Do(radix.Cmd(nil, "XGROUP", "CREATE", stream, groupID, "$", "MKSTREAM"))
	require.NoError(t, err)
	return pool, groupID
}

func generateEvent(t *testing.T, pool *radix.Pool) api.Device {
	d := newFakeDevice(t)
	b, err := proto.Marshal(&d)
	require.NoError(t, err)
	err = pool.Do(radix.Cmd(nil, "XADD", stream, "MAXLEN", "~", "5000", "*",
		key, string(d.Id), value, string(b)))
	// err = pool.Do(radix.Cmd(nil, "XADD", stream, "*",
	// 	key, string(d.Id), value, string(b)))
	require.NoError(t, err)
	return d
}

func newInternalCache(t *testing.T) *ristretto.Cache {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1 << 10, // number of keys to track frequency of (10M).
		MaxCost:     1 << 20, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	require.NoError(t, err)
	return cache
}

func TestIntCache_Start(t *testing.T) {
	pool, groupID := newPool(t)
	device := generateEvent(t, pool)
	ic := newInternalCache(t)
	c := NewCache(ic, pool, groupID, 10, 900*time.Millisecond)
	ctx := context.Background()
	go c.Start(ctx)
	time.Sleep(time.Second)

	data, found := c.Get(device.Id)
	assert.True(t, found)
	d := api.Device{}
	err := proto.Unmarshal([]byte(data.(string)), &d)
	assert.NoError(t, err)
	assert.Equal(t, device.Data, d.Data)

}
