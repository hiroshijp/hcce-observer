package usecase

import (
	"context"
	"fmt"

	"github.com/hiroshijp/try-clean-arch/domain"
)

type HistoryRepository interface {
	Fetch(ctx context.Context, num int) (res []domain.History, err error)
}

type VisitorRepository interface {
	GetByID(ctx context.Context, id int) (res domain.Visitor, err error)
}

type HistoryUsecase struct {
	historyRepo HistoryRepository
	visitorRepo VisitorRepository
}

func NewHistoryUsecase(hr HistoryRepository, vr VisitorRepository) *HistoryUsecase {
	return &HistoryUsecase{
		historyRepo: hr,
		visitorRepo: vr,
	}
}

func (hs *HistoryUsecase) Fetch(ctx context.Context, num int) (res []domain.History, err error) {
	res, err = hs.historyRepo.Fetch(ctx, num)
	if err != nil {
		return nil, err
	}

	// Get visitor by visitor id
	for i, h := range res {
		v, err := hs.visitorRepo.GetByID(ctx, h.Visitor.ID)
		if err != nil {
			return nil, err
		}
		fmt.Println(v)
		res[i].Visitor = v
	}

	return res, nil
}
