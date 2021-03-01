package dto

import (
	"comm/channel"
	"dto/videoContainer"
	"fmt"
	"os"
)

type VideoHandler struct {
	videoFile *os.File
	offset    int64
}

type VideoHandlerModule struct {
}

func (v VideoHandlerModule) HandleRequest(c channel.IChannel, buffer []byte) {
	//fmt.Printf("new video package with %v bytes", len(buffer))
	handler := c.GetVideoHandler().(*VideoHandler)
	if handler == nil {
		vh := interface{}(createVideoHandler())
		c.SetVideoHandler(vh)
		videoHandler := vh.(*VideoHandler)
		frames := videoContainer.ParseIntegerFrames(buffer[12:])
		for _, fr := range frames {
			if fr.Header.FrameType == videoContainer.H264I || fr.Header.FrameType == videoContainer.H264P {
				videoHandler.videoFile.Write(fr.Data)
			}
		}
	} else {
		bb := buffer[12:]
		frames := videoContainer.ParseIntegerFrames(bb)
		for _, fr := range frames {
			if fr.Header.FrameType == videoContainer.H264I || fr.Header.FrameType == videoContainer.H264P {
				handler.videoFile.WriteAt(fr.Data, handler.offset)
				handler.offset += int64(len(fr.Data))
			}
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
		videoFile: createFile("server_file_h264"),
	}
}
