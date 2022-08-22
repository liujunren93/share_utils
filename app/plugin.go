package app

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"plugin"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/liujunren93/share_utils/common/gin/router"
	shareRouter "github.com/liujunren93/share_utils/common/gin/router"
	shErr "github.com/liujunren93/share_utils/errors"
	"github.com/liujunren93/share_utils/helper"
	"github.com/liujunren93/share_utils/netHelper"

	"github.com/liujunren93/share_utils/log"
)

type ShPluginer interface {
	Name() (serverName, pluginName string)
	GetRouter() map[string]string
	Prepare(ctx *gin.Context, method string) (req, res interface{}, err shErr.Error)
}

const PLUGIN_NAME = "ShPlugin"

var autoRouter = shareRouter.NewTree("/", "")

type ShPlugin struct {
	pluginer   ShPluginer
	ServerName string
	PluginName string
	Router     *router.Node
}

type Plugin struct {
	pluginMap map[string]*ShPlugin
}

func (a *App) initPlugins() {
	fmt.Println("initPlugins")
	a.plugin = new(Plugin)
	a.plugin.pluginMap = make(map[string]*ShPlugin)
	PluginPath := a.LocalConf.PluginPath
	files, err := os.ReadDir(PluginPath)
	if err != nil {
		fmt.Println("loadPlugins.ReadDir", err, PluginPath)
		log.Logger.Fatal("loadPlugins.ReadDir", err, PluginPath)
	}

	for _, f := range files {
		log.Logger.Debug("initPlugins.plugin.name", f.Name())
		if filepath.Ext(f.Name()) == ".so" {
			err := a.openPlugin(PluginPath + "/" + f.Name())

			if err != nil {
				log.Logger.Error("loadPlugins.OpenPlugin", PluginPath+"/"+f.Name(), err)
				continue
			}
		}
	}
	// panic(len(a.plugin.pluginMap))
}

func (a *App) AutoRoute(r shareRouter.Router) error {
	log.Logger.Debug("AutoRoute", a.plugin.pluginMap)
	r.NoRoute(func(ctx *gin.Context) {
		plugin, reqPath, method := ParesRequest(ctx, a.LocalConf.ApiPrefix)
		fmt.Println("AutoRoute.ParesRequest", plugin, method)

		// 	isRetry := false
		// retry:
		p, ok := a.plugin.pluginMap[plugin]
		if ok {
			node, param := p.Router.Find(reqPath, method)
			if node == nil {
				fmt.Println(node)
				netHelper.Response(ctx, shErr.NewStatusNotFound(""), nil, nil)
				return
			}
			if len(param.Key) > 0 {
				ctx.Params = append(ctx.Params, param)
			}
			req, res, err := p.Prepare(ctx, node.GrpcPath)
			log.Logger.Debug("AutoRoute.method", req, res, err)
			if err != nil {
				netHelper.Response(ctx, shErr.NewBadRequest(err.Error()), err, nil)
				// log.Logger.Error("noRoute.Prepare", err)
				return
			}
			cc, err := a.shareGrpcClient.Client(p.ServerName)
			if err != nil {
				log.Logger.Error("noRoute.shareGrpcClient.Client", err)
				return
			}
			err = a.shareGrpcClient.Invoke(ctx, node.GrpcPath, req, res, cc)
			netHelper.Response(ctx, res.(netHelper.Responser), err, nil)
		} else {
			netHelper.Response(ctx, shErr.NewStatusNotFound(""), nil, nil)
			return
			// err := a.addPlugin(plugin)
			// if err != nil {
			// 	netHelper.Response(ctx, errors.StatusNotFound, nil, nil)
			// 	return
			// }
			// if !isRetry {
			// 	isRetry = true
			// 	goto retry
			// }
		}

	})
	return nil
}

func (p *ShPlugin) Prepare(ctx *gin.Context, method string) (req, res interface{}, err error) {
	return p.pluginer.Prepare(ctx, method)
}

// func (a *App) addPlugin(name string) error {
// 	a.plugin.pluginMapMu.Lock()
// 	defer a.plugin.pluginMapMu.Unlock()
// 	sp, err := a.openPlugin(a.LocalConf.PluginPath + "/" + name)
// 	if err != nil {
// 		return err
// 	}
// 	a.plugin.pluginMap[sp.PluginName] = sp
// 	return nil
// }

const (
	PLUGIN_METHOD_NAME        = "Name"
	PLUGIN_METHOD_GET_ROUTERE = "GetRouter"
	PLUGIN_METHOD_PREPARE     = "Prepare"
)

func (a *App) openPlugin(pluginPath string) (err error) {

	p, err := plugin.Open(pluginPath)
	if err != nil {
		err = errors.New("openPlugin.Open:" + err.Error())
		return
	}
	shPlugin := new(ShPlugin)
	shPlugin.Plugin = p
	shPlugin.Router = router.NewTree("/", "")
	sym, err := p.Lookup(PLUGIN_NAME)
	if err != nil {
		err = errors.New("OpenPlugin.Lookup:" + err.Error())
		return
	}
	shPluginer, ok := sym.(ShPluginer)
	if ok {
		shPlugin.ServerName, shPlugin.PluginName = shPluginer.Name()

	} else {
		err = fmt.Errorf("the plugin:%s does not implement ShPluginer ", shPlugin.PluginName)
		return err
	}

	shPluginer.GetRouter()

	for k, r := range shPluginer.GetRouter() {
		point := strings.Index(k, ":")
		method := k[0:point]
		shPlugin.Router.Add(k[point+1:], method, r)
	}
	a.plugin.pluginMap[shPlugin.PluginName] = shPlugin

	return
}

//reqPath /plugin/server
//reqPath /plugin/server/:pk(number or len>32)
//reqPath /plugin/server/method(len<32,)
// /configCenter/
//mehod=/plugin.server/mehod
//SetStatus
var methodMap = map[string]string{"POST": "Create", "PUT": "Update", "DELETE": "Delete"}

func ParesRequest(ctx *gin.Context, urlPrefix string) (pluginName, reqPath, method string) {
	ctx.FullPath()

	reqPath = strings.Trim(strings.TrimLeft(path.Clean(ctx.Request.URL.Path), urlPrefix), "/")
	return helper.SubstrLeft(reqPath, "/"), helper.SubstrRight(reqPath, "/"), ctx.Request.Method

}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
