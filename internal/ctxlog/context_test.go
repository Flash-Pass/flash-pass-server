package ctxlog

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRequestID(t *testing.T) {
	t.Run("get without setting", func(t *testing.T) {
		ctx := context.Background()
		requestId := GetRequestID(ctx)
		require.Equal(t, RequestIDNoSet, requestId)
	})

	t.Run("get request id normally", func(t *testing.T) {
		ctx := context.Background()
		ctx, source := WithRequestID(ctx)

		requestId := GetRequestID(ctx)
		require.Equal(t, source, requestId)
	})
}
