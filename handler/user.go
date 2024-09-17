package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewUserHandler(e *echo.Echo) {
	e.GET("/user/login", Login)
}

func Login(c echo.Context) error {
	return c.JSON(http.StatusOK, "Login")
}
