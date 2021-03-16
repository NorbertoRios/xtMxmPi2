package dto

import (
	"dto/videoContainer"
	"fmt"
	"interfaces"
	"os"
)

type VideoHandler struct {
	videoFile       *os.File
	audioFile       *os.File
	rawFile         *os.File
	frameCounter    int
	frameMap        map[int]int
	containerBuffer *TSOBuffer
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
		if fr.Header.FrameType == videoContainer.Audio {
			videoHandler.audioFile.Write(fr.Data[4:])
		}
		if fr.Header.FrameType == videoContainer.H264I {
			fc := videoHandler.frameCounter
			if _, found := videoHandler.frameMap[fc]; found {
				videoHandler.frameMap[fc] = fc + 1
			} else {
				videoHandler.frameMap[fc] = 1
			}
			videoHandler.frameCounter = 1
			videoHandler.videoFile.Write(fr.Data)
		}
		if fr.Header.FrameType == videoContainer.H264P {
			videoHandler.frameCounter++
			videoHandler.videoFile.Write(fr.Data)
		}
	}
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

func createVideoHandler() *VideoHandler {
	return &VideoHandler{
		videoFile:       createFile("server_file_16.03_g711a_lq.video"),
		audioFile:       createFile("server_file_16.03_g711a_lq.audio"),
		rawFile:         createFile("server_file_16.03_g711a_lq.h264"),
		frameCounter:    0,
		containerBuffer: &TSOBuffer{},
		frameMap:        make(map[int]int),
	}
}
