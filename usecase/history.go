package usecase

import (
	"context"
	"log"

	"github.com/hiroshijp/try-clean-arch/domain"
)

type HistoryRepository interface {
	Fetch(ctx context.Context, num int) (res []domain.History, err error)
	Store(ctx context.Context, history *domain.History) error
}

type VisitorRepository interface {
	GetByID(ctx context.Context, id int) (res domain.Visitor, err error)
	GetByMail(ctx context.Context, mail string) (res domain.Visitor, err error)
	Store(ctx context.Context, visitor *domain.Visitor) error
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
	return
}

func (hu *HistoryUsecase) Store(ctx context.Context, history *domain.History) (err error) {
	// visitor := history.Visitor
	// err = hu.storeVisitorIfNotExist(ctx, &visitor)
	// if err != nil {
	// 	return err
	// }
	// history.Visitor = visitor
	// return hu.historyRepo.Store(ctx, history)

	visitor, _ := hu.visitorRepo.GetByMail(ctx, history.Visitor.Mail)
	log.Println(visitor)
	if visitor == (domain.Visitor{}) {
		err = hu.visitorRepo.Store(ctx, &visitor)
		log.Println(visitor)
		if err != nil {
			return err
		}
	}
	history.Visitor = visitor
	return hu.historyRepo.Store(ctx, history)
}

// func (hu *HistoryUsecase) storeVisitorIfNotExist(ctx context.Context, visitor *domain.Visitor) error {
// 	existedVisitor, _ := hu.visitorRepo.GetByMail(ctx, visitor.Mail)
// 	if existedVisitor != (domain.Visitor{}) {
// 		return nil
// 	}
// 	return hu.visitorRepo.Store(ctx, visitor)
// }
