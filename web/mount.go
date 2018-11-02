package web

import (
	"github.com/gin-gonic/gin"
	"goproj-tmpl/mid"
	"goproj-tmpl/web/router"
)

func Mount() *gin.Engine {
	eng := gin.New()
	eng.Use(mid.GinLogger(), mid.GinRecovery())

	router.Load(eng)

	return eng
}
