package commonService

import (
	"fmt"
	"gorm.io/gorm"
)

func RecoverTX(tx *gorm.DB) {
	if r := recover(); r != nil {
		tx.Rollback()
		fmt.Println(r)
	}
}

func NewTx(tx *gorm.DB) *gorm.DB {
	newTx := tx.Begin()
	RecoverTX(newTx)
	return newTx
}
