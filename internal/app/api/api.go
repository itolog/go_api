package api

import "github.com/sirupsen/logrus"

type API struct {
	config *Config
	logger *logrus.Logger
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
	}
}

func (api *API) Start() error {
	if err := api.configureLoggerField(); err != nil {
		return err
	}

	api.logger.Info("server started on port", api.config.Port)

	return nil
}
