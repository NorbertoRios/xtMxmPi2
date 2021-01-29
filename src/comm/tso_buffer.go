package comm

import (
	"dto"
	"time"
)

type TSOBuffer struct {
	header                 *dto.GeneralPackageHeader
	buffer                 []byte
	ttlSeconds             int
	timeCreated            time.Time
	expectedPayloadLen     uint
	currentPayloadLen      int
	bytesNeeded            int
	segmentationInProgress bool
}

func (b TSOBuffer) resetBuffer() {
	b.buffer = nil
	b.ttlSeconds = 30
	b.timeCreated = time.Unix(0, 0)
	b.expectedPayloadLen = 0
	b.currentPayloadLen = 0
	b.bytesNeeded = 0
	b.segmentationInProgress = false
	b.header = nil
}

func (b TSOBuffer) initBuffer(buffer []byte) {
	b.resetBuffer()
	h := &dto.GeneralPackageHeader{}
	b.header = h.FillGeneralPackageHeaderFromPackage(buffer)
	b.buffer = make([]byte, 4096)
	b.timeCreated = time.Now()
	b.expectedPayloadLen = b.header.PayloadLen
	b.currentPayloadLen = len(b.buffer) - 12
	b.bytesNeeded = int(b.header.PayloadLen) - b.currentPayloadLen
	b.segmentationInProgress = true
}

func (b TSOBuffer) addSegment(segment []byte) bool {
	if b.segmentationInProgress {
		if len(segment) <= b.bytesNeeded &&
			b.ttlSeconds > int(time.Now().Sub(b.timeCreated).Seconds()) {
			b.buffer = append(b.buffer, segment...)
			b.currentPayloadLen += len(segment)
			b.bytesNeeded = int(b.expectedPayloadLen) - b.currentPayloadLen
			if b.bytesNeeded < 0 {
				b.resetBuffer()
				return false
			}
			return true
		} else {
			b.resetBuffer()
			return false
		}
	}
	return false
}

func (b TSOBuffer) isBufferReady() bool {
	if b.segmentationInProgress && b.bytesNeeded == 0 && b.currentPayloadLen == int(b.expectedPayloadLen) {
		return true
	} else {
		return false
	}
}

func IsSegmented(buffer []byte) bool {
	pLen := uint(buffer[4])<<24 + uint(buffer[5])<<16 + uint(buffer[6])<<8 + uint(buffer[7])
	if len(buffer) < int(pLen)+12 {
		return true
	} else {
		return false
	}
}
