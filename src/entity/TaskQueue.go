package entity

import "time"

type TaskQueue struct {
	SubtaskId int64     `gorm:"primaryKey;autoIncrement:false"`
	DeviceId  int64     `gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time `gorm:"column:created_time"`
}
