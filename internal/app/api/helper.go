package api

import (
	"fmt"
	"io/ioutil"

	"github.com/kataras/iris/v12"
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

func (a *API) configureAppField() {
	a.app.Get("/json", func(ctx iris.Context) {

		jsonFile, err := ioutil.ReadFile("assets/fake.json")
		if err != nil {
			fmt.Println(err)
		}
		options := iris.JSON{Indent: " ", Secure: true}

		ctx.JSON(jsonFile, options)
	})
}
