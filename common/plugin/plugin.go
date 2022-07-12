package plugin

import (
	"errors"
	"fmt"
	"path"
	"plugin"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	PLUGIN_METHOD_NAME    = "Name"
	PLUGIN_METHOD_PREPARE = "Prepare"
)

type ShPlugin struct {
	Plugin     *plugin.Plugin
	ServerName string
	PluginName string
}

func OpenPlugin(pluginPath string) (sp *ShPlugin, err error) {
	sp = new(ShPlugin)
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return
	}
	sp.Plugin = p
	sym, err := p.Lookup(PLUGIN_METHOD_NAME)
	if err != nil {
		return
	}
	if f, ok := sym.(func() (serverName, pluginName string)); ok {
		sp.ServerName, sp.PluginName = f()

	} else {
		err = errors.New("The name function of this plugin is not 'func() string'")
	}
	return
}

//reqPath /plugin/server/mehod
//mehod=/plugin.server/mehod
func ParesRequest(ctx *gin.Context, urlPrefix string) (pluginName, method string, err error) {
	reqPath := strings.TrimLeft(path.Clean(ctx.Request.URL.Path), urlPrefix)
	req := strings.Split(reqPath, "/")
	fmt.Println("ParesRequest", req)
	if len(req) != 4 {
		err = errors.New("the request url is not autoRoute")
		return
	}
	fmt.Println(req)
	pluginName = req[1]
	method = "/" + req[1] + "." + req[2] + "/" + req[3]
	return
}

func (p *ShPlugin) Prepare(ctx *gin.Context, method string) (req, res interface{}, err error) {
	symbol, err := p.Plugin.Lookup(PLUGIN_METHOD_PREPARE)
	if err != nil {
		return
	}
	f, ok := symbol.(func(ctx *gin.Context, method string) (req, res interface{}, err error))
	if !ok {
		err = errors.New("The Prepare function of this plugin is not 'func(ctx *gin.Context, method string) (req, res interface{}, err error)'")
		return
	}
	return f(ctx, method)
}
