package fplog

import (
	"os"
	"time"

	"github.com/Flash-Pass/flash-pass-server/config"
	"github.com/mattn/go-isatty"
	"github.com/segmentio/ksuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LogFormatDev    string = "dev"
	LogFormatNormal string = "normal"
	LogFormatSplunk string = "splunk"
)

func InitLogger(variables *config.EnvVariable) *zap.Logger {
	cfg := variables.Log
	var loggerConfig zap.Config

	switch cfg.LogFormat {
	case LogFormatDev:
		loggerConfig = zap.NewDevelopmentConfig()
		if isatty.IsTerminal(os.Stdout.Fd()) {
			loggerConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
			loggerConfig.InitialFields = map[string]interface{}{
				"session": ksuid.New(),
			}
		}
	default:
		loggerConfig = zap.NewProductionConfig()

		if cfg.Level == 0 {
			loggerConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		} else {
			loggerConfig.Level = zap.NewAtomicLevelAt(cfg.Level)
		}

		loggerConfig.Development = cfg.Development
		loggerConfig.Encoding = cfg.Encoding
		loggerConfig.EncoderConfig.MessageKey = cfg.EncoderConfig.MessageKey
		loggerConfig.EncoderConfig.LevelKey = cfg.EncoderConfig.LevelKey
		loggerConfig.EncoderConfig.TimeKey = cfg.EncoderConfig.TimeKey
		loggerConfig.EncoderConfig.NameKey = cfg.EncoderConfig.NameKey
		loggerConfig.EncoderConfig.CallerKey = cfg.EncoderConfig.CallerKey
		loggerConfig.EncoderConfig.FunctionKey = cfg.EncoderConfig.FunctionKey
		loggerConfig.EncoderConfig.StacktraceKey = cfg.EncoderConfig.StacktraceKey
		loggerConfig.EncoderConfig.SkipLineEnding = cfg.EncoderConfig.SkipLineEnding
		loggerConfig.EncoderConfig.LineEnding = cfg.EncoderConfig.LineEnding
		loggerConfig.EncoderConfig.EncodeLevel = func() zapcore.LevelEncoder {
			switch cfg.EncoderConfig.LevelEncoder {
			case "capitalColor":
				return zapcore.CapitalColorLevelEncoder
			case "lowercase":
				return zapcore.LowercaseLevelEncoder
			default:
				return zapcore.CapitalLevelEncoder
			}
		}()
		loggerConfig.EncoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(time.Format("2006-01-02 15:04:05.000"))
		}
		loggerConfig.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
		loggerConfig.EncoderConfig.EncodeCaller = func() zapcore.CallerEncoder {
			switch cfg.EncoderConfig.CallerEncoder {
			case "short":
				return zapcore.ShortCallerEncoder
			default:
				return zapcore.FullCallerEncoder
			}
		}()
		loggerConfig.EncoderConfig.EncodeName = func() zapcore.NameEncoder {
			switch cfg.EncoderConfig.NameEncoder {
			case "short":
				return zapcore.FullNameEncoder
			default:
				return zapcore.FullNameEncoder
			}
		}()
		loggerConfig.EncoderConfig.ConsoleSeparator = cfg.EncoderConfig.ConsoleSeparator
		loggerConfig.OutputPaths = cfg.OutputPaths
	}

	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}

	return logger
}
