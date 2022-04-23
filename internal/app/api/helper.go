package api

import (
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

		return c.SendFile("assets/fake.json", true)
	})
}
