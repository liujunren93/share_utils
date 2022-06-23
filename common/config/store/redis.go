package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	namespace string
	timeout   time.Duration
	redis     *redis.Client
}

func NewRedis(r *redis.Client, timeout time.Duration, namespace string) *Redis {
	return &Redis{namespace, timeout, r}
}

func (r *Redis) PublishConfig(ctx context.Context, configName, group, content string) (bool, error) {
	tpkey, configKey := r.GetKey(configName, group)
	res := r.redis.Set(ctx, configKey, content, 0)
	if res.Err() != nil {
		return false, res.Err()
	}
	err := r.redis.Publish(ctx, tpkey, content).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *Redis) GetConfig(ctx context.Context, configName, group string) (interface{}, error) {
	_, key := r.GetKey(configName, group)
	return r.redis.Get(ctx, key).Result()
}

func (r *Redis) ListenConfig(ctx context.Context, configName, group string, f func(interface{})) {
	tpkey, _ := r.GetKey(configName, group)
	pubsub := r.redis.Subscribe(ctx, tpkey)
	ch := pubsub.Channel()
	ctx, _ = context.WithTimeout(context.Background(), r.timeout)
	for {
		select {
		case data := <-ch:
			f(data.Payload)
		case <-ctx.Done():
			return
		}

	}
}

func (r *Redis) DeleteConfig(ctx context.Context, configName, group string) (bool, error) {
	_, configKey := r.GetKey(configName, group)
	err := r.redis.Del(ctx, configKey).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *Redis) GetKey(configName, group string) (string, string) {
	return fmt.Sprintf("topic/%s/%s/%s", r.namespace, configName, group), fmt.Sprintf("config/%s/%s/%s", r.namespace, configName, group)
}
