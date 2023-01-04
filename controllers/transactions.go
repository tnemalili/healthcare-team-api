package controllers

import (
	"billing.api/core/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetTransactions(ctx echo.Context) error {
	message := &models.GenericMessage{Message: "AURA Billing API Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}
