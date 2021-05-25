package entity

import "time"

type SubTaskQueue struct {
	SubTaskId int64      `gorm:"primaryKey;autoIncrement:false;column:subtask_id"`
	TaskId    int64      `gorm:"primaryKey;autoIncrement:false"`
	DeviceId  int64      `gorm:"column:device_id"`
	Status    string     `gorm:"column:status"`
	CreatedAt *time.Time `gorm:"column:created_time"`
}

func (stq SubTaskQueue) TableName() string {
	return "subtask_queue"
}
