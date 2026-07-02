package transfers

import (
	"gorm.io/gorm"
)

type TransferRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) TransferRepo {
	return TransferRepo{
		db: db,
	}
}
