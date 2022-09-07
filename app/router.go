package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"path"
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

func (app *App) getRouterCenter() routerCenter.RouterCenter {
	var rc routerCenter.RouterCenter
	routerConfig := app.baseConfig.GetRouterCenter()
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
		rc = routerRedis.NewRouteCenter(cli, "", "")

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
	if app.appRouter == nil {
		app.appRouter = make(map[string]*shareRouter.Node)
	}
	for appName, routers := range routerMap {

		tree := routerMap2Tree(routers)
		app.appRouter[appName] = tree
	}
	var mu = sync.Mutex{}
	rc.Watch(app.ctx, func(appName string, routers map[string]*routerCenter.Router, err error) {
		mu.Lock()
		defer mu.Unlock()
		fmt.Println(appName, routers)
		if len(routers) == 0 {
			delete(app.appRouter, appName)
		} else {
			tree := routerMap2Tree(routers)
			app.appRouter[appName] = tree
		}

	})

}

func routerMap2Tree(routerMap map[string]*routerCenter.Router) *shareRouter.Node {
	tree := shareRouter.NewTree("/", "")
	for apipath, router := range routerMap {
		index := strings.Index(apipath, ":")
		method := apipath[:index]
		apipath := apipath[index+1:]
		tree.Add(apipath, method, router.GrpcMenthod, router.ReqParams)
	}
	return tree
}

var validate *validator.Validate

func (a *App) AutoRoute(r shareRouter.Router) error {
	a.initRouter()
	validate = validator.New()
	log.Logger.Debug("AutoRoute")
	r.NoRoute(func(ctx *gin.Context) {

		appName, reqPath, method := ParesRequest(ctx, a.LocalConf.ApiPrefix)
		fmt.Println("AutoRoute.ParesRequest", appName, method)
		appName = a.baseConfig.GetRouterCenter().AppPrefix + "_" + appName
		// 	isRetry := false
		// retry:
		p, ok := a.appRouter[appName]
		if ok {
			node, param := p.Find(reqPath, method)
			if node == nil {
				netHelper.Response(ctx, shErr.NewStatusNotFound(""), nil, nil)
				return
			}
			if len(param.Key) > 0 {
				ctx.Params = append(ctx.Params, param)
			}
			log.Logger.Debug("AutoRoute.method", p.ReqParams)
			var req map[string]interface{}
			buf, err := io.ReadAll(ctx.Request.Body)
			if err != nil {
				netHelper.Response(ctx, shErr.NewBadRequest(err), err, nil)
				return
			}
			defer ctx.Request.Body.Close()
			if len(buf) > 0 {
				err = json.Unmarshal(buf, &req)
				if err != nil {
					netHelper.Response(ctx, shErr.NewBadRequest(err), err, nil)
					return
				}
			}

			checkRes := validate.ValidateMap(req, p.ReqParams)
			if len(checkRes) != 0 {
				re, _ := json.Marshal(checkRes)
				netHelper.Response(ctx, shErr.NewBadRequest(errors.New("bad request:"+string(re))), nil, nil)
				// log.Logger.Error("noRoute.Prepare", err)
				return
			}
			cc, err := a.shareGrpcClient.Client(appName)
			if err != nil {
				log.Logger.Error("noRoute.shareGrpcClient.Client", err)
				return
			}
			var res interface{}
			err = a.shareGrpcClient.Invoke(ctx, node.GrpcPath, buf, &res, cc, grpc.CallContentSubtype(codesJson.Name))
			netHelper.ResponseJson(ctx, res.(map[string]interface{}), err, nil)
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

	a.rc.Registry(ctx, a.GetAppName(), rcMap)
	a.RegistryStopFunc(a.delRouter)
}
func ParesRequest(ctx *gin.Context, urlPrefix string) (appName, reqPath, method string) {
	ctx.FullPath()

	reqPath = strings.Trim(strings.TrimLeft(path.Clean(ctx.Request.URL.Path), urlPrefix), "/")
	return helper.SubstrLeft(reqPath, "/"), helper.SubstrRight(reqPath, "/"), ctx.Request.Method

}
