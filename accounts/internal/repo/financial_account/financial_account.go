package financialaccount

import (
	"context"
	"time"

	"github.com/TitusW/accounts/internal/entity"
)

func (m Module) GetFinancialAccount(ctx context.Context, ksuid string) (entity.FinancialAccount, error) {
	var row FinancialAccount

	err := m.db.WithContext(ctx).Where("ksuid = ?", ksuid).First(&row).Error
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
	financialAccount := FinancialAccount{
		Ksuid:         ksuid,
		CurrentAmount: currentAmount,
		UpdatedAt:     time.Now(),
	}

	err := m.db.WithContext(ctx).Where("ksuid = ?", financialAccount.Ksuid).Updates(financialAccount).Error
	if err != nil {
		return entity.FinancialAccount{}, err
	}

	resFinancialAccount := entity.FinancialAccount{
		Ksuid:         financialAccount.Ksuid,
		CurrentAmount: financialAccount.CurrentAmount,
	}

	return resFinancialAccount, nil
}
