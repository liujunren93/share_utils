package databases

import "github.com/go-redis/redis"

func NewRedis(conf *redis.Options) *redis.Client {

	return redis.NewClient(conf)

}
