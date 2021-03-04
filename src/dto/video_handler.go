package dto

import (
	"dto/videoContainer"
	"fmt"
	"interfaces"
	"os"
)

type VideoHandler struct {
	videoFile       *os.File
	offset          int64
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
	segmentBuffer := HandlePackageWithSO(handler.containerBuffer, buffer[12:], c, &videoContainer.VideoFrameHeader{}, handleNormalizedPayload)
	for segmentBuffer != nil {
		segmentBuffer = HandlePackageWithSO(handler.containerBuffer, segmentBuffer, c, &videoContainer.VideoFrameHeader{}, handleNormalizedPayload)
	}
}

func handleNormalizedPayload(c interfaces.IChannel, buffer []byte) {
	videoHandler := c.GetVideoHandler().(*VideoHandler)
	frames := videoContainer.ParseIntegerFrames(buffer[:])
	for _, fr := range frames {
		if fr.Header.FrameType == videoContainer.H264I || fr.Header.FrameType == videoContainer.H264P {
			videoHandler.videoFile.Write(fr.Data)
		}
	}
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
		videoFile:       createFile("server_file_04.03"),
		offset:          0,
		containerBuffer: &TSOBuffer{},
	}
}
