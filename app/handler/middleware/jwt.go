package middleware

import "github.com/labstack/echo/v4"

func NewJWTMiddleware(e *echo.Group) {
	e.Use(JWT)
}

func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("auth", "jwt")
		return next(c)
	}
}
