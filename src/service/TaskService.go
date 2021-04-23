package service

import (
	"entity"
	"time"
)

func CreateTask(deviceDsno string, startTime time.Time, endTime time.Time, channels int, stream int, subStream int, screenshot int) {
	var dev *entity.Devices
	err := DB.First(&dev, "dsno = ?", deviceDsno).Error
	if err != nil {
		dev = &entity.Devices{
			Dsno: deviceDsno,
		}
		DB.Create(dev)
	}
	DB.Create(&entity.Tasks{
		DeviceId:   dev.ID,
		StartTime:  startTime,
		EndTime:    endTime,
		Channels:   channels,
		Stream:     stream,
		SubStream:  subStream,
		Screenshot: screenshot,
	})
}
