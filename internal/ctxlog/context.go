package ctxlog

import (
	"context"
	"github.com/segmentio/ksuid"
)

const (
	RequestIDNoSet = "request id not set"
)

const (
	RequestID = "request_id"
	Logger    = "logger"
)

func WithRequestID(ctx context.Context) (context.Context, string) {
	requestId := ksuid.New().String()[0:20]
	ctx = context.WithValue(ctx, RequestID, requestId)
	return ctx, requestId
}

func GetRequestID(ctx context.Context) string {
	v := ctx.Value(RequestID)
	if v == nil {
		return RequestIDNoSet
	}
	requestId, ok := v.(string)
	if !ok {
		return RequestIDNoSet
	}
	return requestId
}
