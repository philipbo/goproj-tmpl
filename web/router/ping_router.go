package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"web-temp/app"
	"web-temp/web/router/base"
)

func NewPingRouter() *PingRouter {
	return &PingRouter{}
}

type PingRouter struct {
	base.BaseRouter
}

func (p *PingRouter) Load(rg *gin.RouterGroup) {
	rg.HEAD("", pingHandler)
	rg.GET("", pingHandler)
}

func pingHandler(c *gin.Context) {
	c.String(200, fmt.Sprintf("pong for %s\n", app.GetApp().Config.Env))
}
