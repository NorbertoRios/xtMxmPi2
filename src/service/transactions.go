package service

import "gorm.io/gorm"

func RecoverTX(tx *gorm.DB) {
	if r := recover(); r != nil {
		tx.Rollback()
	}
}
