package web

import (
	"github.com/gin-gonic/gin"
	"web-temp/app"
	"web-temp/mid"
	"web-temp/version"
)

func Mount() *gin.Engine {
	router := gin.New()
	router.Use(mid.GinLogger(), mid.GinRecovery())
	fbv := func(c *gin.Context) {
		c.JSON(200, gin.H{
			"env":      app.GetApp().Flags.Env,
			"gin_mode": gin.Mode(),
			"version":  version.GVersion,
		})
	}

	router.GET("/", fbv)

	pong := func(c *gin.Context) {
		c.String(200, "pong\n")
	}

	router.GET("/ping", pong)
	router.HEAD("/ping", pong)

	return router
}
