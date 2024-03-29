package app

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"

	codesJson "github.com/liujunren93/share/codes/json"
	shareRouter "github.com/liujunren93/share_utils/common/gin/router"
	"github.com/liujunren93/share_utils/db/redis"
	shErr "github.com/liujunren93/share_utils/errors"
	"github.com/liujunren93/share_utils/helper"
	"github.com/liujunren93/share_utils/log"
	"github.com/liujunren93/share_utils/netHelper"
	"github.com/liujunren93/share_utils/pkg/routerCenter"
	routerRedis "github.com/liujunren93/share_utils/pkg/routerCenter/redis"
)

type router struct {
	routes      map[string]*shareRouter.Node
	middlewares map[string]middlewareFunc
}

type middlewareFunc func(*gin.Context) error
type MiddlewareItem struct {
	Name           string
	MiddlewareFunc middlewareFunc
}

func (app *App) getRouterCenter() routerCenter.RouterCenter {
	var rc routerCenter.RouterCenter
	routerConfig := app.cloudConfig.GetRouterCenter()
	if routerConfig == nil {
		return nil
		// panic("You must set router center config")
	}
	if !routerConfig.Enable {
		return nil
	}

	switch routerConfig.Type {
	case 1:
		cli, err := redis.NewClient(routerConfig.RedisConf)
		if err != nil {
			panic("initRouter.redis.NewClient:" + err.Error())
		}
		rc = routerRedis.NewRouteCenter(cli, routerConfig.RouterPrefix, app.localConf.GetLocalBase().Namespace)

	}
	return rc
}

func (app *App) delRouter() {

	log.Logger.Info("Application logout from router center")
	err := app.rc.DelRouter(context.Background(), app.GetAppName())
	if err != nil {
		log.Logger.Error("Application logout from router center ", err)
	}

}

func (app *App) initRouter() {
	rc := app.getRouterCenter()
	if rc == nil {
		return
	}
	app.rc = rc
	ctx, _ := context.WithTimeout(app.ctx, time.Second*3)
	routerMap := rc.GetAllRouter(ctx)
	if app.appRouter.routes == nil {
		app.appRouter.routes = make(map[string]*shareRouter.Node)
	}
	for appName, routers := range routerMap {

		tree := routerMap2Tree(routers)
		app.appRouter.routes[appName] = tree
	}
	var mu = sync.Mutex{}
	rc.Watch(app.ctx, func(appName string, routers map[string]*routerCenter.Router, err error) {
		fmt.Println(111)
		mu.Lock()
		defer mu.Unlock()
		if len(routers) == 0 {
			delete(app.appRouter.routes, appName)
		} else {
			tree := routerMap2Tree(routers)
			app.appRouter.routes[appName] = tree
		}

	})

}

func routerMap2Tree(routerMap map[string]*routerCenter.Router) *shareRouter.Node {
	tree := shareRouter.NewTree("/", "")
	for apipath, router := range routerMap {
		index := strings.Index(apipath, ":")
		method := apipath[:index]
		apipath := apipath[index+1:]
		tree.Add(apipath, method, router.GrpcMenthod, router.MiddlewaresWhitelist, router.ReqParams)
	}
	return tree
}

var validate *validator.Validate

