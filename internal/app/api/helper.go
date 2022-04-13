package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/gorilla/mux"
	"github.com/itolog/go_api/storage"
	"github.com/sirupsen/logrus"
)

type Books struct {
	Books []Book `json:"books"`
}

type Book struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)

	return nil
}

func (a *API) configureRouterField() {

	a.router.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		jsonFile, err := ioutil.ReadFile("assets/fake.json")
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonFile)

	})
}

func (a *API) configureStorageField() error {
	storage := storage.New(a.config.Storage)
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
