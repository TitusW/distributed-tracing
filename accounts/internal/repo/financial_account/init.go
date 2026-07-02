package financialaccount

import "gorm.io/gorm"

type Module struct {
	db *gorm.DB
}

func New(db *gorm.DB) Module {
	return Module{
		db: db,
	}
}
