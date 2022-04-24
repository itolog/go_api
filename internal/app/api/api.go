package api

import (
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

type API struct {
	config *Config
	logger *logrus.Logger
	app    *iris.Application
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		app:    iris.New(),
	}
}

func (api *API) Start() error {
	if err := api.configureLoggerField(); err != nil {
		return err
	}

	api.logger.Info("server started on port", api.config.Port)

	api.configureAppField()

	return api.app.Listen(api.config.Port)
}
