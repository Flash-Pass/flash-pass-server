package middlewares

import (
	"context"
	"github.com/Flash-Pass/flash-pass-server/internal/constants"
	"time"

	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		ctx := context.Background()
		ctx = ctxlog.WithLogger(ctx)
		c.Set(constants.CtxKey, ctx)
		c.Next()

		cost := time.Since(start)

		logger := ctxlog.Extract(ctx)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}
