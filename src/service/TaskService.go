package service

import (
	"entity"
	"time"
)

func CreateTask(deviceDsno string, startTime time.Time, endTime time.Time, channels int, stream int, subStream int, screenshot int) error {
	var dev *entity.Devices
	tx := DB.Begin()
	err := tx.First(&dev, "dsno = ?", deviceDsno).Error
	if err != nil || dev == nil {
		dev = &entity.Devices{
			Dsno: deviceDsno,
		}
		tx.Create(dev)
	}
	tx.Create(&entity.Tasks{
		Status:     "CREATED",
		DeviceId:   dev.ID,
		StartTime:  startTime,
		EndTime:    endTime,
		Channels:   channels,
		Stream:     stream,
		SubStream:  subStream,
		Screenshot: screenshot,
	})
	return tx.Commit().Error
}

func PutFlagToInt(flagPos int, val int) int {
	var argument int = 1
	if flagPos < 1 {
		return val
	}
	if flagPos > 1 {
		argument = argument << (flagPos - 1)
	}
	return val | argument
}
