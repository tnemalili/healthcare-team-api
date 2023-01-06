package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"suppliers.api/core/models"
)

func CreateSupplier(ctx echo.Context) error {
	message := &models.GenericMessage{Message: "AURA Create Supplier API Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}

func FetchSuppliers(ctx echo.Context) error {
	message := &models.GenericMessage{Message: "AURA Fetch Suppliers API Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}

func FetchSupplier(ctx echo.Context) error {
	message := &models.GenericMessage{Message: "AURA Fetch a Supplier API  Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}

func UpdateSupplier(ctx echo.Context) error {
	message := &models.GenericMessage{Message: "AURA Update a Supplier API  Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}

func DeleteSupplier(ctx echo.Context) error {
	message := &models.GenericMessage{Message: "AURA Delete a Supplier API  Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}
