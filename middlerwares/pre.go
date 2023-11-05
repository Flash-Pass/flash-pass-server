package middlewares

import (
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"github.com/gin-gonic/gin"
)

func PreRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctxlog.WithRequestID(c)
		ctxlog.WithLogger(c)
		c.Next()
	}
}
