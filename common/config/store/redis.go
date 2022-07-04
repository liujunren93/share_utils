package store

import (
	"context"
	"fmt"

	"github.com/liujunren93/share_utils/log"

	"github.com/go-redis/redis/v8"
	"github.com/liujunren93/share_utils/common/config"
)

type Redis struct {
	namespace string
	redis     *redis.Client
}

func NewRedis(r *redis.Client, namespace string) config.Configer {
	return &Redis{namespace, r}
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

func (r *Redis) GetConfig(ctx context.Context, confName, group string, callback config.Callback) error {
	_, key := r.GetKey(confName, group)
	content, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return callback(confName, group, content)

}

func (r *Redis) ListenConfig(ctx context.Context, confName, group string, callback config.Callback) error {
	tpkey, _ := r.GetKey(confName, group)
	pubsub := r.redis.Subscribe(ctx, tpkey)
	ch := pubsub.Channel()
	for {
		select {
		case data := <-ch:
			log.Logger.Info("config update", data.Payload)
			callback(confName, group, data.Payload)

		case <-ctx.Done():
			return nil
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
