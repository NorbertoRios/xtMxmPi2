package repository

import (
	"gorm.io/gorm"
	"streamax-go/entity"
)

func GetDeviceByDSNO(dsno string, tx *gorm.DB) (int64, error) {
	var dev entity.Devices
	err := tx.First(&dev, "dsno = ?", dsno).Error
	return dev.ID, err
}
