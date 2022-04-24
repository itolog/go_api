package api

import (
	"github.com/gin-gonic/gin"
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
	a.app.GET("/json", func(ctx *gin.Context) {
		ctx.File("assets/fake.json")
	})
}
