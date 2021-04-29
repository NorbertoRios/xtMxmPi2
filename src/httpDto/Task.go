package httpDto

import "time"

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
	if time.Unix(t.StartTime, 0).After(y0) ||
		time.Unix(t.EndTime, 0).Before(y100) ||
		time.Unix(t.StartTime, 0).Before(y100) ||
		time.Unix(t.EndTime, 0).After(y0) {
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
