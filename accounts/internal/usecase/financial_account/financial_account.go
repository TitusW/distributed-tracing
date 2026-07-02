package financialaccount

import (
	"context"

	"github.com/TitusW/accounts/internal/entity"
)

func (uc Usecase) DebitFinancialAccount(ctx context.Context, debitFinancialAccount entity.DebitFinancialAccount) (entity.FinancialAccount, error) {
	var updatedFinancialAccount entity.FinancialAccount

	uc.uow.Do(ctx, func(txCtx context.Context) error {
		financialAccount, err := uc.repo.GetFinancialAccount(txCtx, debitFinancialAccount.Ksuid)
		if err != nil {
			return err
		}

		currentAmount := financialAccount.CurrentAmount - debitFinancialAccount.Amount
		if currentAmount < 0 {
			return err
		}

		updatedFinancialAccount, err = uc.repo.UpdateFinancialAccount(txCtx, financialAccount.Ksuid, currentAmount)
		if err != nil {
			return err
		}

		return nil
	})

	return updatedFinancialAccount, nil
}

func (uc Usecase) CreditFinancialAccount(ctx context.Context, creditFinancialAccount entity.CreditFinancialAccount) (entity.FinancialAccount, error) {
	var updatedFinancialAccount entity.FinancialAccount

	uc.uow.Do(ctx, func(txCtx context.Context) error {
		financialAccount, err := uc.repo.GetFinancialAccount(txCtx, creditFinancialAccount.Ksuid)
		if err != nil {
			return err
		}

		currentAmount := financialAccount.CurrentAmount + creditFinancialAccount.Amount
		updatedFinancialAccount, err = uc.repo.UpdateFinancialAccount(txCtx, financialAccount.Ksuid, currentAmount)
		if err != nil {
			return err
		}

		return nil
	})

	return updatedFinancialAccount, nil
}
