package router

import (
	"github.com/gin-gonic/gin"
	"web-temp/app"
	"web-temp/version"
	"web-temp/web/router/base"
	"web-temp/web/router/v1"
)

func Load(e *gin.Engine) {
	portal := NewPortalRouter()
	portal.Load(e.Group("/"))
}

func NewPortalRouter() *PortalRouter {
	return (&PortalRouter{}).init()
}

type PortalRouter struct {
	base.BaseRouter
}

func (p *PortalRouter) init() *PortalRouter {
	p.Register("/ping", NewPingRouter())
	p.Register("/api/v1", v1.NewV1Router())
	return p
}

func (p *PortalRouter) Load(rg *gin.RouterGroup) {
	rg.GET("", portalHandler)
	rg.GET("/fbv", fbvHandler)

	p.BaseRouter.Load(rg)
}

func portalHandler(c *gin.Context) {
	c.String(200, "hello portal\n")
}

func fbvHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"env":      app.GetApp().Config.Env,
		"version":  version.GVersion,
		"gin_mode": gin.Mode(),
	})
}
