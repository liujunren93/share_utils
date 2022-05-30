package redis

import re "github.com/go-redis/redis/v8"

func NewRedis(conf *re.Options) *re.Client {

	return re.NewClient(conf)
}
