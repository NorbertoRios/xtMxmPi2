package dto

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"streamax-go/config"
	"streamax-go/interfaces"
)

type MediaStreamModel struct {
	GeneralPackageHeader `json:"-"`
	MODULE               string
	OPERATION            string
	PARAMETER            interface{} //MEDIATASKSTART CREATESTREAM REQUESTSTREAM CONTROLSTREAM MEDIAREGISTEFAILACK MEDIATASKSTOP
	SESSION              string
	RESPONSE             interface{} `json:",omitempty"`
}

func (m MediaStreamModel) HandleRequest(c interfaces.IChannel, buffer []byte) {
	switch m.OPERATION {
	case "REQUESTDOWNLOADVIDEO":
		m.handleRequestDownloadVideoResponse(c, buffer)
	case "MEDIATASKSTART":
		m.handleMediaTaskStartResponse(c, buffer)
	case "MEDIATASKSTOP":
		m.handleMediaTaskStopResponse(c, buffer)
	}
}

func (m MediaStreamModel) ParseDtoFromData(buffer []byte) interface{} {
	var result MediaStreamModel
	m.FillGeneralPackageHeaderFromPackage(buffer)
	json.Unmarshal(m.PayloadBody, &result)
	marshalParam, _ := json.Marshal(result.PARAMETER)
	switch m.OPERATION {
	case "REQUESTDOWNLOADVIDEO":
		var msr *MediaStreamModelRequestDownloadVideoResponse
		json.Unmarshal(marshalParam, msr)
		result.PARAMETER = &msr
	case "MEDIATASKSTART":
		var msr *MediaStreamModelMediaTaskStartResponseParameter
		json.Unmarshal(marshalParam, msr)
		result.PARAMETER = &msr
	case "MEDIATASKSTOP":
		var mts *MediaStreamModelMediaTaskStopResponseParameter
		json.Unmarshal(marshalParam, mts)
		result.PARAMETER = &mts
	}
	return result
}

func (m MediaStreamModel) handleRequestDownloadVideoResponse(c interfaces.IChannel, buffer []byte) {

}
func (m MediaStreamModel) handleMediaTaskStartResponse(c interfaces.IChannel, buffer []byte) {

}

func (m MediaStreamModel) handleMediaTaskStopResponse(c interfaces.IChannel, buffer []byte) {

}

func RequestFile(c interfaces.IChannel, streamName string, streamType int, recordId string, channel int, startTime string, endTime string) {
	m := &MediaStreamModel{
		GeneralPackageHeader: GeneralPackageHeader{},
		MODULE:               "MEDIASTREAMMODEL",
		OPERATION:            "REQUESTDOWNLOADVIDEO",
		PARAMETER: MediaStreamModelRequestDownloadVideoParameter{
			PT:         3,
			SSRC:       128,
			STREAMNAME: streamName,
			STREAMTYPE: streamType,
			RECORDID:   recordId,
			CHANNEL:    channel,
			STARTTIME:  startTime,
			ENDTIME:    endTime,
			OFFSETFLAG: 0,
			OFFSET:     0,
			IPANDPORT:  config.GetConfig().VideoServerWANIP + ":" + strconv.Itoa(config.GetConfig().VideoServerPort),
			SERIAL:     rand.Intn(0xffff),
		},
		SESSION:  c.GetCurrentSession(),
		RESPONSE: nil,
	}
	header := GeneralPackageHeader{}
	marshal, _ := json.Marshal(m)
	headerBytes := header.toHeaderBytes(uint(len(marshal)))
	bytes := append(headerBytes, marshal...)
	PrintDebugPackageInfo(bytes)
	c.SendBytes(bytes)
}

func ControlDownloadVideo(c interfaces.IChannel, streamName string, cmd int) {
	m := &MediaStreamModel{
		GeneralPackageHeader: GeneralPackageHeader{},
		MODULE:               "MEDIASTREAMMODEL",
		OPERATION:            "CONTROLDOWNLOADVIDEO",
		PARAMETER: MediaStreamModelControlDownloadVideoParameter{
			CSRC:       "",
			PT:         0,
			SSRC:       0,
			STREAMNAME: streamName,
			CMD:        cmd,
			DT:         0,
		},
		SESSION:  c.GetCurrentSession(),
		RESPONSE: nil,
	}
	header := GeneralPackageHeader{}
	marshal, _ := json.Marshal(m)
	headerBytes := header.toHeaderBytes(uint(len(marshal)))
	bytes := append(headerBytes, marshal...)
	PrintDebugPackageInfo(bytes)
	c.SendBytes(bytes)
}

//"OPERATION": "REQUESTDOWNLOADVIDEO"  ACK
type MediaStreamModelRequestDownloadVideoResponse struct {
	ERRORCAUSE   string
	ERRORCODE    int
	FILESIZE     int
	LEFTFILESIZE int
	SERIAL       int
	STREAMNAME   string
}

type MediaStreamModelRequestDownloadVideoParameter struct {
	PT         int //payload type as in package header
	SSRC       int
	STREAMNAME string
	STREAMTYPE int
	RECORDID   string
	CHANNEL    int
	STARTTIME  string
	ENDTIME    string
	OFFSETFLAG int
	OFFSET     int
	IPANDPORT  string
	SERIAL     int
}

type MediaStreamModelControlDownloadVideoParameter struct {
	CSRC       string
	PT         int //payload type as in package header
	SSRC       int
	STREAMNAME string
	//Control the media task operation (0: stop, 1: restore download, 2: pause, 3: switch download mode.
	//If issue this command, will immediately execute the new command, continue to download until the completion of the download task)
	CMD int
	//Download mode, the default is normal download.
	//If this field does not exist, it is also normal download
	//0: normal download
	DT int
}

//"OPERATION":"MEDIATASKSTART"  ACK (callback)
type MediaStreamModelMediaTaskStartResponseParameter struct {
	CSRC       string
	IPANDPORT  string
	PT         int //payload type as in package header
	SSRC       int //128
	STREAMNAME string
}

//"OPERATION":"MEDIATASKSTOP"  ACK (callback)
type MediaStreamModelMediaTaskStopResponseParameter struct {
	MediaStreamModelMediaTaskStartResponseParameter
	ERRORCAUSE string
	ERRORCODE  int
}
