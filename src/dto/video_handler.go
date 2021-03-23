package dto

import (
	"crypto/rand"
	"dto/videoContainer"
	"fmt"
	"interfaces"
	"log"
	"os"
	"strconv"
	"time"
)

type VideoHandler struct {
	videoFile        *os.File
	audioFile        *os.File
	rawFile          *os.File
	frameCounter     int
	frameMap         map[int]int
	containerBuffer  *TSOBuffer
	audioCodecId     int
	videoCodecId     int
	timeStart        time.Time
	filePrefix       string
	videoFileCounter int
}

type VideoHandlerModule struct {
}

func (v VideoHandlerModule) HandleRequest(c interfaces.IChannel, buffer []byte) {
	//fmt.Printf("new video package with %v bytes", len(buffer))
	handler := c.GetVideoHandler().(*VideoHandler)
	if handler == nil {
		vh := interface{}(createVideoHandler())
		c.SetVideoHandler(vh)
	}
	handler = c.GetVideoHandler().(*VideoHandler)
	handler.rawFile.Write(buffer)
	segmentBuffer := HandlePackageWithSO(handler.containerBuffer, buffer[12:], c, &videoContainer.VideoFrameHeader{}, handleNormalizedPayload)
	for segmentBuffer != nil {
		segmentBuffer = HandlePackageWithSO(handler.containerBuffer, segmentBuffer, c, &videoContainer.VideoFrameHeader{}, handleNormalizedPayload)
	}
}

func handleNormalizedPayload(c interfaces.IChannel, buffer []byte) {
	videoHandler := c.GetVideoHandler().(*VideoHandler)
	frames := videoContainer.ParseIntegerFrames(buffer[:])
	for _, fr := range frames {
		switch fr.Header.FrameType {
		case videoContainer.H264I:
			process264IFrame(videoHandler, fr)
		case videoContainer.H264P:
			process264PFrame(videoHandler, fr)
		case videoContainer.H265I:
			process265IFrame(videoHandler, fr)
		case videoContainer.H265P:
			process265PFrame(videoHandler, fr)
		case videoContainer.MJPEG:
			processMJPEGFrame(videoHandler, fr)
		case videoContainer.Audio:
			processAudioFrame(fr, videoHandler)
		}
	}
}

func generateFilename(videoHandler *VideoHandler) string {
	return videoHandler.timeStart.String() + "_" + videoHandler.filePrefix + "_part"
}

func processMJPEGFrame(videoHandler *VideoHandler, fr *videoContainer.VideoFrame) {

}

func process265IFrame(videoHandler *VideoHandler, fr *videoContainer.VideoFrame) {
	if setVideoCodec(videoHandler, videoContainer.H265I) {
		updateVideoFileAfterCodecChanged(videoHandler)
	}
	updateICounterMap(videoHandler)
	videoHandler.videoFile.Write(fr.Data)
}

func updateVideoFileAfterCodecChanged(videoHandler *VideoHandler) {
	videoHandler.videoFileCounter++
	videoHandler.videoFile.Close()
	videoHandler.videoFile = createFile(generateFilename(videoHandler) + strconv.Itoa(videoHandler.videoFileCounter) + ".video")
}

func process265PFrame(videoHandler *VideoHandler, fr *videoContainer.VideoFrame) {
	videoHandler.frameCounter++
	if setVideoCodec(videoHandler, videoContainer.H265P) {
		updateVideoFileAfterCodecChanged(videoHandler)
	}
	videoHandler.videoFile.Write(fr.Data)
}

func (v *VideoHandler) CloseFiles() {
	if v.audioFile != nil {
		v.audioFile.Close()
	}
	if v.videoFile != nil {
		v.videoFile.Close()
	}
	if v.rawFile != nil {
		v.rawFile.Close()
	}
}

//return true if value were updated
func setVideoCodec(videoHandler *VideoHandler, codec int) bool {
	vc := switchVideoCodec(codec)
	if videoHandler.videoCodecId == -1 {
		videoHandler.videoCodecId = vc
		return true
	}
	if videoHandler.videoCodecId == vc {
		return false
	} else {
		videoHandler.videoCodecId = vc
		return true
	}
}

func switchVideoCodec(cc int) int {
	if cc == videoContainer.H264I || cc == videoContainer.H264P {
		return 0
	}
	if cc == videoContainer.H265I || cc == videoContainer.H265P {
		return 1
	}
	if cc == videoContainer.MJPEG {
		return 2
	}
	return -1
}

func process264PFrame(videoHandler *VideoHandler, fr *videoContainer.VideoFrame) {
	videoHandler.frameCounter++
	videoHandler.videoFile.Write(fr.Data)
}

func process264IFrame(videoHandler *VideoHandler, fr *videoContainer.VideoFrame) {
	updateICounterMap(videoHandler)
	videoHandler.videoFile.Write(fr.Data)
}

func updateICounterMap(videoHandler *VideoHandler) {
	fc := videoHandler.frameCounter
	if _, found := videoHandler.frameMap[fc]; found {
		videoHandler.frameMap[fc] = fc + 1
	} else {
		videoHandler.frameMap[fc] = 1
	}
	videoHandler.frameCounter = 1
}

func processAudioFrame(fr *videoContainer.VideoFrame, videoHandler *VideoHandler) {
	headers := fr.InfoHeaders
	for _, h := range headers {
		if h.InfoType == 5 {
			pt := (h.InfoPayload[0] & 0xF0) >> 4
			if int(pt) != videoHandler.audioCodecId {
				if videoHandler.audioCodecId == -1 {
					videoHandler.audioCodecId = int(pt)
					continue
				}
				renewAudioCodec(videoHandler, int(pt))
			}
		}
	}
	videoHandler.audioFile.Write(fr.Data[4:])
}

func renewAudioCodec(vh *VideoHandler, codecId int) {
	if vh.audioFile != nil {
		vh.audioFile.Close()
	}
	vh.audioFile = createFile("server_file_" + time.Now().String() + "_acodec_" + strconv.Itoa(codecId))
	vh.audioCodecId = codecId
}

func (v VideoHandler) GetFrameMedian() int {
	cK, maxV := 0, 0
	for k, v := range v.frameMap {
		if maxV < v {
			cK, maxV = k, v
		}
	}
	return cK
}

func (v VideoHandlerModule) ParseDtoFromData(buffer []byte) interface{} {
	return v
}

func createFile(name string) *os.File {
	fmt.Println("creating file")
	f, err := os.Create(name) // creates a file at current directory
	if err != nil {
		fmt.Println(err)
	}
	return f
}

func CreateFileRandomPrefix() string {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x", b)
	return uuid
}

func createVideoHandler() *VideoHandler {
	tNow := time.Now()
	t := tNow.Format("2006-01-02T15:04:05Z07:00")
	prefix := CreateFileRandomPrefix()
	return &VideoHandler{
		videoFile:        createFile(t + "_" + prefix + "_part0" + ".video"),
		audioFile:        createFile(t + "_" + prefix + "_part0" + ".audio"),
		rawFile:          createFile(t + "_" + prefix + ".h264"),
		frameCounter:     0,
		containerBuffer:  &TSOBuffer{},
		frameMap:         make(map[int]int),
		audioCodecId:     -1,
		videoCodecId:     -1,
		filePrefix:       prefix,
		timeStart:        tNow,
		videoFileCounter: 0,
	}
}

func (v *VideoHandler) DownloadFinished() {

}
