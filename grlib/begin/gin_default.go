package begin

import (

	"github.com/gin-gonic/gin"
)

type defaultGin struct {
	g               *gin.Engine
	apiVersionGroup RouterGroup
}

//New create new Gin interface
func New() Gin {

	g := gin.New()

	return &defaultGin{
		g: g,
	}
}
