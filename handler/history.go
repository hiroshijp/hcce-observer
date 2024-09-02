package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/hiroshijp/try-clean-arch/domain"
	"github.com/labstack/echo/v4"
)

type HistoryUsecase interface {
	Fetch(ctx context.Context, num int) (res []domain.History, err error)
}

type HistoryHandler struct {
	historyUsecase HistoryUsecase
}

func NewHistoryHandler(e *echo.Echo, hu HistoryUsecase) {
	handler := &HistoryHandler{
		historyUsecase: hu,
	}
	e.GET("/history", handler.Fetch)
}

func (h HistoryHandler) Fetch(c echo.Context) error {
	numS := c.QueryParam("num")
	num, err := strconv.Atoi(numS)
	if err != nil {
		num = 10
	}

	ctx := c.Request().Context()
	data, err := h.historyUsecase.Fetch(ctx, num)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}
	return c.JSON(http.StatusOK, data)
}
