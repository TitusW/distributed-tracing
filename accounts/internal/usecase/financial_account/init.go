package financialaccount

import (
	"context"

	"github.com/TitusW/accounts/internal/entity"
	unitofwork "github.com/TitusW/unit-of-work"
)

type ResourceItf interface {
	GetFinancialAccount(ctx context.Context, ksuid string) (entity.FinancialAccount, error)
	UpdateFinancialAccount(ctx context.Context, ksuid string, amount float64) (entity.FinancialAccount, error)
}

type Usecase struct {
	repo ResourceItf
	uow  unitofwork.UnitOfWorkItf
}

func New(repo ResourceItf, uow unitofwork.UnitOfWorkItf) Usecase {
	return Usecase{
		repo: repo,
		uow:  uow,
	}
}
