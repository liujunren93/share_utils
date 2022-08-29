package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	router "github.com/liujunren93/share_utils/pkg/routerCenter"
)

type Redis struct {
	router.RouterCentry
	client *redis.Client
}

func (r *Redis) GetSubChannel(key string) string {
	return r.GetKey(key) + "/" + "subscribe"
}

func (r *Redis) GetRouter(ctx context.Context, key string) (routers map[string]router.Router, err error) {
	res := r.client.Get(ctx, r.GetKey(key))
	if res.Err() != nil && res.Err() != redis.Nil {
		return nil, res.Err()
	}
	data, err := res.Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &routers)
	return
}

func (r *Redis) Registry(ctx context.Context, key string, router map[string]router.Router) error {
	data, err := json.Marshal(router)
	if err != nil {
		return err
	}
	err = r.client.Set(ctx, r.GetKey(key), string(data), time.Minute*30).Err()
	if err != nil {
		return err
	}
	err = r.client.Publish(ctx, r.GetSubChannel(key), "1").Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) DelRouter(ctx context.Context, key string) error {
	err := r.client.Del(ctx, r.GetKey(key)).Err()
	if err != nil {
		return err
	}
	err = r.client.Publish(ctx, r.GetSubChannel(key), "-1").Err()
	if err != nil {
		return err
	}
	return nil

}

func (r *Redis) Watch(ctx context.Context, key string, callback func(router map[string]router.Router, err error)) {
	pub := r.client.Subscribe(ctx, key)
	for {
		select {
		case msg := <-pub.Channel():

			if msg.Payload == "1" { //add
				tctx, _ := context.WithTimeout(ctx, time.Second*3)
				callback(r.GetRouter(tctx, key))
			} else { // del
				callback(nil, nil)
			}
		case <-ctx.Done():
			return
		}
	}
}

func NewRedis(cli *redis.Client) *Redis {
	return &Redis{client: cli}
}
