package router

import (
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

const WHITE_PREFIX = "White"

var whiteMap = make(map[string][]string)

type RouterGroup struct {
	group        *gin.RouterGroup
	relativePath string
}

func NewRouterGroup(irouter *gin.RouterGroup) RouterGroup {
	return RouterGroup{
		group: irouter,
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
func (g RouterGroup) White(table string) RouterGroup {
	g.group.BasePath()
	var rPath string

	rPath = g.group.BasePath() + "////" + g.relativePath
	rPath = path.Clean(rPath)
	whiteMap[table] = append(whiteMap[table], rPath)
	return g
}

func (g RouterGroup) Use(ms ...gin.HandlerFunc) RouterGroup {
	g.group.Use(ms...)
	return g

}

func (g RouterGroup) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) RouterGroup {
	g.group.Handle(httpMethod, relativePath, handlers...)
	g.relativePath = relativePath
	return g
}

func (g RouterGroup) Any(relativePath string, handlers ...gin.HandlerFunc) RouterGroup {
	g.group.Any(relativePath, handlers...)
	g.relativePath = relativePath
	return g
}

func (g RouterGroup) GET(relativePath string, handlers ...gin.HandlerFunc) RouterGroup {
	g.group.GET(relativePath, handlers...)
	g.relativePath = relativePath
	return g
}

func (g RouterGroup) POST(relativePath string, handlers ...gin.HandlerFunc) RouterGroup {
	g.group.POST(relativePath, handlers...)
	g.relativePath = relativePath
	return g
}

func (g RouterGroup) DELETE(relativePath string, handlers ...gin.HandlerFunc) RouterGroup {
	g.group.DELETE(relativePath, handlers...)
	g.relativePath = relativePath
	return g
}

func (g RouterGroup) PATCH(relativePath string, handlers ...gin.HandlerFunc) RouterGroup {
	g.group.PATCH(relativePath, handlers...)
	g.relativePath = relativePath
	return g
}

func (g RouterGroup) PUT(relativePath string, handlers ...gin.HandlerFunc) RouterGroup {
	g.group.PUT(relativePath, handlers...)
	g.relativePath = relativePath
	return g
}

func (g RouterGroup) OPTIONS(relativePath string, handlers ...gin.HandlerFunc) RouterGroup {
	g.group.OPTIONS(relativePath, handlers...)
	g.relativePath = relativePath
	return g
}

func (g RouterGroup) HEAD(relativePath string, handlers ...gin.HandlerFunc) RouterGroup {
	g.group.Any(relativePath, handlers...)
	g.relativePath = relativePath
	return g
}

func (g RouterGroup) StaticFile(relativePath, filepath string) RouterGroup {
	g.group.StaticFile(relativePath, filepath)
	g.relativePath = relativePath
	return g
}

func (g RouterGroup) StaticFileFS(relativePath, filepath string, fs http.FileSystem) RouterGroup {
	g.group.StaticFileFS(relativePath, filepath, fs)
	g.relativePath = relativePath
	return g
}

func (g RouterGroup) Static(relativePath, root string) RouterGroup {
	g.group.Static(relativePath, root)
	g.relativePath = relativePath
	return g
}

func (g RouterGroup) StaticFS(relativePath string, fs http.FileSystem) RouterGroup {
	g.group.StaticFS(relativePath, fs)
	g.relativePath = relativePath
	return g
}

func (g RouterGroup) Group(relativePath string, handlers ...gin.HandlerFunc) RouterGroup {
	g.group = g.group.Group(relativePath, handlers...)
	g.relativePath = ""
	return g
}
