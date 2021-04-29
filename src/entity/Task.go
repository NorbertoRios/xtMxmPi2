package entity

import "time"

type Tasks struct {
	ID         int64 `gorm:"autoIncrement"`
	DeviceId   int64
	Status     string
	StartTime  *time.Time
	EndTime    *time.Time
	Channels   int
	Stream     int
	SubStream  int
	Screenshot int
	CreatedAt  *time.Time `gorm:"column:created_time"`
	UpdatedAt  *time.Time `gorm:"column:updated_time"`
	DeletedAt  *time.Time `gorm:"column:deleted_time"`
}
