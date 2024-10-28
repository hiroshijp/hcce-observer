package usecase

import (
	"context"

	"github.com/hiroshijp/hcce-observer/domain"
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

type TxRepository interface {
	BeginTx(ctx context.Context, f func(ctc context.Context) (err error)) (err error)
}

type HistoryUsecase struct {
	txRepo      TxRepository
	historyRepo HistoryRepository
	visitorRepo VisitorRepository
}

func NewHistoryUsecase(tr TxRepository, hr HistoryRepository, vr VisitorRepository) *HistoryUsecase {
	return &HistoryUsecase{
		txRepo:      tr,
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
	existedVisitor, _ := hu.visitorRepo.GetByMail(ctx, history.Visitor.Mail)
	if existedVisitor == (domain.Visitor{}) {
		existedVisitor.Mail = history.Visitor.Mail
		err = hu.visitorRepo.Store(ctx, &existedVisitor)
		if err != nil {
			return err
		}
	}
	history.Visitor = existedVisitor
	return hu.historyRepo.Store(ctx, history)
}

func (hu *HistoryUsecase) FetchWithTx(ctx context.Context, num int) (res []domain.History, err error) {
	err = hu.txRepo.BeginTx(ctx, func(ctx context.Context) (err error) {
		res, err = hu.historyRepo.Fetch(ctx, num)
		if err != nil {
			return err
		}

		// Get visitor by visitor id
		for i, h := range res {
			visitor, err := hu.visitorRepo.GetByID(ctx, h.Visitor.ID)
			if err != nil {
				return err
			}
			res[i].Visitor = visitor
		}
		return nil
	})
	return
}
