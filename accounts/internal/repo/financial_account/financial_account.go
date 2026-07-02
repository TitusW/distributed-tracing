package financialaccount

import (
	"context"
	"time"

	"github.com/TitusW/accounts/internal/entity"
	unitofwork "github.com/TitusW/unit-of-work"
)

func (m Module) GetFinancialAccount(ctx context.Context, ksuid string) (entity.FinancialAccount, error) {
	tx := unitofwork.GetTX(ctx, m.db)

	var row FinancialAccount

	err := tx.WithContext(ctx).Where("ksuid = ?", ksuid).First(&row).Error
	if err != nil {
		return entity.FinancialAccount{}, err
	}

	resFinancialAccount := entity.FinancialAccount{
		Ksuid:         row.Ksuid,
		CurrentAmount: row.CurrentAmount,
	}

	return resFinancialAccount, nil
}

func (m Module) UpdateFinancialAccount(ctx context.Context, ksuid string, currentAmount float64) (entity.FinancialAccount, error) {
	tx := unitofwork.GetTX(ctx, m.db)

	financialAccount := FinancialAccount{
		Ksuid:         ksuid,
		CurrentAmount: currentAmount,
		UpdatedAt:     time.Now(),
	}

	err := tx.WithContext(ctx).Where("ksuid = ?", financialAccount.Ksuid).Updates(financialAccount).Error
	if err != nil {
		return entity.FinancialAccount{}, err
	}

	resFinancialAccount := entity.FinancialAccount{
		Ksuid:         financialAccount.Ksuid,
		CurrentAmount: financialAccount.CurrentAmount,
	}

	return resFinancialAccount, nil
}
