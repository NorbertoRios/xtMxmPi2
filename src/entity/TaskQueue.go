package entity

import "time"

type TaskQueue struct {
	TaskId    int64      `gorm:"primaryKey;autoIncrement:false"`
	DeviceId  int64      `gorm:"primaryKey;autoIncrement:false"`
	CreatedAt *time.Time `gorm:"column:created_time"`
}

func (tq TaskQueue) TableName() string {
	return "task_queue"
}
