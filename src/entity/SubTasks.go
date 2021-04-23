package entity

import (
	"gorm.io/gorm"
	"time"
)

type SubTasks struct {
	ID        int64 `gorm:"autoIncrement"`
	TaskId    int64
	Channel   int
	DataType  int
	Status    string
	DeviceId  int64
	StartTime time.Time
	EndTime   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
