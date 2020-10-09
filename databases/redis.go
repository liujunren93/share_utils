package databases

import "github.com/go-redis/redis"

var RedisDB *redis.Client


func NewRedis(conf *redis.Options) *redis.Client {
	if RedisDB == nil {
		RedisDB = redis.NewClient(conf)
	}
	return RedisDB
}
