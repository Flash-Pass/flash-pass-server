package ctxlog

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestExtractLogger(t *testing.T) {
	cases := []struct {
		description  string
		given        context.Context
		expected     []zap.Field
		loggerExists bool
	}{
		{
			description: "empty context",
			given:       context.Background(),
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
				ctx = context.WithValue(ctx, loggerKey{}, existingLogger)

				extractLogger := Extract(ctx)
				extractLogger.Info("doing log")
				allLogs := observedLogs.All()
				require.ElementsMatch(t, item.expected, allLogs[0].Context)
			}

			expectedLogger := zap.NewNop().With(item.expected...)
			require.Equal(t, expectedLogger, Extract(ctx))
		})
	}
}

func TestAddLoggerToContext(t *testing.T) {
	given := DefaultLogger
	ctx := context.Background()
	ctx = WithLogger(ctx)

	v := ctx.Value(loggerKey{})
	require.NotNil(t, v)
	logger, ok := v.(*zap.Logger)
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

	ctx := context.Background()
	ctx = context.WithValue(ctx, namespaceKey, originalFields)
	ctx = AddFields(ctx, namespace, zap.String("new field", "new value"))
	require.Equal(t, map[zap.Field][]zap.Field{
		namespace: {
			zap.String("original field", "original value"),
		},
	}, originalFields)

	namespaces, ok := ctx.Value(namespaceKey).(map[zap.Field][]zap.Field)
	require.True(t, ok)
	require.Equal(t, map[zap.Field][]zap.Field{
		namespace: {
			zap.String("original field", "original value"),
			zap.String("new field", "new value"),
		},
	}, namespaces)
}
