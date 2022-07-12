package app

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/gin-gonic/gin"
	shareRouter "github.com/liujunren93/share_utils/common/gin/router"
	myplugin "github.com/liujunren93/share_utils/common/plugin"
	"github.com/liujunren93/share_utils/errors"
	"github.com/liujunren93/share_utils/log"
	"github.com/liujunren93/share_utils/netHelper"
)

type plugin struct {
	pluginMapMu *sync.RWMutex
	pluginMap   map[string]*myplugin.ShPlugin
}

func (a *App) initPlugins() {
	fmt.Println("initPlugins")
	a.plugin = new(plugin)
	a.plugin.pluginMapMu = &sync.RWMutex{}
	a.plugin.pluginMap = make(map[string]*myplugin.ShPlugin)
	PluginPath := a.LocalConf.PluginPath
	files, err := os.ReadDir(PluginPath)
	if err != nil {
		fmt.Println("loadPlugins.ReadDir", err, PluginPath)
		log.Logger.Fatal("loadPlugins.ReadDir", err, PluginPath)
	}

	for _, f := range files {
		fmt.Println("loadPlug", f.Name())
		if filepath.Ext(f.Name()) == ".so" {
			sp, err := myplugin.OpenPlugin(PluginPath + "/" + f.Name())

			if err != nil {
				fmt.Println("loadPlugins.OpenPlugin", PluginPath+"/"+f.Name(), err)
				log.Logger.Error("loadPlugins.OpenPlugin", PluginPath+"/"+f.Name(), err)
				continue
			}
			fmt.Println(sp)
			log.Logger.Debug("initPlugins", sp)
			a.plugin.pluginMap[sp.PluginName] = sp
		}
	}
}

func (a *App) AutoRoute(r shareRouter.Router) error {
	fmt.Println("AutoRoute", a.plugin.pluginMap)
	r.NoRoute(func(ctx *gin.Context) {
		plugin, method, err := myplugin.ParesRequest(ctx, a.LocalConf.ApiPrefix)
		fmt.Println("AutoRoute.ParesRequest", plugin, method, err)
		if err != nil {
			log.Logger.Error("noRoute.ParesRequest", err)
			return
		}
		// 	isRetry := false
		// retry:
		a.plugin.pluginMapMu.RLock()
		p, ok := a.plugin.pluginMap[plugin]
		fmt.Println(plugin, p)
		a.plugin.pluginMapMu.RUnlock()
		if ok {
			req, res, err := p.Prepare(ctx, method)
			if err != nil {
				netHelper.Response(ctx, errors.NewBadRequest(err.Error()), err, nil)
				// log.Logger.Error("noRoute.Prepare", err)
				return
			}
			cc, err := a.shareGrpcClient.Client(p.ServerName)
			if err != nil {
				log.Logger.Error("noRoute.shareGrpcClient.Client", err)
				return
			}
			err = a.shareGrpcClient.Invoke(ctx, method, req, res, cc)
			netHelper.Response(ctx, res.(netHelper.Responser), err, nil)
		} else {
			netHelper.Response(ctx, errors.StatusNotFound, nil, nil)
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
func (a *App) addPlugin(name string) error {
	a.plugin.pluginMapMu.Lock()
	defer a.plugin.pluginMapMu.Unlock()
	sp, err := myplugin.OpenPlugin(a.LocalConf.PluginPath + "/" + name)
	if err != nil {
		return err
	}
	a.plugin.pluginMap[sp.PluginName] = sp
	return nil
}
