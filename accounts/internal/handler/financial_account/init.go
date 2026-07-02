package financialaccount

import (
	"context"

	"github.com/TitusW/accounts/internal/entity"
)

type UsecaseItf interface {
	DebitFinancialAccount(ctx context.Context, debitFinancialAccount entity.DebitFinancialAccount) (entity.FinancialAccount, error)
	CreditFinancialAccount(ctx context.Context, creditFinancialAccount entity.CreditFinancialAccount) (entity.FinancialAccount, error)
}

type Handler struct {
	uc UsecaseItf
}

func New(uc UsecaseItf) Handler {
	return Handler{
		uc: uc,
	}
}
