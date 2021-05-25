package dto

//"OPERATION": "REQUESTDOWNLOADVIDEO"  ACK
type MediaStreamModelRequestDownloadVideoResponse struct {
	ERRORCAUSE   string
	ERRORCODE    int
	FILESIZE     int
	LEFTFILESIZE int
	SERIAL       int
	STREAMNAME   string
}
