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

type nameFunc func() string
type PrepareFunc func(ctx *gin.Context, method string) (req, res interface{}, err error)

func OpenPlugin(pluginPath string) (p *plugin.Plugin, pluginName string, err error) {
	p, err = plugin.Open(pluginPath)
	if err != nil {
		return nil, "", err
	}
	sym, err := p.Lookup(PLUGIN_METHOD_NAME)
	if err != nil {
		return nil, "", err
	}
	if f, ok := sym.(func() string); ok {
		pluginName = f()
	} else {
		return nil, "", errors.New("The name function of this plugin is not 'func() string'")
	}
	return
}

//reqPath plugin/server/mehod
//mehod=plugin.server/mehod
func ParesRequest(reqPath, urlPrefix string) (plugin, server, method string) {
	reqPath = strings.TrimLeft(strings.TrimLeft(reqPath, "/"), urlPrefix)
	reqPath = path.Clean(reqPath)
	req := strings.Split(reqPath, "/")
	fmt.Println(req)
	plugin = req[0]
	server = req[0] + "." + req[1]
	method = req[0] + "." + req[1] + "/" + req[2]
	return
}
