package config

import (
	"errors"

	"github.com/spf13/viper"
)

var (
	config *Config
)

// 服务器地址
type Server struct {
	Url     string `yaml:"url"`
	Secret  string `yaml:"secret"`
	Content string `yaml:"content"`
}

type Logger struct {
	MaxAge      int    `yaml:"maxAge"`
	MaxSize     int    `yaml:"maxSize"`
	MaxBackup   int    `yaml:"maxBackup"`
	Compress    bool   `yaml:"compress"`
	Level       int    `yaml:"level"`
	LogPath     string `yaml:"logPath"`
	ServiceName string `yaml:"serviceName"`
}

type Config struct {
	Server
	Logger
}

func init() {
	newConfig()
	loadConfig()
}

func newConfig() *Config {
	config = &Config{}
	return config
}

func GetRestUrl() string {
	return config.Server.Url
}

func GetServerSecret() string {
	return config.Server.Secret
}

func GetServerContent() string {
	return config.Server.Content
}

func GetLoggerLevel() int {
	return config.Logger.Level
}
func GetLoggerPath() string {
	return config.Logger.LogPath
}

func GetLoggerServiceName() string {
	return config.Logger.ServiceName
}

func GetLoggerMaxAge() int {
	return config.Logger.MaxAge
}

func GetLoggerMaxSize() int {
	return config.Logger.MaxSize
}

func GetLoggerMaxBackup() int {
	return config.Logger.MaxBackup
}

func GetLoggerCompress() bool {
	return config.Logger.Compress
}

func loadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()

	if err != nil {
		panic(errors.New("反序列化配置文件出错"))
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(errors.New("反序列化配置文件出错"))
	}
}
