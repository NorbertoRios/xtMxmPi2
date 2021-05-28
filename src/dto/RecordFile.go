package dto

import "time"

type RecordFile struct {
	ID         int64 `gorm:"autoIncrement"`
	SubTaskId  int64 `gorm:"column:subtask_id"`
	Channel    int
	DataType   string
	Status     string
	DeviceId   int64
	AT         int
	RecordSize int
	StampId    int
	Cmd        int
	RecordId   string
	CreatedAt  *time.Time `gorm:"column:created_time"`
	UpdatedAt  *time.Time `gorm:"column:updated_time"`
	DeletedAt  *time.Time `gorm:"column:deleted_time"`
}
