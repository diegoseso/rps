package config

import (
	"github.com/spf13/viper"
	"net/url"
)

const CONFIG_FILE_PATTERN = "config"

type SocketServerConfig struct {
	Host   string
	Port   string
	Path   string
}

type LoggerConfig struct {
	Level   string
}

var initialized bool

func Load(configPath *string) error {
	viper.SetConfigName(CONFIG_FILE_PATTERN) 
	viper.AddConfigPath(*configPath)      
	err := viper.ReadInConfig()          
	if err == nil {
		initialized = true
	}
	return err
}

func GetSocketServerConfig() *SocketServerConfig {
	if !initialized {
		panic("Config is not initialized")
	}
	return &SocketServerConfig{
		Host:   viper.GetString("http-server.host"),
		Port:   viper.GetString("http-server.port"),
		Path:   viper.GetString("http-server.path"),
	}
}

func GetHttpServerUrl() *url.URL {
	httpServerConfig := GetHttpServerConfig()
	return &url.URL{Scheme: httpServerConfig.Scheme, Host: httpServerConfig.Host + ":" + httpServerConfig.Port, Path: httpServerConfig.Path}
}

func GetLoggerConfig() *LoggerConfig{
	if !initialized {
		panic("Config is not initialized")
	}
	return &LoggerConfig{
		Level:   viper.GetString("logger.level"),
	}
}