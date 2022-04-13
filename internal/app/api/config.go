package api

import "github.com/itolog/go_api/storage"

type Config struct {
	Port        string `toml:"port"`
	LoggerLevel string `toml:"logger_level"`
	Storage     *storage.Config
}

func NewConfig() *Config {
	return &Config{
		Port:        ":8080",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
