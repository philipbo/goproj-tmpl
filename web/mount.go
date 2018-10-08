package web

import (
	"github.com/gin-gonic/gin"
	"web-temp/mid"
	"web-temp/web/router"
)

func Mount() *gin.Engine {
	eng := gin.New()
	eng.Use(mid.GinLogger(), mid.GinRecovery())

	router.Load(eng)

	return eng
}
