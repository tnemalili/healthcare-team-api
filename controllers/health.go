package controllers

import (
	"net/http"
	"suppliers.api/core/models"

	"github.com/labstack/echo/v4"
)

func HealthCheck(ctx echo.Context) error {

	message := &models.GenericMessage{Message: "AURA Suppliers API Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}
