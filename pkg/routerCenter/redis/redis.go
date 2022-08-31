package redis

import (
	"context"
	"encoding/json"
	"time"

	re "github.com/go-redis/redis/v8"
	"github.com/liujunren93/share_utils/db/redis"
	router "github.com/liujunren93/share_utils/pkg/routerCenter"
)

type RouteCenter struct {
	router.RouterCentry
	client *redis.Client
}

func NewRouteCenter(cli *redis.Client) *RouteCenter {
	return &RouteCenter{client: cli}
}

func (r *RouteCenter) GetSubChannel(app string) string {
	return r.GetKey(app) + "/" + "subscribe"
}

func (r *RouteCenter) GetRouter(ctx context.Context, app string) (routers map[string]router.Router, err error) {
	res := r.client.Get(ctx, r.GetKey(app))
	if res.Err() != nil && res.Err() != re.Nil {
		return nil, res.Err()
	}
	data, err := res.Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &routers)
	return
}

func (r *RouteCenter) Registry(ctx context.Context, app string, router map[string]router.Router) error {
	data, err := json.Marshal(router)
	if err != nil {
		return err
	}
	err = r.client.Set(ctx, r.GetKey(app), string(data), time.Minute*30).Err()
	if err != nil {
		return err
	}
	err = r.client.Publish(ctx, r.GetSubChannel(app), "1").Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RouteCenter) DelRouter(ctx context.Context, app string) error {
	err := r.client.Del(ctx, r.GetKey(app)).Err()
	if err != nil {
		return err
	}
	err = r.client.Publish(ctx, r.GetSubChannel(app), "-1").Err()
	if err != nil {
		return err
	}
	return nil

}

func (r *RouteCenter) Watch(ctx context.Context, app string, callback func(router map[string]router.Router, err error)) {
	pub := r.client.Subscribe(ctx, app)
	for {
		select {
		case msg := <-pub.Channel():

			if msg.Payload == "1" { //add
				tctx, _ := context.WithTimeout(ctx, time.Second*3)
				callback(r.GetRouter(tctx, app))
			} else { // del
				callback(nil, nil)
			}
		case <-ctx.Done():
			return
		}
	}
}
