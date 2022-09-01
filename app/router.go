package app

import (
	"context"
	"time"

	shareRouter "github.com/liujunren93/share_utils/common/gin/router"
	"github.com/liujunren93/share_utils/db/redis"
	"github.com/liujunren93/share_utils/pkg/routerCenter"
	routerRedis "github.com/liujunren93/share_utils/pkg/routerCenter/redis"
)

func (app *App) getRouterCenter() routerCenter.RouterCenter {
	var rc routerCenter.RouterCenter
	routerConfig := app.defaultConf.GetRouterCenter()
	if !routerConfig.Enable {
		return nil
	}
	if routerConfig != nil {
		panic("You must set router center config")
	}

	switch routerConfig.Type {
	case 1:
		cli, err := redis.NewClient(routerConfig.RedisConf)
		if err != nil {
			panic("initRouter.redis.NewClient:" + err.Error())
		}
		rc = routerRedis.NewRouteCenter(cli)

	}
	return rc
}

func (app *App) initRouter() {
	rc := app.getRouterCenter()
	if rc == nil {
		return
	}
	ctx, _ := context.WithTimeout(app.ctx, time.Second*3)
	routerMap, err := rc.GetRouter(ctx)
	if err != nil {
		panic(err)
	}
	tree := shareRouter.NewTree("/", "")

	for k, r := range routerMap {
		tree.Add(k, r.Method, r.GrpcMenthod)
	}

	app.appRouter = tree

	rc.Watch(app.ctx, func(router map[string]routerCenter.Router, error error) {

	})

}
