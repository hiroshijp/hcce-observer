package public

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/hiroshijp/hcce-observer/handler/middleware"

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
	name, password, ok := c.Request().BasicAuth()
	if !ok {
		return c.JSON(http.StatusBadRequest, "invalid basic auth")
	}

	ctx := c.Request().Context()
	if err := h.userUsecase.Signin(ctx, name, password); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnauthorized, "failed to signin")
	}

	isAdmin := false
	if name == os.Getenv("ADMIN_NAME") && password == os.Getenv("ADMIN_PASS") {
		isAdmin = true
	}

	t, err := middleware.CreateToken(name, isAdmin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to create token")
	}
	return c.JSON(http.StatusOK, echo.Map{"token": t})
}
