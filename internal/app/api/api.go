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

func (api *API) Start() error {
	if err := api.configureLoggerField(); err != nil {
		return err
	}

	api.logger.Info("server started on port", api.config.Port)

	api.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	return api.app.Listen(api.config.Port)
}
