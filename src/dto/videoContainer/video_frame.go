package videoContainer

type VideoFrame struct {
	RawHeader []byte
	Header    *VideoFrameHeader // 12 byte
	Extension []byte            // <ExtendedLen> is total len of all ex. blocks
	Data      []byte            // audio and video
}

type VideoFrameHeader struct { // 12 byte
	ChannelNum    int // 8 bit //right order
	FrameType     int // 24 bit // right order
	FrameLen      int // 24 bit // reverse order by byte cell // only for video frame data length
	StreamSum     int // 8 bit  // no flip for just 1 byte
	ExtendedLen   int // 24 bit // reverse order by byte cell // 0 for no extension block
	ExtendedCount int // 8 bit  // no flip for just 1 byte
}

func ParseVideoFrameHeader(buffer []byte) *VideoFrameHeader {
	if len(buffer) < 12 {
		return nil
	}
	h := &VideoFrameHeader{}
	h.ChannelNum = int(buffer[0])
	h.FrameType = int(buffer[1])<<16 + int(buffer[2])<<8 + int(buffer[3])
	flr := ReverseOrder(buffer[4:7])
	h.FrameLen = int(flr[0])<<16 + int(flr[1])<<8 + int(flr[2])
	h.StreamSum = int(buffer[7])
	elr := ReverseOrder(buffer[8:11])
	h.ExtendedLen = int(elr[0])<<16 + int(elr[1])<<8 + int(elr[2])
	h.ExtendedCount = int(buffer[11])
	return h
}

func ParseVideoFrame(b []byte) *VideoFrame {
	if len(b) < 12 {
		return nil
	}
	v := &VideoFrame{}
	return v
}

type InfoTypeHeader struct { // 4 byte
	infoType   int // 8 bit
	infoLength int // 24 bit  // reverse order, counts this header itself
}

const (
	H264I = 0x326463 //FrameTypes
	H264P = 0x336463
	Audio = 0x346463
	H265I = 0x356463
	H265P = 0x366463
	MJPEG = 0x376463
)

func ReverseOrder(arr []byte) []byte {
	rev := make([]byte, 0)
	copy(rev, arr)
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	return rev
}
