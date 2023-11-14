package config

import (
	"context"
	"encoding/json"
	"log"

	"github.com/caarlos0/env/v9"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap/zapcore"
)

type EnvVariable struct {
	BASE   BaseConfig   `envPrefix:"BASE_"`
	MySQL  MySQLConfig  `envPrefix:"MYSQL_"`
	WeChat WeChatConfig `envPrefix:"WECHAT_"`
	Log    LogConfig    `envPrefix:"LOG_"`
}

type BaseConfig struct {
	Name    string      `env:"SERVER_NAME" envDefault:"flash_pass_server"`
	Address string      `env:"SERVER_ADDRESS" envDefault:"localhost"`
	Port    int         `env:"SERVER_PORT" envDefault:"8080"`
	Nacos   NacosConfig `envPrefix:"NACOS_"`
}

type NacosConfig struct {
	Address       string `env:"ADDRESS" envDefault:"localhost"`
	Port          uint64 `env:"PORT" envDefault:"8848"`
	Namespace     string `env:"NAMESPACE" envDefault:""`
	NamespaceUser string `env:"NAMESPACE_USER" envDefault:""`
}

type MySQLConfig struct {
	Host      string `env:"HOST" envDefault:"localhost" json:"host"`
	Port      int    `env:"PORT" envDefault:"3306" json:"port"`
	Username  string `env:"USERNAME" envDefault:"root" json:"username"`
	Password  string `env:"PASSWORD" envDefault:"root" json:"password"`
	Database  string `env:"DATABASE" envDefault:"flash_pass" json:"database"`
	CharSet   string `env:"CHARSET" envDefault:"utf8" json:"charset"`
	ParseTime string `env:"PARSE_TIME" envDefault:"true" json:"parse_time"`
	Loc       string `env:"LOC" envDefault:"Local" json:"loc"`
}

type WeChatConfig struct {
	AppId  string `env:"APP_ID" envDefault:""`
	Secret string `env:"SECRET" envDefault:""`
}

type LogConfig struct {
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
}

func LoadEnv(ctx context.Context) (*EnvVariable, error) {
	config := &EnvVariable{}
	if err := env.Parse(config); err != nil {
		return nil, err
	}

	dbConfig, err := loadDBFromNacos(config.BASE)
	if err == nil {
		config.MySQL = dbConfig
	}

	return config, nil
}

func loadDBFromNacos(cfg BaseConfig) (MySQLConfig, error) {
	dbConfig := MySQLConfig{}

	sc := []constant.ServerConfig{
		{
			IpAddr: cfg.Nacos.Address,
			Port:   cfg.Nacos.Port,
		},
	}
	if cfg.Address != "" {
		sc[0].IpAddr = cfg.Address
	}

	cc := constant.ClientConfig{
		NamespaceId: cfg.Nacos.Namespace,
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})

	if err != nil {
		log.Println("创建 Nacos 配置客户端失败:", err)
		return dbConfig, err
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "database",
		Group:  "FLASH_PASS",
	})

	if err != nil {
		log.Println("获取 Nacos 配置失败:", err)
		return dbConfig, err
	}

	err = json.Unmarshal([]byte(content), &dbConfig)
	if err != nil {
		log.Println("解析配置失败:", err)
		return dbConfig, err
	}

	// 本地 dev 环境开发人员的 db 相互独立，通过 username 区分，该 db 需要事先联系管理员创建
	if dbConfig.Database == "" {
		dbConfig.Database = cfg.Nacos.NamespaceUser
	}

	return dbConfig, nil
}
