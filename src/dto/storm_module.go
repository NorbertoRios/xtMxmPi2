package dto

import (
	"comm/channel"
	"encoding/json"
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

func OperationGetCalendarRequest(c channel.IChannel, qMonth string, session string) {
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

func (s Storm) HandleRequest(c channel.IChannel, buffer []byte) {
	var m map[string]interface{}
	json.Unmarshal(s.PayloadBody, &m)
	rm := m["RESPONSE"].(map[string]interface{})
	switch m["OPERATION"] {
	case "GETCALENDAR":
		marshal, _ := json.Marshal(rm)
		s.OperationGetCalendarResponse(marshal)
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

type GetCalendarResponse struct {
	CALENDER   []string
	CHANNEL    int
	CHCALENDER []map[string]interface{}
	COUNT      int
	ERRORCAUSE string
	ERRORCODE  int
}
