package redis

import "github.com/go-redis/redis/v8"

func NewRedis(conf *redis.Options) *redis.Client {

	return redis.NewClient(conf)
}
