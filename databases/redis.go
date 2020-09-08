package databases

import "github.com/go-redis/redis"

var RedisDB *redis.Client

type RedisConf struct {
	Network  string `json:"network"`
	Host     string `json:"host"`
	DB       int    `json:"db"`
	Password string `json:"password"`
}

func NewRedis(conf *redis.Options) *redis.Client {
	if RedisDB == nil {

		RedisDB = redis.NewClient(conf)
	}
	return RedisDB
}
