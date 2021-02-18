package dto

import (
	"comm/channel"
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
	fmt.Printf("new video package with %v bytes", len(buffer))
	handler := c.GetVideoHandler().(*VideoHandler)
	if handler == nil {
		vh := interface{}(createVideoHandler())
		c.SetVideoHandler(vh)
		videoHandler := vh.(*VideoHandler)
		videoHandler.videoFile.Write(buffer)
	} else {
		bb := buffer[12:]
		handler.videoFile.WriteAt(bb, handler.offset)
		handler.offset += int64(len(bb))
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
		videoFile: createFile("server_file"),
	}
}
