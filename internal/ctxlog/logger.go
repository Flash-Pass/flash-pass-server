package ctxlog

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	DefaultNamespace = zap.Namespace("DouTokDefault")
)

var DefaultLogger = zap.NewNop()

//type loggerKey struct{}
//type namespaceKey struct{}
//type globalFieldKey struct{}

const (
	loggerKey      = "logger_key"
	namespaceKey   = "namespace_key"
	globalFieldKey = "global_field_key"
)

func SetLogger(logger *zap.Logger) *zap.Logger {
	DefaultLogger = logger
	return logger
}

// AddFields will add or override fields to context when zap.Namespace is provided.
// It will nest any following fields until another zap.Namespace is provided.
// The difference with pure zap.Namespaces are that when you define a namespace, any following attached
// logger will be nested, not only when you call logger.With
func AddFields(ctx *gin.Context, fields ...zap.Field) *gin.Context {
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

	ctx.Set(namespaceKey, namespaces)
	ctx.Set(globalFieldKey, globalFieldKey)

	return ctx
}

// WithContextFields adds context fields to the zap.Logger
func WithContextFields(ctx *gin.Context, logger *zap.Logger) *zap.Logger {
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

func extractNamespaces(ctx *gin.Context) map[zapcore.Field][]zapcore.Field {
	ctxValue, ok := ctx.Get(namespaceKey)
	if !ok {
		return nil
	}
	namespaces := make(map[zapcore.Field][]zapcore.Field)
	if ctxValue != nil {
		ctxNamespaces, ok := ctxValue.(map[zapcore.Field][]zapcore.Field)
		if ok {
			for k, v := range ctxNamespaces {
				namespaces[k] = v
			}
		}
	}
	return namespaces
}

func extractGlobalfields(ctx *gin.Context) []zapcore.Field {
	ctxValue, ok := ctx.Get(globalFieldKey)
	if !ok {
		return nil
	}
	globalFields := make([]zap.Field, 0)
	if ctxValue != nil {
		value, ok := ctxValue.([]zap.Field)
		if ok {
			globalFields = value
		}
	}
	return globalFields
}
