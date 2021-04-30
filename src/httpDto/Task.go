package httpDto

import (
	"time"
)

type Task struct {
	StartTime int64     `json:"startTime,omitempty"`
	EndTime   int64     `json:"endTime,omitempty"`
	Dsno      string    `json:"dsno,omitempty"`
	Channels  *Channels `json:"channels,omitempty"`
}

type Channels struct {
	Adas *Adas `json:"adas,omitempty"`
	Dsm  *Dsm  `json:"dsm,omitempty"`
	Ip   *Ip   `json:"ip,omitempty"`
}

//"stream", "substream", "screenshot"
type Adas struct {
	FileTypes []string `json:"fileTypes,omitempty"`
}

//"stream", "substream", "screenshot"
type Dsm struct {
	FileTypes []string `json:"fileTypes,omitempty"`
}

//"stream", "substream", "screenshot"
type Ip struct {
	FileTypes []string `json:"fileTypes,omitempty"`
}

func (t Task) ValidateTask() bool {
	form := "January 1, 2006"
	y0, _ := time.Parse(form, "January 1, 2000")
	y100, _ := time.Parse(form, "January 1, 2100")
	if !(time.Unix(t.StartTime, 0).After(y0) ||
		time.Unix(t.EndTime, 0).Before(y100) ||
		time.Unix(t.StartTime, 0).Before(y100) ||
		time.Unix(t.EndTime, 0).After(y0)) {
		return false
	}
	if len(t.Dsno) < 3 {
		return false
	}
	if t.Channels == nil || (t.Channels.Adas == nil && t.Channels.Dsm == nil && t.Channels.Ip == nil) {
		return false
	}
	return true
}

//1,2,3 adas,dsm,ip   bits
//0 empty
func (t Task) GetTypeFromTask(s string) int {
	var res = 0
	if t.Channels.Ip != nil {
		res = t.iterateTypes(s, t.Channels.Ip.FileTypes, res, 3)
	}
	if t.Channels.Dsm != nil {
		res = t.iterateTypes(s, t.Channels.Dsm.FileTypes, res, 2)
	}
	if t.Channels.Adas != nil {
		res = t.iterateTypes(s, t.Channels.Adas.FileTypes, res, 1)
	}
	return res
}

func (t Task) iterateTypes(s string, ar []string, res int, ch int) int {
	typesAdas := ar
	for i := 0; i < len(typesAdas); i++ {
		if typesAdas[i] == s {
			res = t.putFlagToInt(ch, res)
		}
	}
	return res
}

func (t Task) putFlagToInt(flagPos int, val int) int {
	var argument int = 1
	if flagPos < 1 {
		return val
	}
	if flagPos > 1 {
		argument = argument << (flagPos - 1)
	}
	return val | argument
}
