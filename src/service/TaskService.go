package service

import (
	"gorm.io/gorm"
	"streamax-go/entity"
	"time"
)

func CreateTask(deviceDsno string, startTime time.Time, endTime time.Time, channels int, stream int, subStream int, screenshot int) (int64, int, error) {
	var dev *entity.Devices
	tx := DB.Begin()
	defer RecoverTX(tx)
	err := tx.First(&dev, "dsno = ?", deviceDsno).Error
	if err == gorm.ErrRecordNotFound || dev == nil {
		dev = &entity.Devices{
			Dsno: deviceDsno,
		}
		tx.Create(dev)
	}
	task := &entity.Tasks{
		Status:     "CREATED",
		DeviceId:   dev.ID,
		StartTime:  &startTime,
		EndTime:    &endTime,
		Channels:   channels,
		Stream:     stream,
		SubStream:  subStream,
		Screenshot: screenshot,
	}
	tx.Create(task)
	utc := time.Now().UTC()
	utcP := &utc
	tx.Create(&entity.TaskQueue{
		TaskId:    task.ID,
		DeviceId:  dev.ID,
		CreatedAt: utcP,
	})
	var taskQ = make([]entity.TaskQueue, 0)
	err = tx.Where("device_id = ?", dev.ID).Order("created_time").Find(taskQ).Error
	if err != nil {
		tx.Rollback()
		return 0, 0, err
	}
	qPos := 0
	if len(taskQ) > 0 {
		for i := 0; i < len(taskQ); i++ {
			if task.ID == taskQ[i].TaskId {
				qPos = i
			}
		}
	}
	return task.ID, qPos, tx.Commit().Error
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
