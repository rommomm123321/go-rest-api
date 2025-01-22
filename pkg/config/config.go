package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB     DBConfig
	Server ServerConfig
}

type DBConfig struct {
	DSN string
}

type ServerConfig struct {
	Port string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Config{
		DB: DBConfig{
			DSN: viper.GetString("DB_DSN"),
		},
		Server: ServerConfig{
			Port: viper.GetString("PORT"),
		},
	}, nil
}
