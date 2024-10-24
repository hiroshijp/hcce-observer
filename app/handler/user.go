package handler

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hiroshijp/try-clean-arch/domain"
	"github.com/hiroshijp/try-clean-arch/handler/middleware"
	"github.com/labstack/echo/v4"
)

type UserUsecase interface {
	Create(ctx context.Context, user *domain.User) (err error)
}

type UserHandler struct {
	userUsecase UserUsecase
}

func NewUserHandler(e *echo.Group, uu UserUsecase) {
	handler := &UserHandler{
		userUsecase: uu,
	}
	e.POST("/user", handler.Create)
}

// admin作成はreppositoryのCreateを呼び出すだけ
func (h *UserHandler) Create(c echo.Context) error {

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*middleware.JwtCustomClaims)
	isAdmin := claims.Admin

	if !isAdmin {
		return c.JSON(403, "you are not admin")
	}

	var user domain.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if user.Name == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, "name and password are required")
	}

	ctx := c.Request().Context()
	err = h.userUsecase.Create(ctx, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
