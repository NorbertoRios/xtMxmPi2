package entity

import (
	"streamax-go/dto"
	"streamax-go/httpDto"
	"time"
)

type SubTasks struct {
	ID         int64 `gorm:"autoIncrement"`
	TaskId     int64
	Channel    int
	DataType   string
	Status     string
	DeviceId   int64
	InProgress *time.Time `gorm:"column:in_progress"`
	StartTime  *time.Time
	EndTime    *time.Time
	CreatedAt  *time.Time `gorm:"column:created_time"`
	UpdatedAt  *time.Time `gorm:"column:updated_time"`
	DeletedAt  *time.Time `gorm:"column:deleted_time"`
}

func CreateSubTasksFromTask(taskDto *httpDto.Task, task *Tasks) []*SubTasks {
	stream := "stream"
	subStream := "substream"
	screenshot := "screenshot"
	streamBits := taskDto.GetTypeFromTask(stream)
	subStreamBits := taskDto.GetTypeFromTask(subStream)
	screenshotBits := taskDto.GetTypeFromTask(screenshot)
	tasks := generateSubTasksFromInt(streamBits, stream, task)
	tasks = append(tasks, generateSubTasksFromInt(subStreamBits, subStream, task)...)
	tasks = append(tasks, generateSubTasksFromInt(screenshotBits, screenshot, task)...)
	return tasks
}

func generateSubTasksFromInt(mask int, dataType string, task *Tasks) []*SubTasks {
	subTasks := make([]*SubTasks, 0)
	t := time.Now().UTC()
	utc := &t
	for i := 31; i >= 0; i-- {
		if mask&(1<<i) > 0 {
			st := &SubTasks{
				TaskId:     task.ID,
				Channel:    i,
				DataType:   dataType,
				Status:     "CREATED",
				DeviceId:   task.DeviceId,
				InProgress: utc,
				StartTime:  task.StartTime,
				EndTime:    task.EndTime,
				CreatedAt:  task.CreatedAt,
				UpdatedAt:  task.UpdatedAt,
				DeletedAt:  task.DeletedAt,
			}
			subTasks = append(subTasks, st)
		}
	}
	return subTasks
}

func (st SubTasks) GetN9MStreamType() int {
	switch st.DataType {
	case dto.STREAM:
		return 1
	case dto.SUBSTREAM:
		return 0
	}
	return -1
}
