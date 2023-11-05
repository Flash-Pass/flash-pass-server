package ctxlog

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestRequestID(t *testing.T) {
	t.Run("get without setting", func(t *testing.T) {
		ctx := &gin.Context{}
		requestId := GetRequestID(ctx)
		require.Equal(t, RequestIDNoSet, requestId)
	})

	t.Run("get request id normally", func(t *testing.T) {
		ctx := &gin.Context{}
		source := WithRequestID(ctx)

		requestId := GetRequestID(ctx)
		require.Equal(t, source, requestId)
	})
}
