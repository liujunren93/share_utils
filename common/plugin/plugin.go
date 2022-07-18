package plugin

import (
	"errors"
	"path"
	"plugin"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	shErr "github.com/liujunren93/share_utils/errors"
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
		err = errors.New("OpenPlugin.Open:" + err.Error())
		return
	}
	sp.Plugin = p
	sym, err := p.Lookup(PLUGIN_METHOD_NAME)
	if err != nil {
		err = errors.New("OpenPlugin.Lookup:" + err.Error())
		return
	}
	if f, ok := sym.(func() (serverName, pluginName string)); ok {
		sp.ServerName, sp.PluginName = f()

	} else {
		err = errors.New("The name function of this plugin is not 'func() (string,string)'")
	}
	return
}

//reqPath /plugin/server
//reqPath /plugin/server/:pk(number or len>32)
//reqPath /plugin/server/method(len<32,)
// /configCenter/
//mehod=/plugin.server/mehod
//SetStatus
var methodMap = map[string]string{"POST": "Create", "PUT": "Update", "DELETE": "Delete"}

func ParesRequest(ctx *gin.Context, urlPrefix string) (pluginName, method string, err error) {
	ctx.FullPath()

	reqPath := strings.Trim(strings.TrimLeft(path.Clean(ctx.Request.URL.Path), urlPrefix), "/")
	req := strings.Split(reqPath, "/")
	if len(req) < 2 {
		err = errors.New("the request url is not autoRoute")
		return
	}
	m := methodMap[ctx.Request.Method]

	if len(req) > 2 && len(req[2]) < 32 && !IsNum(req[2]) {
		m = req[2]
	} else {

		if ctx.Request.Method == "GET" {
			if len(req) > 2 { // info
				ctx.Params = append(ctx.Params, gin.Param{Key: "pk", Value: req[2]})
				m = "Info"
			} else {
				m = "List"
			}
		}

		if (ctx.Request.Method == "PUT" || ctx.Request.Method == "DELETE") && len(req) > 2 {

			ctx.Params = append(ctx.Params, gin.Param{Key: "pk", Value: req[2]})
		}
	}

	pluginName = req[0]
	method = "/" + req[0] + "." + req[1] + "/" + m
	return
}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// func ParesRequest(ctx *gin.Context, urlPrefix string) (pluginName, method string, err error) {

// 	reqPath := strings.TrimLeft(path.Clean(ctx.Request.URL.Path), urlPrefix)
// 	req := strings.Split(reqPath, "/")
// 	fmt.Println("ParesRequest", req)
// 	if len(req) != 4 {
// 		err = errors.New("the request url is not autoRoute")
// 		return
// 	}
// 	fmt.Println(req)
// 	pluginName = req[1]
// 	method = "/" + req[1] + "." + req[2] + "/" + req[3]
// 	return
// }

func (p *ShPlugin) Prepare(ctx *gin.Context, method string) (req, res interface{}, err error) {
	symbol, err := p.Plugin.Lookup(PLUGIN_METHOD_PREPARE)
	if err != nil {
		return
	}
	f, ok := symbol.(func(ctx *gin.Context, method string) (req, res interface{}, err shErr.Error))
	if !ok {
		err = errors.New("The Prepare function of this plugin is not 'func(ctx *gin.Context, method string) (req, res interface{}, err errors.Error)'")
		return
	}
	return f(ctx, method)
}
