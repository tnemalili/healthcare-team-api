package controllers

import (
	"billing.api/core/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func MakePayment(ctx echo.Context) error {
	message := &models.GenericMessage{Message: "AURA Billing API OTP Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}
