package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Books []Book

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

	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		jsonFile, err := ioutil.ReadFile("assets/fake.json")
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonFile)

	})
}
