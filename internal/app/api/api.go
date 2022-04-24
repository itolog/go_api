package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type API struct {
	config *Config
	logger *logrus.Logger
	app    *gin.Engine
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		app:    gin.Default(),
	}
}

func (api *API) Start() error {
	if err := api.configureLoggerField(); err != nil {
		return err
	}

	api.logger.Info("server started on port", api.config.Port)

	api.configureRouterField()

	api.app.Run()

	return nil
}
