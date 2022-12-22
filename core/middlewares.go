package core

import (
	"os"

	"github.com/labstack/echo/v4"
)

func APIKEYValidator(key string, ctx echo.Context) (bool, error) {
	authorized := key == os.Getenv("X_API_KEY")
	return authorized, nil
}
