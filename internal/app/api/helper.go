package api

import (
	"net/http"

	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)

	return nil
}

func (a *API) configureRouterField() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("GET /")
		w.Write([]byte("Hello Word"))
	})
}
