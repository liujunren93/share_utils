package databases

import "github.com/go-redis/redis"

var RedisDB *redis.Client

func NewRedis(host, password string, db int) *redis.Client {
	if RedisDB == nil {
		options := redis.Options{
			Network:  "tcp",
			Addr:     host,
			DB:       5,
			Password: password,
		}
		RedisDB = redis.NewClient(&options)
	}
	return RedisDB
}
