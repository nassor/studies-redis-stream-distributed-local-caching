package device

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/mediocregopher/radix/v3"
)

const (
	stream = "cache:device"
	key    = "k"
	value  = "v"
)

type ristrettoCache interface {
	Set(interface{}, interface{}, int64) bool
	Get(interface{}) (interface{}, bool)
}

type redisStreamReader interface {
	Next() (string, []radix.StreamEntry, bool)
	Err() error
}

type redisClient interface {
	Do(a radix.Action) error
}

type Cache struct {
	ristrettoCache
	redisClient  redisClient
	updateStream redisStreamReader
	updateFreq   time.Duration
	poolSize     int
}

func NewCache(c ristrettoCache, rc radix.Client, groupID string, poolSize int, updateFreq time.Duration) *Cache {
	rsr := radix.NewStreamReader(rc, radix.StreamReaderOpts{
		Group:    groupID,
		Consumer: uuid.New().String(),
		NoBlock:  true,
		NoAck:    true,
		Count:    poolSize,
		Streams:  map[string]*radix.StreamEntryID{stream: nil},
	})
	return &Cache{
		ristrettoCache: c,
		redisClient:    rc,
		updateStream:   rsr,
		updateFreq:     updateFreq,
		poolSize:       poolSize,
	}
}

func (c *Cache) Start(ctx context.Context) {
	freq := time.NewTicker(c.updateFreq)
	for {
		select {
		case <-ctx.Done():
			freq.Stop()
			return
		case <-freq.C:
			if err := c.process(); err != nil {
				// TODO add stats
			}
		}
	}
}

func (c *Cache) process() error {
	k, v, err := c.receive()
	if err != nil {
		return err
	}
	for i := 0; i < len(k); i++ {
		if ok := c.Set(k[i], v[i], 1); !ok {
			// TODO add an metric
		}
	}
	return nil
}

func (c *Cache) receive() (keys, values []string, err error) {
	stream, entries, ok := c.updateStream.Next()
	if stream == "" {
		return
	}
	if !ok {
		err = c.updateStream.Err()
		return
	}
	size := len(entries)
	keys = make([]string, size)
	values = make([]string, size)
	for i := range entries {
		keys[i] = entries[i].Fields[key]
		values[i] = entries[i].Fields[value]
	}
	return
}
