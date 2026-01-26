package middleware

import (
	"ayam_bangkok/source/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        latency := time.Since(start)
        status := c.Writer.Status()

        ev := logger.Info()
        if status >= 500 {
            ev = logger.Error()
        } else if status >= 400 {
            ev = logger.Warn()
        }

        ev.
            Str("method", c.Request.Method).
            Str("path", c.FullPath()).
            Int("status", status).
            Dur("latency", latency).
            Str("client_ip", c.ClientIP()).
            Msg("HTTP request")
    }
}