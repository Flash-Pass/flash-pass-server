package config

import (
	"context"

	"github.com/caarlos0/env/v9"
	"go.uber.org/zap/zapcore"
)

type EnvVariable struct {
	BASE struct {
		Name    string `env:"SERVER_NAME" envDefault:"flash_pass_server"`
		Address string `env:"SERVER_ADDRESS" envDefault:"localhost"`
		Port    int    `env:"SERVER_PORT" envDefault:"8080"`
	} `envPrefix:"BASE_"`
	MySQL struct {
		Host      string `env:"HOST" envDefault:"localhost"`
		Port      string `env:"PORT" envDefault:"3306"`
		Username  string `env:"USERNAME" envDefault:"root"`
		Password  string `env:"PASSWORD" envDefault:"root"`
		Database  string `env:"DATABASE" envDefault:"flash_pass"`
		CharSet   string `env:"CHARSET" envDefault:"utf8"`
		ParseTime string `env:"PARSE_TIME" envDefault:"true"`
		Loc       string `env:"LOC" envDefault:"Local"`
	} `envPrefix:"MYSQL_"`
	Log struct {
		LogFormat     string        `env:"LOGGER_FORMAT" envDefault:"dev"` // json, console
		Level         zapcore.Level `env:"LOGGER_LEVEL" envDefault:"info"`
		Development   bool          `env:"LOGGER_DEVELOPMENT" envDefault:"false"`
		Encoding      string        `env:"LOGGER_ENCODING" envDefault:"json"`
		EncoderConfig struct {
			MessageKey       string `env:"MESSAGE_KEY" envDefault:"message"`
			LevelKey         string `env:"LEVEL_KEY" envDefault:"level"`
			TimeKey          string `env:"TIME_KEY" envDefault:"ts"`
			NameKey          string `env:"NAME_KEY" envDefault:"DouTokLogger"`
			CallerKey        string `env:"CALLER_KEY" envDefault:"caller"`
			FunctionKey      string `env:"FUNCTION_KEY" envDefault:"function"`
			StacktraceKey    string `env:"STACKTRACE_KEY" envDefault:"stacktrace"`
			SkipLineEnding   bool   `env:"SKIP_LINE_ENDING" envDefault:"false"`
			LineEnding       string `env:"LINE_ENDING" envDefault:"\n"`
			LevelEncoder     string `env:"LEVEL_ENCODER" envDefault:"capital"`    // capitalColor, capital, color, lowercase
			DurationEncoder  string `env:"DURATION_ENCODER" envDefault:"seconds"` // string, nanos, ms, seconds
			CallerEncoder    string `env:"CALLER_ENCODER" envDefault:"full"`      // short, full
			NameEncoder      string `env:"NAME_ENCODER" envDefault:"full"`        // short, full
			ConsoleSeparator string `env:"CONSOLE_SEPARATOR" envDefault:" "`
		} `envPrefix:"ENCODER_"`
		OutputPaths []string `env:"OUTPUT_PATHS" envDefault:"stdout" envSeparator:","`
	} `envPrefix:"LOG_"`
}

func LoadEnv(ctx context.Context) (*EnvVariable, error) {
	config := &EnvVariable{}
	if err := env.Parse(config); err != nil {
		return nil, err
	}

	return config, nil
}
