package api

type Config struct {
	Port string `toml:"port"`
	LoggerLevel string `toml:"logger_level"`
}

func NewConfig() *Config {
	return &Config{
		Port: "8080",
		LoggerLevel: "debug",
	}
}