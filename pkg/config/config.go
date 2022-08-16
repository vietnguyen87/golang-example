package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

func Load() {
	viper.AddConfigPath("$APP_HOME")
	viper.AddConfigPath(".")

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("viper.ReadInConfig returns error: %s", err.Error())
	}

	viper.SetConfigName("override")
	if err := viper.MergeInConfig(); err != nil {
		if !strings.Contains(strings.ToLower(err.Error()), "not found") {
			log.Fatalf("viper.MergeInConfig returns error: %s", err.Error())
		}
	}
	viper.AutomaticEnv()
}

type AppConfig struct {
	Env string `json:"env"`
}

func GetAppConfig() AppConfig {
	return AppConfig{
		Env: viper.GetString(".env"),
	}
}

type ElasticApmConfig struct {
	ServiceName    string `json:"elastic_apm_service_name"`
	ServiceVersion string `json:"elastic_apm_service_version"`
}

func GetElasticApmConfig() ElasticApmConfig {
	return ElasticApmConfig{
		ServiceName:    viper.GetString("elastic_apm_service_name"),
		ServiceVersion: viper.GetString("elastic_apm_service_version"),
	}
}

type HTTPConfig struct {
	Port uint64 `json:"port"`
}

func GetHTTPConfig() HTTPConfig {
	return HTTPConfig{
		Port: viper.GetUint64("http_port"),
	}
}

type LoggingFormatter string

var (
	LoggingJSONFormatter LoggingFormatter = "json"
	LoggingTextFormatter LoggingFormatter = "text"
)

type LoggingLevel string

var (
	LoggingInfoLevel  LoggingLevel = "info"
	LoggingDebugLevel LoggingLevel = "debug"
)

type LoggingConfig struct {
	Formatter LoggingFormatter `json:"formatter"`
	Level     LoggingLevel     `json:"level"`
}

func GetLoggingConfig() LoggingConfig {
	env := GetAppConfig().Env

	switch env {
	case "production":
		return LoggingConfig{
			Formatter: LoggingJSONFormatter,
			Level:     LoggingInfoLevel,
		}

	case "staging":
		return LoggingConfig{
			Formatter: LoggingJSONFormatter,
			Level:     LoggingDebugLevel,
		}

	default:
		return LoggingConfig{
			Formatter: LoggingTextFormatter,
			Level:     LoggingDebugLevel,
		}
	}
}

type DatabaseConfig struct {
	Dsn string
}

func GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Dsn: viper.GetString("database"),
	}
}

func AsMap() map[string]interface{} {
	return map[string]interface{}{
		"app":     GetAppConfig(),
		"http":    GetHTTPConfig(),
		"logging": GetLoggingConfig(),
	}
}
