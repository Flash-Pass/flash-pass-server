package ctxlog

import (
	"context"
	"github.com/Flash-Pass/flash-pass-server/internal/constants"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Export(c *gin.Context) (context.Context, *zap.Logger) {
	var ctx context.Context
	v, ok := c.Get(constants.CtxKey)
	if !ok {
		ctx = context.Background()
	}
	ctxValue, ok := v.(context.Context)
	if !ok {
		ctx = context.Background()
	}
	ctx = ctxValue
	logger := Extract(ctx)
	return ctx, logger
}
