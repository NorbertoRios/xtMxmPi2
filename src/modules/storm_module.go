package modules

import (
	"encoding/json"
	"fmt"
	"streamax-go/interfaces"
)

type Storm struct {
	GeneralPackageHeader `json:"-"`
	MODULE               string //STORM
	OPERATION            string
	PARAMETER            *interface{} `json:",omitempty"`
	SESSION              string
	RESPONSE             *interface{} `json:",omitempty"`
}

type GetCalendarRequestParameter struct {
	CALENDARTYPE int
	STREAMTYPE   int
	FILETYPE     int
	CHANNEL      int
	QUERYTIME    string
}

func OperationGetCalendarRequest(c interfaces.IChannel, qMonth string, session string) {
	s := &Storm{
		GeneralPackageHeader: GeneralPackageHeader{},
		MODULE:               "STORM",
		OPERATION:            "GETCALENDAR",
		SESSION:              session,
		RESPONSE:             nil,
	}
	var crp interface{} = &GetCalendarRequestParameter{
		CALENDARTYPE: 1,
		STREAMTYPE:   1,
		FILETYPE:     127,    //select all types
		CHANNEL:      4095,   //select all channels
		QUERYTIME:    qMonth, //yyyymm
	}
	s.PARAMETER = &crp
	marshal, _ := json.Marshal(s)
	c.SendBytes(append(s.GeneralPackageHeader.toHeaderBytes(uint(len(marshal))), marshal...))
}

func (s Storm) HandleRequest(c interfaces.IChannel, buffer []byte) {
	fmt.Println(string(buffer))
	var m map[string]interface{}
	json.Unmarshal(s.PayloadBody, &m)
	rm := m["RESPONSE"].(map[string]interface{})
	switch m["OPERATION"] {
	case "GETCALENDAR":
		marshal, _ := json.Marshal(rm)
		s.OperationGetCalendarResponse(marshal)
	case "QUERYFILELIST":
		marshal, _ := json.Marshal(rm)
		s.OperationQueryFileListResponse(marshal, c)
	}
}

func (s Storm) ParseDtoFromData(buffer []byte) interface{} {
	s.FillGeneralPackageHeaderFromPackage(buffer)
	return s
}

func (s Storm) OperationGetCalendarResponse(payload []byte) {
	var res *GetCalendarResponse
	json.Unmarshal(payload, res)
}

func (s Storm) OperationQueryFileListResponse(payload []byte, c interfaces.IChannel) {
	var res *QueryFileListResponse
	err := json.Unmarshal(payload, &res)
	if err == nil && res != nil {
		//service.HandleFileListResponse(res, c)
	}
}

type GetCalendarResponse struct {
	CALENDER   []string
	CHANNEL    int
	CHCALENDER []map[string]interface{}
	COUNT      int
	ERRORCAUSE string
	ERRORCODE  int
}

func OperationQueryFileListRequest(c interfaces.IChannel, session string, startTime string, endTime string, chMask int) {
	s := &Storm{
		GeneralPackageHeader: GeneralPackageHeader{},
		MODULE:               "STORM",
		OPERATION:            "QUERYFILELIST",
		SESSION:              session,
		RESPONSE:             nil,
	}
	var qfl interface{} = &QueryFileListRequestParameter{
		SERIAL:     95255,
		STARTTIME:  startTime,
		CHANNEL:    chMask,
		ENDTIME:    endTime,
		STREAMTYPE: 256,
		FILETYPE:   65535,
		RFSTORAGE:  0,
	}
	s.PARAMETER = &qfl
	marshal, _ := json.Marshal(s)
	c.SendBytes(append(s.GeneralPackageHeader.toHeaderBytes(uint(len(marshal))), marshal...))
}

type QueryFileListRequestParameter struct {
	SERIAL     uint //unique request identifier
	STARTTIME  string
	CHANNEL    int
	ENDTIME    string
	STREAMTYPE int
	FILETYPE   int
	RFSTORAGE  int // 0 hard drive, 1 sd card
}

type QueryFileListResponse struct {
	ERRORCAUSE string
	ERRORCODE  int
	SERIAL     uint //unique request identifier
	//The alarm type, each element bit reprensent, corresponds to RECORD one by one.
	//BIT0 = 0, // IO alarm 1
	//BIT1 = 1, // IO alarm 2
	//BIT2 = 2, // IO alarm 3
	//BIT3 = 3, // IO alarm 4
	//BIT4 = 4, // IO alarm 5
	//BIT5 = 5, // IO alarm 6
	//BIT6 = 6, // IO alarm 7
	//BIT7 = 7, // IO alarm 8
	//BIT8 = 8, // panel alarm (emergency alarm (robbery))
	//BIT9 = 9, // speed alarm
	//BIT10 = 10, // video is missing
	//BIT11 = 11, // motion detection
	//BIT12 = 12, // video occlusion
	//BIT13 = 13, // gesture alarm
	//BIT19 = 19, // network is issued
	//BIT20 = 20, // electronic fence
	//BIT21 = 21, // ACC alarm
	//BIT22 = 22, / / ​​reported station alarm, including stagnation and station
	//BIT23 = 23, / / ​​peripheral dropped alarm
	//BIT24 = 24, // rollover alarm
	//BIT25 = 25, // antenna abnormal alarm
	//BIT26 = 26, // timeout stop
	//BIT27 = 27,
	//BIT28 = 28,
	//BIT29 = 29,
	//BIT30 = 30,
	//BIT31 = 31,
	//BIT32 = 32,
	//BIT33 = 33,
	//BIT34 = 34,
	//BIT35 = 35,
	//BIT36 = 36,
	//BIT37 = 37,
	//BIT38 = 38,
	//BIT39 = 39, // dedicated customer class alarm
	//BIT40 = 40, // serial key
	//BIT41 = 41, // hard disk button
	//BIT42 = 42, // vehicle collision
	//BIT43 = 43, // door open
	//BIT44 = 44, // door off
	//BIT45 = 45, // the door is closed by the opening, the speed from less than 20 to more than 20
	//BIT46 = 46, // server evaluation
	//BIT47 = 47, // taxi business, began to carry passengers
	//BIT48 = 48, // Pre alarm
	AT            []int //RECORD arr index,
	SENDFILECOUNT int   //files in this payload
	FILETYPE      []int
	LASTRECORD    int   //1 for last response of request, 0 other
	LOCK          []int //array index is for RECORD, 1 is for locked, 0 not
	//1-30 record time quantum
	//(20110928090909-20110929101010: stand for to record from 09:09:09 2011-09-28 to 10:10:10 2011-09-29, the research can step over day)
	//Picture name (20110928090909.JPG), transaction record, the number of element is  SENDFILECOUNT
	RECORD        []string //actual names of files
	RECORDCHANNEL []int    // num of channel, array index matched to RECORD array index
	RECORDID      []string //some string order counter ex: "0-0-2"
	RECORDSIZE    []uint   //RECORD index, bytes size
	SENDTIME      int      // counter for payloads started from 1
	STAMPID       []int    //creation order count, not unique
	STREAMTYPE    []int    //no docs
}
