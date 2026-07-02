package transfers

import "gorm.io/gorm"

func (tr *TransferRepo) GetDB() *gorm.DB {
	return tr.db
}
