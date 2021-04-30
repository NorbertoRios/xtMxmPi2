package service

import (
	"gorm.io/gorm"
	"streamax-go/entity"
	"streamax-go/httpDto"
	"time"
)

func CreateTask(deviceDsno string, startTime time.Time, endTime time.Time, channels int, stream int, subStream int, screenshot int) (httpDto.TaskResponse, error) {
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
	var count int64
	err = tx.Table("task_queue").
		Where("device_id = ?", dev.ID).
		Order("created_time").
		Count(&count).
		Error
	if err != nil {
		tx.Rollback()
		return httpDto.TaskResponse{}, err
	}
	return httpDto.TaskResponse{
		Id:       task.ID,
		QueuedAt: count,
	}, tx.Commit().Error
}
