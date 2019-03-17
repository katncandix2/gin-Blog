package utils

import "github.com/go-redis/redis"

var r *redis.Client

func init() {
	r = GetRedis()
}

func GetUUID() int64 {

	if r == nil {
		return -1
	}

	defer r.Close()
	res := r.Incr("article")
	return res.Val()
}
