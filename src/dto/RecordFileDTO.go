package dto

import (
	"streamax-go/com"
	"time"
)

type RecordFileDTO struct {
	AT            int
	ERRORCAUSE    string
	ERRORCODE     int
	FILETYPE      int
	LASTRECORD    int
	LOCK          int
	RECORD        string
	RECORDCHANNEL int
	RECORDID      string
	RECORDSIZE    int64
	SENDFILECOUNT int
	SENDTIME      int
	SERIAL        int
	STAMPID       int
	STREAMTYPE    int
}

//"20210401080429-20210401081003"
func (f RecordFileDTO) GetStartTime() time.Time {
	st := f.RECORD[:14]
	return com.N9MTimeToTime(st)
}

//"20210401080429-20210401081003"
func (f RecordFileDTO) GetEndTime() time.Time {
	et := f.RECORD[15:]
	return com.N9MTimeToTime(et)
}