func (a *App) autoRoute(r shareRouter.Router) error {

	a.initRouter()
	validate = validator.New()
	log.Logger.Debug("AutoRoute")

	r.NoRoute(func(ctx *gin.Context) {
		appName, reqPath, method, reqData, err := ParesRequest(ctx, a.localConf.GetLocalBase().ApiPrefix)
		if err != nil {

			netHelper.Response(ctx, shErr.NewBadRequest(nil), nil, nil)
			return
		}
		appPrefix := a.cloudConfig.GetRouterCenter().AppPrefix
		if appPrefix == "" {
			appPrefix = "share_app"
		}
		appName = appPrefix + "_" + appName
		// 	isRetry := false
		// retry:
		p, ok := a.appRouter.routes[appName]
		if ok {
			node, param := p.Find(reqPath, method)
			if node == nil {
				netHelper.Response(ctx, shErr.NewStatusNotFound(""), nil, nil)
				return
			}

			for _, w := range node.MiddlewaresWhitelist {
				if mid, ok := a.appRouter.middlewares[w]; !ok {
					err = mid(ctx)
					if err != nil {
						return
					}
				}
			}
			var req = make(map[string]interface{})

			if len(reqData) > 0 {
				err = json.Unmarshal(reqData, &req)
				if err != nil {
					netHelper.Response(ctx, shErr.NewBadRequest(err), err, nil)
					return
				}
			}
			// param 参数
			if len(param.Key) > 0 {

				req["pk"] = param.Value
				reqData, err = json.Marshal(req)
				if err != nil {
					log.Logger.Error("AutoRoute.param.Marshal", err)
					netHelper.Response(ctx, shErr.NewBadRequest(err), err, nil)
					return
				}
				ctx.Params = append(ctx.Params, param)
			}

			checkRes := validate.ValidateMap(req, node.ReqParams)

			if len(checkRes) != 0 {
				errMsg := bytes.Buffer{}
				for k, v := range checkRes {
					errMsg.WriteString(fmt.Sprintf("%s:%v;", k, v))
				}

				netHelper.Response(ctx, shErr.NewBadRequest(errors.New("bad request:"+errMsg.String())), nil, nil)
				// log.Logger.Error("noRoute.Prepare", err)
				return
			}

			cc, err := a.shareGrpcClient.Client(appName)
			if err != nil {
				log.Logger.Error("noRoute.shareGrpcClient.Client", err)
				return
			}
			var res interface{}
			err = a.shareGrpcClient.Invoke(ctx, node.GrpcPath, reqData, &res, cc, grpc.CallContentSubtype(codesJson.Name))
			log.Logger.Debug(res)
			if err != nil {
				log.Logger.Error("grpc.Invoke", err)
			}
			if re, ok := res.(netHelper.Responser); ok {
				if re.GetCode() != 200 {
					log.Logger.Info(re)
				}
			}

			netHelper.ResponseJson(ctx, res, err, nil)
		} else {
			netHelper.Response(ctx, shErr.NewStatusNotFound(""), nil, nil)
			return
		}

	})
	return nil
}

func (a *App) RegistryRouter(rcMap map[string]*routerCenter.Router) {
	if len(rcMap) == 0 {
		return
	}
	if a.rc == nil {
		rc := a.getRouterCenter()
		if rc == nil {
			panic("init routerCenter failed")
		}
		a.rc = rc
	}
	ctx, _ := context.WithTimeout(a.ctx, time.Second*10)
	// appnames := strings.Split(appName, "_")
	// appName = appnames[len(appnames)-1]

	err := a.rc.Registry(ctx, a.GetAppName(), rcMap)
	if err != nil {
		panic("registry router failed:" + err.Error())
	}
	err = a.rc.Lease(ctx, a.GetAppName())
	if err != nil {
		panic("registry router.Lease failed:" + err.Error())
	}
	a.RegistryStopFunc(a.delRouter)
}
func ParesRequest(ctx *gin.Context, urlPrefix string) (appName, reqPath, method string, body []byte, err error) {
	method = ctx.Request.Method

	reqPath = strings.Trim(strings.TrimLeft(path.Clean(ctx.Request.URL.Path), urlPrefix), "/")
	appName = helper.SubstrLeft(reqPath, "/")
	reqPath = helper.SubstrRight(reqPath, "/")
	if method == "GET" {
		var req = make(map[string]interface{}, len(ctx.Request.URL.Query()))
		if len(ctx.Request.URL.Query()) != 0 {
			for k, v := range ctx.Request.URL.Query() {
				if strings.LastIndex(k, "_str") >= 0 && strings.LastIndex(k, "_str") == len(k)-4 {
					// if strings.LastIndex(k, "_str") == len(k)-4 {
					fmt.Println(k, v)
					req[k[:len(k)-4]] = v[0]
				} else if nv, err := strconv.ParseFloat(v[0], 64); err == nil {
					req[k] = nv
				}
			}
		}
		if _, ok := ctx.GetQuery("sort_order_str"); !ok {
			req["sort_order"] = ""
		}

		body, err = json.Marshal(req)
		if err != nil {
			return
		}
	} else {
		body, err = io.ReadAll(ctx.Request.Body)
		if err != nil {
			return
		}
		defer ctx.Request.Body.Close()
	}

	return

}
