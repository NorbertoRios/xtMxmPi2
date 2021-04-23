package entity

import "time"

type Tasks struct {
	ID         int64 `gorm:"autoIncrement"`
	DeviceId   int64
	Status     string
	StartTime  time.Time
	EndTime    time.Time
	Channels   int
	Stream     int
	SubStream  int
	Screenshot int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
