package httpService

import (
	"github.com/nu7hatch/gouuid"
	"gorm.io/gorm"
	"strconv"
	"streamax-go/commonService"
	"streamax-go/entity"
	"streamax-go/httpDto"
	"streamax-go/interfaces"
	"streamax-go/modules"
	"time"
)

func CreateTask(taskDto *httpDto.Task, deviceDsno string, startTime time.Time, endTime time.Time, channels int, stream int, subStream int, screenshot int) (httpDto.TaskResponse, error) {
	var dev *entity.Devices
	tx := DB.Begin()
	defer commonService.RecoverTX(tx)
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
	st, errSt := CreateSubTaskFromTask(taskDto, task, tx)
	if errSt != nil {
		tx.Rollback()
		return httpDto.TaskResponse{}, err
	}
	_, errSTQ := QueueSubTasks(st, task, tx)
	if errSTQ != nil {
		tx.Rollback()
		return httpDto.TaskResponse{}, err
	}
	return httpDto.TaskResponse{
		Id:       task.ID,
		QueuedAt: count,
	}, tx.Commit().Error
}

func CreateSubTaskFromTask(taskDto *httpDto.Task, task *entity.Tasks, tx *gorm.DB) ([]*entity.SubTasks, error) {
	sTasks := entity.CreateSubTasksFromTask(taskDto, task)
	tx.CreateInBatches(sTasks, len(sTasks))
	return sTasks, tx.Error
}

func QueueSubTasks(st []*entity.SubTasks, task *entity.Tasks, tx *gorm.DB) ([]*entity.SubTaskQueue, error) {
	q := make([]*entity.SubTaskQueue, 0)
	for i := 0; i < len(st); i++ {
		q = append(q, &entity.SubTaskQueue{
			SubTaskId: st[i].ID,
			TaskId:    task.ID,
			DeviceId:  task.DeviceId,
			Status:    "CREATED",
		})
	}
	return q, tx.CreateInBatches(q, len(q)).Error
}

func SetSubTaskToProgress(st entity.SubTasks, task entity.Tasks, tx *gorm.DB) (*entity.SubTaskQueue, error) {
	var stq *entity.SubTaskQueue
	err := tx.First(&stq, "subtask_id = ?", st.ID).Error

	if err == gorm.ErrRecordNotFound {
		stq = &entity.SubTaskQueue{
			SubTaskId: st.ID,
			TaskId:    task.ID,
			DeviceId:  task.DeviceId,
			Status:    commonService.Dispatching,
		}
	} else {
		stq.Status = commonService.Dispatching
	}
	tx.Save(stq)
	return stq, err
}

func ProcessSubTask(st entity.SubTasks, task entity.Tasks, tx *gorm.DB) (*entity.SubTaskQueue, error) {
	var device entity.Devices
	tx.First(device, "id = ?", st.DeviceId)
	//channel := scontext.DeviceChannelMap[device.Dsno]
	//dto.RequestFile(*channel, "58895_video", 3,
	//	"0-0-101", 1, "20210525075506", "20210525081417")
	return nil, nil
}

func processStream(c interfaces.IChannel, st entity.SubTasks) {
	session, _ := uuid.NewV4()
	modules.OperationQueryFileListRequest(c, session.String(),
		strconv.FormatInt(st.StartTime.Unix(), 10),
		strconv.FormatInt(st.EndTime.Unix(), 10),
		st.Channel)
}
