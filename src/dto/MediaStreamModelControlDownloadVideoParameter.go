package dto

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
