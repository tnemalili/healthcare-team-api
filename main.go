package main

import (
	"auth.api/routes"
	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

func main() {
	// SETTING LOGGER
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, ForceColors: true})
	router := echo.New()
	err := routes.ConfigureRouter(router)
	if err != nil {
		log.Errorf("[auth.api] Error: %v", err.Error())
	}
}
