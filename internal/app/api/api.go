package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type API struct {
	config *Config
	logger *logrus.Logger
	app    *fiber.App
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		app:    fiber.New(),
	}
}

type Books struct {
	Book []Book `json:"books"`
}

type Book struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (api *API) Start() error {
	if err := api.configureLoggerField(); err != nil {
		return err
	}

	api.configureAppField()

	return api.app.Listen(api.config.Port)
}
