package dto

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
