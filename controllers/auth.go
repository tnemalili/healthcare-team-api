package controllers

import (
	"auth.api/core/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RequestOTP(ctx echo.Context) error {
	message := &models.GenericMessage{Message: "AURA Auth API OTP Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}

func RequestAToken(ctx echo.Context) error {
	message := &models.GenericMessage{Message: "AURA Auth API Token Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}
