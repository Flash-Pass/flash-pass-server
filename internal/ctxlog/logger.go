package ctxlog

import (
	"context"
	"reflect"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	DefaultNamespace = zap.Namespace("DouTokDefault")
)

var DefaultLogger = zap.NewNop()

type loggerKey struct{}

//type namespaceKey struct{}
//type globalFieldKey struct{}

const (
	namespaceKey   = "namespace_key"
	globalFieldKey = "global_field_key"
)

func Extract(ctx context.Context) *zap.Logger {
	v := ctx.Value(loggerKey{})
	if v == nil {
		return DefaultLogger.With(zap.String("logger", "default"))
	}

	logger, ok := v.(*zap.Logger)
	if !ok {
		return DefaultLogger.With(zap.String("logger", "default"))
	}

	logger = logger.With(zap.String("logger", "ctx"))
	return WithContextFields(ctx, logger)
}

func WithLogger(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, loggerKey{}, DefaultLogger)
	return ctx
}

func SetLogger(logger *zap.Logger) *zap.Logger {
	DefaultLogger = logger
	return logger
}

// AddFields will add or override fields to context when zap.Namespace is provided.
// It will nest any following fields until another zap.Namespace is provided.
// The difference with pure zap.Namespaces are that when you define a namespace, any following attached
// logger will be nested, not only when you call logger.With
func AddFields(ctx context.Context, fields ...zap.Field) context.Context {
	namespaces := extractNamespaces(ctx)
	globalFields := extractGlobalfields(ctx)

	var namespace zap.Field
	for _, field := range fields {
		if field.Type == zapcore.NamespaceType {
			namespace = field
			continue
		}

		if !reflect.ValueOf(namespace).IsZero() {
			namespaces[namespace] = append(namespaces[namespace], field)
		} else {
			globalFields = append(globalFields, field)
		}
	}

	ctx = context.WithValue(ctx, namespaceKey, namespaces)
	ctx = context.WithValue(ctx, globalFieldKey, globalFields)

	return ctx
}

// WithContextFields adds context fields to the zap.Logger
func WithContextFields(ctx context.Context, logger *zap.Logger) *zap.Logger {
	ctx = AddFields(
		ctx, zap.String("RequestId", GetRequestID(ctx)),
	)

	// TODO: Add more public fields here.

	logger = logger.With(extractGlobalfields(ctx)...)
	for namespace, fields := range extractNamespaces(ctx) {
		logger = logger.With(Nest(namespace.Key, fields...))
	}

	return logger
}

func extractNamespaces(ctx context.Context) map[zapcore.Field][]zapcore.Field {
	v := ctx.Value(namespaceKey)
	if v == nil {
		return nil
	}

	ctxValue, ok := v.(map[zapcore.Field][]zapcore.Field)
	if !ok {
		return nil
	}
	namespaces := make(map[zapcore.Field][]zapcore.Field)
	if ctxValue != nil {
		ctxNamespaces := ctxValue
		if ok {
			for k, v := range ctxNamespaces {
				namespaces[k] = v
			}
		}
	}
	return namespaces
}

func extractGlobalfields(ctx context.Context) []zapcore.Field {
	v := ctx.Value(globalFieldKey)
	if v == nil {
		return nil
	}

	ctxValue, ok := v.([]zapcore.Field)
	if !ok {
		return nil
	}
	globalFields := make([]zap.Field, 0)
	if ctxValue != nil {
		globalFields = ctxValue
	}
	return globalFields
}
