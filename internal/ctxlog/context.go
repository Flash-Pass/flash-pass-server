package ctxlog

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"go.uber.org/zap"
)

const (
	RequestIDNoSet = "request id not set"
)

const (
	RequestID = "request_id"
	Logger    = "logger"
)

func WithRequestID(ctx *gin.Context) string {
	requestId := ksuid.New().String()[0:20]
	ctx.Set(RequestID, requestId)
	return requestId
}

func GetRequestID(ctx *gin.Context) string {
	requestId, ok := ctx.Get(RequestID)
	if !ok {
		return RequestIDNoSet
	}
	return requestId.(string)
}

func WithLogger(ctx *gin.Context) {
	ctx.Set(Logger, DefaultLogger)
}

func GetLogger(ctx *gin.Context) *zap.Logger {
	value, ok := ctx.Get(Logger)
	if !ok {
		return DefaultLogger.With(zap.String("logger", "default"))
	}
	logger, ok := value.(*zap.Logger)
	if !ok {
		return DefaultLogger.With(zap.String("logger", "default"))
	}

	logger = logger.With(zap.String("logger", "ctx"))
	return WithContextFields(ctx, logger)
}
