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
	panic("implement me")
}

func (s Storm) ParseDtoFromData(buffer []byte) interface{} {
	panic("implement me")
}
