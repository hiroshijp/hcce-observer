package public

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserUsecase interface {
	Signin(ctx context.Context, name string, password string) (err error)
}

type SigninHandler struct {
	userUsecase UserUsecase
}

func NewSigninHandler(e *echo.Echo, uu UserUsecase) {
	handler := &SigninHandler{
		userUsecase: uu,
	}
	e.POST("/signin", handler.Signin)
}

func (h *SigninHandler) Signin(c echo.Context) error {
	// get id and password from request
	name, password, ok := c.Request().BasicAuth()
	if !ok {
		return c.String(http.StatusBadRequest, "invalid request \n")
	}

	ctx := c.Request().Context()
	if err := h.userUsecase.Signin(ctx, name, password); err != nil {
		fmt.Println(err)
		return c.String(http.StatusUnauthorized, "user not found\n")
	}
	return c.String(http.StatusOK, "signin\n")
}
