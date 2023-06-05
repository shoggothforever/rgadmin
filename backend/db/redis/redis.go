package dbredis

import (
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	CurrentNumKey string = "Count:Now-%d-%d"
	AllNumKey     string = "Count:All"
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		DB:           6,
		ReadTimeout:  time.Millisecond * time.Duration(500),
		WriteTimeout: time.Millisecond * time.Duration(500),
		IdleTimeout:  time.Second * time.Duration(60),
		PoolSize:     64,
		MinIdleConns: 16,
		PoolFIFO:     true,
		//MaxRetries:   3,
	})
}
