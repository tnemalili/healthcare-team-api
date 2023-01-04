package controllers

import (
	"billing.api/core/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthCheck(ctx echo.Context) error {

	message := &models.GenericMessage{Message: "AURA Billing API Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}
