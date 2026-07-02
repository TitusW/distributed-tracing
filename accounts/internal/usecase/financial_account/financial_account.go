package financialaccount

import (
	"context"

	"github.com/TitusW/accounts/internal/entity"
)

func (uc Usecase) DebitFinancialAccount(ctx context.Context, debitFinancialAccount entity.DebitFinancialAccount) (entity.FinancialAccount, error) {
	financialAccount, err := uc.repo.GetFinancialAccount(ctx, debitFinancialAccount.Ksuid)
	if err != nil {
		return entity.FinancialAccount{}, err
	}

	currentAmount := financialAccount.CurrentAmount - debitFinancialAccount.Amount
	if currentAmount < 0 {
		return financialAccount, err
	}

	updatedFinancialAccount, err := uc.repo.UpdateFinancialAccount(ctx, financialAccount.Ksuid, currentAmount)
	if err != nil {
		return entity.FinancialAccount{}, err
	}

	return updatedFinancialAccount, nil
}

func (uc Usecase) CreditFinancialAccount(ctx context.Context, creditFinancialAccount entity.CreditFinancialAccount) (entity.FinancialAccount, error) {
	financialAccount, err := uc.repo.GetFinancialAccount(ctx, creditFinancialAccount.Ksuid)
	if err != nil {
		return entity.FinancialAccount{}, err
	}

	currentAmount := financialAccount.CurrentAmount + creditFinancialAccount.Amount
	updatedFinancialAccount, err := uc.repo.UpdateFinancialAccount(ctx, financialAccount.Ksuid, currentAmount)
	if err != nil {
		return entity.FinancialAccount{}, err
	}

	return updatedFinancialAccount, nil
}
