package v1

import (
	"github.com/gin-gonic/gin"
	"web-temp/web/router/base"
)

func NewExampleRouter() *ExampleRouter {
	return &ExampleRouter{}
}

type ExampleRouter struct {
	base.BaseRouter
}

func (r *ExampleRouter) Load(rg *gin.RouterGroup) {
	rg.GET("", exampleHandler)
}

func exampleHandler(c *gin.Context) {
	c.String(200, "example router for v1\n")
}
