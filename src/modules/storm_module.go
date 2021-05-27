package modules

import (
	"encoding/json"
	"fmt"
	"streamax-go/dto"
	"streamax-go/interfaces"
	"streamax-go/tcpService"
)

type Storm struct {
	GeneralPackageHeader `json:"-"`
	MODULE               string //STORM
	OPERATION            string
	PARAMETER            *interface{} `json:",omitempty"`
	SESSION              string
	RESPONSE             *interface{} `json:",omitempty"`
}

func OperationGetCalendarRequest(c interfaces.IChannel, qMonth string, session string) {
	s := &Storm{
		GeneralPackageHeader: GeneralPackageHeader{},
		MODULE:               "STORM",
		OPERATION:            "GETCALENDAR",
		SESSION:              session,
		RESPONSE:             nil,
	}
	var crp interface{} = &dto.GetCalendarRequestParameter{
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
	var res *dto.GetCalendarResponse
	json.Unmarshal(payload, res)
}

func (s Storm) OperationQueryFileListResponse(payload []byte, c interfaces.IChannel) {
	var res *dto.QueryFileListResponse
	err := json.Unmarshal(payload, &res)
	if err == nil && res != nil {
		tcpService.HandleFileListResponse(res, c)
	}
}

func OperationQueryFileListRequest(c interfaces.IChannel, session string, startTime string, endTime string, chMask int) {
	s := &Storm{
		GeneralPackageHeader: GeneralPackageHeader{},
		MODULE:               "STORM",
		OPERATION:            "QUERYFILELIST",
		SESSION:              session,
		RESPONSE:             nil,
	}
	var qfl interface{} = &dto.QueryFileListRequestParameter{
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
