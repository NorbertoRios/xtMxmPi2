package service

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
