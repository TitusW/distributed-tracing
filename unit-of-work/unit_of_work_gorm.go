package unitofwork

import (
	"context"

	"gorm.io/gorm"
)

type ctxKey string

const txCtx ctxKey = "tx"

type gormUnitOfWork struct {
	db *gorm.DB
}

type UnitOfWorkItf interface {
	Do(ctx context.Context, fn func(txCtx context.Context) error) error
}

func NewGormUnitOfWork(db *gorm.DB) (uow UnitOfWorkItf) {
	return gormUnitOfWork{db: db}
}

func (uow gormUnitOfWork) Do(ctx context.Context, fn func(context.Context) error) error {
	return uow.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txCtx := context.WithValue(ctx, txCtx, tx)

		return fn(txCtx)
	})
}

func GetTX(ctx context.Context, defaultDB *gorm.DB) *gorm.DB {
	if tx, ok := ctx.Value(txCtx).(*gorm.DB); ok {
		return tx
	}

	return defaultDB
}
