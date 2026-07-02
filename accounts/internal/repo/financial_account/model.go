package financialaccount

import (
	"time"

	"gorm.io/gorm"
)

type FinancialAccount struct {
	Ksuid         string
	CurrentAmount float64
	InsertedAt    time.Time
	UpdatedAt     time.Time
	DeletedAt     *gorm.DeletedAt
}
