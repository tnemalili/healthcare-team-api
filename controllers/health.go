package controllers

import (
	"auth.api/core/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthCheck(ctx echo.Context) error {

	message := &models.GenericMessage{Message: "AURA Auth API Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}
