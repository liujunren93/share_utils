package router

import (
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

const WHITE_PREFIX = "White"

var whiteMap = make(map[string][]string)

type Router struct {
	IsGroup      bool
	Engine       *gin.Engine
	group        *gin.RouterGroup
	relativePath string
}

func NewRouter(eng *gin.Engine) Router {
	return Router{
		Engine: eng,
		group:  eng.Group(""),
	}

}

func GetChan() {
}
func InWhitelist(ctx *gin.Context, table string) bool {
	fullpath := ctx.FullPath()

	if list, ok := whiteMap[table]; ok {
		for _, v := range list {
			if strings.Index(fullpath, v) == 0 {
				return true
			}
		}
	}
	return false
}

//White set WhiteList
// prefix:White
func (g Router) White(table string) Router {

	rPath := g.group.BasePath()
	if !g.IsGroup {
		rPath += "/" + g.relativePath
	}
	rPath = path.Clean(rPath)
	whiteMap[table] = append(whiteMap[table], rPath)
	return g
}

func (g Router) Use(ms ...gin.HandlerFunc) Router {
	g.group.Use(ms...)
	return g

}

func (g Router) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) Router {
	g.group.Handle(httpMethod, relativePath, handlers...)
	g.relativePath = relativePath
	return g
}

func (g Router) Any(relativePath string, handlers ...gin.HandlerFunc) Router {
	g.group.Any(relativePath, handlers...)
	g.relativePath = relativePath
	return g
}

func (g Router) GET(relativePath string, handlers ...gin.HandlerFunc) Router {
	g.group.GET(relativePath, handlers...)
	g.relativePath = relativePath
	return g
}

func (g Router) POST(relativePath string, handlers ...gin.HandlerFunc) Router {
	g.group.POST(relativePath, handlers...)
	g.relativePath = relativePath
	return g
}

func (g Router) DELETE(relativePath string, handlers ...gin.HandlerFunc) Router {
	g.group.DELETE(relativePath, handlers...)
	g.relativePath = relativePath
	return g
}

func (g Router) PATCH(relativePath string, handlers ...gin.HandlerFunc) Router {
	g.group.PATCH(relativePath, handlers...)
	g.relativePath = relativePath
	return g
}

func (g Router) PUT(relativePath string, handlers ...gin.HandlerFunc) Router {
	g.group.PUT(relativePath, handlers...)
	g.relativePath = relativePath
	return g
}

func (g Router) NoRoute(handlers ...gin.HandlerFunc) {
	// groupHandler := g.group.Handlers
	// handlers = append(groupHandler, handlers...)
	g.Engine.NoRoute(handlers...)
}

func (g Router) OPTIONS(relativePath string, handlers ...gin.HandlerFunc) Router {
	g.group.OPTIONS(relativePath, handlers...)
	g.relativePath = relativePath
	return g
}

func (g Router) HEAD(relativePath string, handlers ...gin.HandlerFunc) Router {
	g.group.Any(relativePath, handlers...)
	g.relativePath = relativePath
	return g
}

func (g Router) StaticFile(relativePath, filepath string) Router {
	g.group.StaticFile(relativePath, filepath)
	g.relativePath = relativePath
	return g
}

func (g Router) StaticFileFS(relativePath, filepath string, fs http.FileSystem) Router {
	g.group.StaticFileFS(relativePath, filepath, fs)
	g.relativePath = relativePath
	return g
}

func (g Router) Static(relativePath, root string) Router {
	g.group.Static(relativePath, root)
	g.relativePath = relativePath
	return g
}

func (g Router) StaticFS(relativePath string, fs http.FileSystem) Router {
	g.group.StaticFS(relativePath, fs)
	g.relativePath = relativePath
	return g
}

func (g Router) Group(relativePath string, handlers ...gin.HandlerFunc) Router {
	return Router{
		IsGroup:      true,
		Engine:       g.Engine,
		group:        g.group.Group(relativePath, handlers...),
		relativePath: relativePath,
	}
}
