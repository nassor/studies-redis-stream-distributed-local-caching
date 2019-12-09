package main

import (
	"github.com/mediocregopher/radix/v3"
	"github.com/rs/zerolog/log"
)

func main() {
	pool, err := radix.NewPool("tcp", "127.0.0.1:6379", 10)
	if err != nil {
		log.Panic().Err(err).Msg("initializing redis pool")
	}

}
