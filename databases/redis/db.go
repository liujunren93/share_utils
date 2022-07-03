package redis

import (
	"sync"

	re "github.com/go-redis/redis/v8"
)

var Redis *re.Client
var mu sync.Mutex
var redisVersion int64

func InitRedis(conf *re.Options) error {
	return newRedis(conf)
}

func newRedis(conf *re.Options) (err error) {
	tmpVersion := redisVersion
	mu.Lock()
	defer mu.Unlock()
	if tmpVersion == redisVersion {
		Redis, err = NewRedis(conf)
	}
	return
}
func UpdateRedis(conf *re.Options) error {
	return newRedis(conf)
}
