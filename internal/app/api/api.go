package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/itolog/go_api/storage"
	"github.com/sirupsen/logrus"
)

type API struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (api *API) Start() error {
	if err := api.configureLoggerField(); err != nil {
		return err
	}

	api.logger.Info("server started on port", api.config.Port)

	api.configureRouterField()

	if err := api.configureStorageField(); err != nil {
		return err
	}

	return http.ListenAndServe(api.config.Port, api.router)
}
