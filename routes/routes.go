package routes

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"os"
	"suppliers.api/controllers"
)

func ConfigureRouter(router *echo.Echo) error {
	version := fmt.Sprintf("/api/%s/", os.Getenv("API_VERSION"))
	port := os.Getenv("API_PORT")
	router.Use(middleware.Recover())
	//router.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
	//	KeyLookup: "header:X-API-KEY",
	//	Validator: core.APIKEYValidator,
	//}))
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{AllowOrigins: []string{"*"}}))

	// DEBUG MODE (OPTIONAL)
	//router.Debug = true
	// JUST TO HIDE THE ECHO BANNER
	router.HideBanner = true
	api := router.Group(version)

	///////////////////// HEALTH SECTION //////////////////////////////////////

	api.GET("health", controllers.HealthCheck)

	///////////////////////////// AUTH SECTION ///////////////////////////////

	api.POST("suppliers", controllers.CreateSupplier)
	api.POST("responer", controllers.CreateSupplier)

	data, err := json.MarshalIndent(router.Routes(), "", "  ")
	if err != nil {
		return err
	}

	// LOG ALL REGISTERED ROUTES
	log.Info(string(data))
	log.Info("[AURA.CORE.SUPPLIERS.API - STARTED!]")
	err = router.Start(fmt.Sprintf(":%s", port))
	return err
}
