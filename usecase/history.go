package usecase

import (
	"context"
	"log"

	"github.com/hiroshijp/try-clean-arch/domain"
	"golang.org/x/sync/errgroup"
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

func (hu *HistoryUsecase) Fetch(ctx context.Context, num int) (res []domain.History, err error) {
	res, err = hu.historyRepo.Fetch(ctx, num)
	if err != nil {
		return nil, err
	}

	// Get visitor by visitor id
	for i, h := range res {
		visitor, err := hu.visitorRepo.GetByID(ctx, h.Visitor.ID)
		if err != nil {
			return nil, err
		}
		res[i].Visitor = visitor
	}

	// res, err = hu.fillVisitorDetails(ctx, res)
	// if err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }

	return
}

func (hu *HistoryUsecase) fillVisitorDetails(ctx context.Context, data []domain.History) ([]domain.History, error) {
	g, ctx := errgroup.WithContext(ctx)
	mapVisitors := map[int]domain.Visitor{}

	for _, history := range data {
		mapVisitors[history.Visitor.ID] = domain.Visitor{}
	}

	chanVisitors := make(chan domain.Visitor)
	for visitorID := range mapVisitors {
		visitorID := visitorID
		g.Go(func() error {
			res, err := hu.visitorRepo.GetByID(ctx, visitorID)
			if err != nil {
				return err
			}
			chanVisitors <- res
			return nil
		})
	}

	go func() {
		defer close(chanVisitors)
		err := g.Wait()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	for visitor := range chanVisitors {
		if visitor != (domain.Visitor{}) {
			mapVisitors[visitor.ID] = visitor
		}
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	for i, item := range data {
		if visitor, ok := mapVisitors[item.Visitor.ID]; ok {
			data[i].Visitor = visitor
		}
	}
	return data, nil
}
