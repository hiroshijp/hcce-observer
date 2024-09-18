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
	Store(ctx context.Context, history *domain.History) (err error)
	FetchWithTx(ctx context.Context, num int) (res []domain.History, err error)
}

type HistoryHandler struct {
	historyUsecase HistoryUsecase
}

func NewHistoryHandler(e *echo.Group, hu HistoryUsecase) {
	handler := &HistoryHandler{
		historyUsecase: hu,
	}
	e.GET("/history", handler.Fetch)
	e.POST("/history", handler.Store)
	e.GET("/history/tx", handler.FetchWithTx)

}

func (h *HistoryHandler) Fetch(c echo.Context) error {
	numS := c.QueryParam("num")
	num, err := strconv.Atoi(numS)
	if err != nil {
		num = 10
	}

	ctx := c.Request().Context()
	data, err := h.historyUsecase.Fetch(ctx, num)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func (h *HistoryHandler) Store(c echo.Context) error {
	var history domain.History
	err := c.Bind(&history)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = h.historyUsecase.Store(ctx, &history)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, history)
}

func (h *HistoryHandler) FetchWithTx(c echo.Context) error {
	numS := c.QueryParam("num")
	num, err := strconv.Atoi(numS)
	if err != nil {
		num = 10
	}

	ctx := c.Request().Context()
	data, err := h.historyUsecase.FetchWithTx(ctx, num)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}
