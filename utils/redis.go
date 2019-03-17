package utils

import "github.com/go-redis/redis"

func GetRedis() *redis.Client {

	client := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			PoolSize: 10,
		})

	return client
}
