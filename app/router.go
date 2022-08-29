package app

import "github.com/liujunren93/share_utils/databases/redis"

type Router struct {
	Method    string                 `json:"method" yaml:"method"`
	ReqParams map[string]interface{} `json:"req_params" yaml:"req_params"`
}

type RouterCenter interface {
	GetRouter() map[string]Router //map[method:path]map[]
	Registry(map[string]Router) error
}

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
