package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	rrdis "github.com/go-redis/redis/v8"
	"github.com/liujunren93/share_utils/db/redis"
	"github.com/liujunren93/share_utils/helper"
	"github.com/liujunren93/share_utils/log"
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
		app := helper.SubstrRight(v, r.GetKey(""))
		tmp, err := r.getRouter(ctx, v)
		if err != nil {
			log.Logger.Error(err)
			continue
		}
		resData[app] = tmp

	}
	return resData
}

func (r *RouteCenter) getRouter(ctx context.Context, key string) (routers map[string]*router.Router, err error) {
	res := r.client.HGet(ctx, key, "router")
	if res.Err() != nil && res.Err() != rrdis.Nil {
		return nil, res.Err()
	}

	err = json.Unmarshal([]byte(res.Val()), &routers)
	return
}

func (r *RouteCenter) GetRouter(ctx context.Context, app string) (routers map[string]*router.Router, err error) {
	return r.getRouter(ctx, r.GetKey(app))
}

func (r *RouteCenter) Registry(ctx context.Context, app string, router map[string]*router.Router) error {

	data, err := json.Marshal(router)
	if err != nil {
		return err
	}
	lua := `
		if (redis.call('EXISTS', KEYS[1]) ~= 1) then
				 		redis.call('HSET',KEYS[1],KEYS[2],ARGV[1]) 
						redis.call('PUBLISH', KEYS[3],ARGV[2])
					end
			redis.call('HINCRBY',KEYS[1],'life',1)
		
		`
	fmt.Println(r.GetSubChannelReg())
	res := r.client.Eval(ctx, lua, []string{r.GetKey(app), "router", r.GetSubChannelReg()}, string(data), app)
	if res.Err() == rrdis.Nil {
		return nil
	}
	return res.Err()

}

func (r *RouteCenter) Lease(ctx context.Context, app string) error {
	var err error
	go func() {
		for {
			res := r.client.Expire(context.TODO(), r.GetKey(app), time.Second*60)
			if res.Err() != nil && res.Err() != rrdis.Nil {
				err = res.Err()
				break
			}
			time.Sleep(time.Second * 50)
		}
	}()
	return err
}

func (r *RouteCenter) DelRouter(ctx context.Context, app string) error {

	lua := `
			redis.call('HINCRBY',KEYS[1],'life',-1)
			local cnt= tonumber(redis.call('HGET',KEYS[1],'life'))
			if (cnt<=0) then
				redis.call('DEL',KEYS[1])
			end
			if (cnt==0)then
			redis.call('PUBLISH', KEYS[2],ARGV[1])
			end
		`
	res := r.client.Eval(ctx, lua, []string{r.GetKey(app), r.GetSubChannelDel()}, app)
	if res.Err() == rrdis.Nil {
		return nil
	}
	return res.Err()
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
