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

func ParseIntegerFrames(b []byte) []*VideoFrame {
	if len(b) < 12 {
		return nil
	}
	vf := make([]*VideoFrame, 0)
	for i, l := 0, len(b); i < l; {
		var frame *VideoFrame
		if i+12 < l {
			header := ParseVideoFrameHeader(b[i : i+12])
			frL := header.ExtendedLen + header.FrameLen + 12
			if frL <= len(b) {
				frame = ParseVideoFrame(b[i : i+frL])
				if frame != nil {
					vf = append(vf, frame)
				}
				i += frL
			} else {
				break
			}
		} else {
			break
		}
	}
	return vf
}

func ParseVideoFrame(b []byte) *VideoFrame {
	if len(b) < 12 {
		return nil
	}
	v := &VideoFrame{}
	v.RawHeader = b[:12]
	v.Header = ParseVideoFrameHeader(v.RawHeader)
	start := 11 + v.Header.ExtendedLen
	end := 12 + v.Header.ExtendedLen + v.Header.FrameLen
	if start < len(b) && end <= len(b) {
		v.Data = b[start:end]
		v.Extension = b[11 : 12+v.Header.ExtendedLen]
	} else {
		return nil
	}
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
	rev = append(rev, arr...)
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	return rev
}
