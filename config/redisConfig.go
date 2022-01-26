package config

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

var Cache *redis.Client

func RedisInit() {
	Cache = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println("Connected to Redis")
}
