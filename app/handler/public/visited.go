package public

import (
	"context"
	"net/http"

	"github.com/hiroshijp/hcce-observer/domain"
	"github.com/labstack/echo/v4"
)

type HistoryUsecase interface {
	Fetch(ctx context.Context, num int) (res []domain.History, err error)
	Store(ctx context.Context, history *domain.History) (err error)
	FetchWithTx(ctx context.Context, num int) (res []domain.History, err error)
}

type VisitedHandler struct {
	historyUsecase HistoryUsecase
}

func NewVisitedHandler(e *echo.Echo, hu HistoryUsecase) {
	handler := &VisitedHandler{
		historyUsecase: hu,
	}
	e.POST("/visited", handler.Visited)

}

func (h *VisitedHandler) Visited(c echo.Context) error {
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
