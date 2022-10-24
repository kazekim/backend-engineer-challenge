package begin

import (
	"github.com/gin-gonic/gin"
	begincontext "github.com/kazekim/backend-engineer-challenge/grlib/begin/context"
)

type RouterGroup interface {
	Group(relativePath string) RouterGroup

	POST(relativePath string, handler HandlerFunc)
	GET(relativePath string, handler HandlerFunc)
	PUT(relativePath string, handler HandlerFunc)
	DELETE(relativePath string, handler HandlerFunc)
}

type defaultRouterGroup struct {
	g gin.RouterGroup
}

// Group create new sub RouterGroup from relativePath
func (rg *defaultRouterGroup) Group(relativePath string) RouterGroup {
	g := rg.g.Group(relativePath)
	return &defaultRouterGroup{
		g: *g,
	}
}

// POST create POST request with handler
func (rg *defaultRouterGroup) POST(relativePath string, handler HandlerFunc) {
	rg.g.POST(relativePath, func(context *gin.Context) {
		handler(begincontext.New(*context))
	})
}

// GET create GET request with handler
func (rg *defaultRouterGroup) GET(relativePath string, handler HandlerFunc) {
	rg.g.GET(relativePath, func(context *gin.Context) {
		handler(begincontext.New(*context))
	})
}

// PUT create PUT request with handler
func (rg *defaultRouterGroup) PUT(relativePath string, handler HandlerFunc) {
	rg.g.PUT(relativePath, func(context *gin.Context) {
		handler(begincontext.New(*context))
	})
}

// DELETE create DELETE request with handler
func (rg *defaultRouterGroup) DELETE(relativePath string, handler HandlerFunc) {
	rg.g.DELETE(relativePath, func(context *gin.Context) {
		handler(begincontext.New(*context))
	})
}
