package tcpService

import (
	"gorm.io/gorm"
	"strconv"
	"streamax-go/commonService"
	"streamax-go/dto"
	"streamax-go/entity"
	"streamax-go/interfaces"
	"streamax-go/l"
	"streamax-go/repository"
)

func HandleFileListResponse(res *dto.QueryFileListResponse, c interfaces.IChannel) {
	tx := commonService.NewTx(DB)
	var stq *entity.SubTaskQueue
	dsno, err2 := repository.GetDeviceByDSNO(c.GetDSNO(), tx)
	if err2 != nil {
		return
	}
	err := tx.First(&stq, "device_id = ? AND status = ?", dsno, commonService.Dispatching).Error
	if err == gorm.ErrRecordNotFound {
		l.Inf("no active tasks found for DSNO " + strconv.FormatInt(dsno, 10))
	} else if err != nil {
		l.Err("error in tx found for DSNO " + strconv.FormatInt(dsno, 10))
	}
	records := res.SplitToRecordFilesDTO()
	filtered := FilterRecordsByTaskConditions(stq, records)

	//tx.First()
}

func FilterRecordsByTaskConditions(stq *entity.SubTaskQueue, d []*dto.RecordFileDTO) []*dto.RecordFileDTO {
	res := make([]*dto.RecordFileDTO, len(d))
	var st *entity.SubTasks
	err := DB.First(&st, stq.SubTaskId).Error
	if err == gorm.ErrRecordNotFound {
		l.Inf("No record for subtask in FilterRecordsByTaskConditions subtask id:" + strconv.FormatInt(stq.SubTaskId, 10))
	} else if err != nil {
		l.Err("db error in FilterRecordsByTaskConditions  subtask id:" + strconv.FormatInt(stq.SubTaskId, 10))
	}
	for i := 0; i < len(d); i++ {
		if d[i].RECORDCHANNEL == st.Channel {
			startTime := st.StartTime
			endTime := st.EndTime
			if (startTime.Before(d[i].GetStartTime()) || startTime.Equal(d[i].GetStartTime())) &&
				(endTime.After(d[i].GetEndTime()) || endTime.Equal(d[i].GetEndTime())) {
				if d[i].STREAMTYPE == st.GetN9MStreamType() {
					res = append(res, d[i])
				}
			}
		}
	}
	return res
}
