package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"suppliers.api/core/models"
)

func CreateResponder(ctx echo.Context) error {
	message := &models.GenericMessage{Message: "AURA Create Responder API Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}

func FetchResponders(ctx echo.Context) error {
	message := &models.GenericMessage{Message: "AURA Fetch Responder API Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}

func FetchResponder(ctx echo.Context) error {
	message := &models.GenericMessage{Message: "AURA Fetch a Responder API  Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}

func UpdateResponder(ctx echo.Context) error {
	message := &models.GenericMessage{Message: "AURA Update a Responder API  Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}

func DeleteResponder(ctx echo.Context) error {
	message := &models.GenericMessage{Message: "AURA Delete a Responder API  Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}
