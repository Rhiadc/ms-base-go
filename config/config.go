package config

import (
	"fmt"

	"github.com/pingcap/log"
	"github.com/spf13/viper"
)

type Config struct {
	APIPort           string
	ENV               string
	ServiceName       string
	DBConnectionHost  string
	MigrationLocation string
	MigrationsEnabled bool
	Log               Log
}

type Log struct {
	DevMode  bool
	Encoding string
}

func LoadEnvVars() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.SetDefault("ENV", "dev")
	viper.SetDefault("API_PORT", "8080")
	viper.SetDefault("DB_CONNECTION_HOST", "postgres://postgres:postgres@localhost:5432")
	viper.SetDefault("MigrationLocation", "./migrations")
	viper.SetDefault("MigrationsEnabled", true)

	viper.SetDefault("DEV_MODE", true)
	viper.SetDefault("ENCODING", "json")

	if err := viper.ReadInConfig(); err != nil {
		log.Info(fmt.Sprintf("unable to find or read config file: %s", err))
	}

	return &Config{
		APIPort:     viper.GetString("API_PORT"),
		ENV:         viper.GetString("ENV"),
		ServiceName: viper.GetString("SERVICE_NAME"),
		Log: Log{
			DevMode:  viper.GetBool("DEV_MODE"),
			Encoding: viper.GetString("ENCODING"),
		},
		DBConnectionHost: viper.GetString("DB_CONNECTION_HOST"),
	}
}
