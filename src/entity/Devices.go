package entity

type Devices struct {
	ID   int64  `gorm:"autoIncrement"`
	Dsno string `gorm:"uniqueNullTime"`
}
