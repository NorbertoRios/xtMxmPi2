package entity

import (
	"time"
)

type SubTasks struct {
	ID        int64 `gorm:"autoIncrement"`
	TaskId    int64
	Channel   int
	DataType  int
	Status    string
	DeviceId  int64
	StartTime *time.Time
	EndTime   *time.Time
	CreatedAt *time.Time `gorm:"column:created_time"`
	UpdatedAt *time.Time `gorm:"column:updated_time"`
	DeletedAt *time.Time `gorm:"column:deleted_time"`
}
