package begincontext

import (
	"github.com/gin-gonic/gin"
)

type Context struct {
	gin.Context
	Parameters interface{}
}

// New create new context from gin's context
func New(ctx gin.Context) *Context {
	return &Context{
		Context: ctx,
	}
}
