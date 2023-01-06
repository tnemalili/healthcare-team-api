package main

import (
	"github.com/labstack/echo/v4"
	"suppliers.api/routes"

	log "github.com/sirupsen/logrus"
)

func main() {
	// SETTING LOGGER
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, ForceColors: true})
	router := echo.New()
	err := routes.ConfigureRouter(router)
	if err != nil {
		log.Errorf("[suppliers.api] Error: %v", err.Error())
	}
}
