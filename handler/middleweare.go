package handler

import "github.com/labstack/echo/v4"

func NewMiddleware(e *echo.Echo) {
	e.Use(CORS)
}

func CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}
