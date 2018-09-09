package mid

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		latency := time.Now().Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		comment := c.Errors.ByType(gin.ErrorTypePrivate).String()
		if method == "HEAD" {
			return
		}

		entry := log.WithField("mod", "gin").
			WithField("latency", latency.String()).
			WithField("ip", clientIP).
			WithField("method", method).
			WithField("path", path).
			WithField("q", c.Request.URL.RawQuery).
			WithField("lat", fmt.Sprintf("%.2f", float64(latency.Nanoseconds())/1e6))
		if comment != "" {
			entry = entry.WithField("err", comment)
		}
		entry.Infoln(statusCode)
	}
}
