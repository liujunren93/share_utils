package app

import "github.com/liujunren93/share_utils/db/redis"

func (app *App) initRouter() {
	if !app.LocalConf.EnableAutoRoute {
		return
	}
	routerConfig := app.defaultConf.GetRouterCenter()
	if routerConfig != nil {
		panic("You must set router center config")
	}
	switch routerConfig.Type {
	case 1:
		redis.NewRedis(routerConfig.RedisConf)
	}

}
