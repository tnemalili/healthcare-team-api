package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"suppliers.api/core/models"
)

func CreateDevice(ctx echo.Context) error {
	message := &models.GenericMessage{Message: "AURA Create Device API Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}

func FetchDevices(ctx echo.Context) error {
	message := &models.GenericMessage{Message: "AURA Fetch Device API Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}

func FetchDevice(ctx echo.Context) error {
	message := &models.GenericMessage{Message: "AURA Fetch a Device API  Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}

func UpdateDevice(ctx echo.Context) error {
	message := &models.GenericMessage{Message: "AURA Update a Device API  Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}

func DeleteDevice(ctx echo.Context) error {
	message := &models.GenericMessage{Message: "AURA Delete a Device API  Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}
