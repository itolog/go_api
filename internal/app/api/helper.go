package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gofiber/fiber/v2"
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
	a.app.Get("/json", func(c *fiber.Ctx) error {
		jsonFile, err := ioutil.ReadFile("assets/fake.json")
		if err != nil {
			fmt.Println(err)
		}
		var books Books

		jerr := json.Unmarshal(jsonFile, &books)
		if jerr != nil {
			log.Println(jerr)
		}

		return c.JSON(books)
	})
}
