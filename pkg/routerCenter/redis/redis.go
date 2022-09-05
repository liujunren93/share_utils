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

func NewRouteCenter(cli *redis.Client, prefix, namespace string) *RouteCenter {
	return &RouteCenter{
		RouterCentry: router.RouterCentry{
			Namespace: namespace,
			Prefix:    prefix,
		},
		client: cli,
	}
}

func (r *RouteCenter) GetSubChannelReg() string {
	return r.GetKey("") + "subscribeReg"
}

func (r *RouteCenter) GetSubChannelDel() string {
	return r.GetKey("") + "subscribeDel"
}
func (r *RouteCenter) GetAllRouter(ctx context.Context) map[string]map[string]*router.Router {
	var resData = make(map[string]map[string]*router.Router)
	res := r.client.Keys(ctx, r.GetKeys(""))
	keys := res.Val()
	// var routerDatas = make(map[string]map[string]router.Router)
	for _, v := range keys {
		var data map[string]*router.Router
		re := r.client.Get(ctx, v)
		if re.Err() != nil {
			continue
		}
		json.Unmarshal([]byte(re.Val()), &data)
		app := v[len(r.GetKey("")):]
		resData[app] = data

	}
	return resData
}

func (r *RouteCenter) GetRouter(ctx context.Context, app string) (routers map[string]*router.Router, err error) {
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

func (r *RouteCenter) Registry(ctx context.Context, app string, router map[string]*router.Router) error {
	data, err := json.Marshal(router)
	if err != nil {
		return err
	}
	err = r.client.Set(ctx, r.GetKey(app), string(data), time.Minute*30).Err()
	if err != nil {
		return err
	}
	err = r.client.Publish(ctx, r.GetSubChannelReg(), app).Err()
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
	err = r.client.Publish(ctx, r.GetSubChannelDel(), app).Err()
	if err != nil {
		return err
	}
	return nil

}

func (r *RouteCenter) Watch(ctx context.Context, callback func(app string, router map[string]*router.Router, err error)) {
	go func() {
		pub := r.client.Subscribe(ctx, r.GetSubChannelReg())
		for {
			select {
			case msg := <-pub.Channel():

				tctx, _ := context.WithTimeout(ctx, time.Second*3)
				data, err := r.GetRouter(tctx, msg.Payload)
				callback(msg.Payload, data, err)

			case <-ctx.Done():
				return
			}
		}

	}()
	go func() {
		pub := r.client.Subscribe(ctx, r.GetSubChannelDel())
		for {
			for {
				select {
				case msg := <-pub.Channel():

					callback(msg.Payload, nil, nil)

				case <-ctx.Done():
					return
				}
			}

		}

	}()

}

// func (r *RouteCenter) Watch1(ctx context.Context, callback func(app string, router map[string]router.Router, err error)) {
// 	pub := r.client.Subscribe(ctx, r.GetSubChannelReg())
// 	pubDel := r.client.Subscribe(ctx, r.GetSubChannelDel())

// 	for {
// 		select {
// 		case msg := <-pub.Channel():

// 			tctx, _ := context.WithTimeout(ctx, time.Second*3)
// 			data, err := r.GetRouter(tctx, msg.Payload)
// 			callback(msg.Payload, data, err)
// 		case msg := <-pubDel.Channel():
// 			callback(msg.Payload, nil, nil)
// 		case <-ctx.Done():
// 			return
// 		}
// 	}

// go func() {
// 	pub := r.client.Subscribe(ctx, r.GetSubChannelDel())
// 	for {
// 		for {
// 			select {
// 			case msg := <-pub.Channel():

// 				callback(msg.Payload, nil, nil)

// 			case <-ctx.Done():
// 				return
// 			}
// 		}

// 	}

// }()

// }
