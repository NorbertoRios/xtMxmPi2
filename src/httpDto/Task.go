package httpDto

type Task struct {
	StartTime int64    `json:"startTime,omitempty"`
	EndTime   int64    `json:"endTime,omitempty"`
	Dsno      string   `json:"dsno,omitempty"`
	Channels  Channels `json:"channels,omitempty"`
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
