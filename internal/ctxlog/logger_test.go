package ctxlog

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestExtractLogger(t *testing.T) {
	cases := []struct {
		description  string
		given        *gin.Context
		expected     []zap.Field
		loggerExists bool
	}{
		{
			description: "empty context",
			given:       &gin.Context{},
			expected:    []zap.Field{zap.String("logger", "default")},
		},
	}

	for _, item := range cases {
		item := item
		t.Run(item.description, func(t *testing.T) {
			ctx := item.given
			if item.loggerExists {
				coreLogger, observedLogs := observer.New(zap.InfoLevel)
				existingLogger := zap.New(coreLogger)
				ctx.Set(loggerKey, existingLogger)

				extractLogger := GetLogger(ctx)
				extractLogger.Info("doing log")
				allLogs := observedLogs.All()
				require.ElementsMatch(t, item.expected, allLogs[0].Context)
			}

			expectedLogger := zap.NewNop().With(item.expected...)
			require.Equal(t, expectedLogger, GetLogger(ctx))
		})
	}
}

func TestAddLoggerToContext(t *testing.T) {
	given := DefaultLogger
	ctx := &gin.Context{}
	WithLogger(ctx)

	value, ok := ctx.Get(Logger)
	require.True(t, ok)

	logger, ok := value.(*zap.Logger)
	require.True(t, ok)
	require.Equal(t, given, logger)
}

func TestAddFields_DoesNotModifyOriginalNamespaces(t *testing.T) {
	namespace := zap.Namespace("test_space")
	originalFields := map[zap.Field][]zap.Field{
		namespace: {
			zap.String("original field", "original value"),
		},
	}

	ctx := &gin.Context{}
	ctx.Set(namespaceKey, originalFields)
	ctx = AddFields(ctx, namespace, zap.String("new field", "new value"))
	require.Equal(t, map[zap.Field][]zap.Field{
		namespace: {
			zap.String("original field", "original value"),
		},
	}, originalFields)

	namespaces, ok := ctx.Get(namespaceKey)
	require.True(t, ok)
	require.Equal(t, map[zap.Field][]zap.Field{
		namespace: {
			zap.String("original field", "original value"),
			zap.String("new field", "new value"),
		},
	}, namespaces)
}
