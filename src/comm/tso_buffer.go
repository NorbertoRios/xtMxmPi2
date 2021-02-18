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
	segmentedHeaderPrefix  []byte
}

func (b *TSOBuffer) resetBuffer() {
	b.buffer = nil
	b.ttlSeconds = 30
	b.timeCreated = time.Unix(0, 0)
	b.expectedPayloadLen = 0
	b.currentPayloadLen = 0
	b.bytesNeeded = 0
	b.segmentationInProgress = false
	b.header = nil
	b.segmentedHeaderPrefix = nil
}

func (b *TSOBuffer) initBuffer(buffer []byte) {
	prefPresent := b.segmentedHeaderPrefix != nil
	b.buffer = make([]byte, 0)
	if !prefPresent {
		b.resetBuffer()
	} else {
		b.buffer = append(b.buffer, b.segmentedHeaderPrefix...)
		b.segmentedHeaderPrefix = nil
	}
	b.buffer = append(b.buffer, buffer...)
	h := &dto.GeneralPackageHeader{}
	b.header = h.FillGeneralPackageHeaderFromPackage(b.buffer)
	b.timeCreated = time.Now()
	b.expectedPayloadLen = b.header.PayloadLen
	b.currentPayloadLen = len(b.buffer) - 12
	b.bytesNeeded = int(b.header.PayloadLen) - b.currentPayloadLen
	b.segmentationInProgress = true
}

func (b *TSOBuffer) initBufferWithPartialHeader(buffer []byte) {
	b.resetBuffer()
	b.segmentationInProgress = true
	b.segmentedHeaderPrefix = make([]byte, 0)
	b.segmentedHeaderPrefix = append(b.segmentedHeaderPrefix, buffer...)
}

func (b *TSOBuffer) addSegment(segment []byte) (bool, []byte) {
	if prefPresent := b.segmentedHeaderPrefix != nil; prefPresent {
		leftForHeader := 12 - len(b.segmentedHeaderPrefix)
		b.initBuffer(segment[:leftForHeader])
		segment = segment[leftForHeader:]
	}
	if b.segmentationInProgress {
		if len(segment) <= b.bytesNeeded {
			b.buffer = append(b.buffer, segment...)
			b.currentPayloadLen += len(segment)
			b.bytesNeeded = int(b.expectedPayloadLen) - b.currentPayloadLen
			return true, nil
		} else {
			ovIndex := b.bytesNeeded
			b.buffer = append(b.buffer, segment[:b.bytesNeeded]...)
			b.currentPayloadLen += b.bytesNeeded
			b.bytesNeeded = int(b.expectedPayloadLen) - b.currentPayloadLen
			overflowBuffer := segment[ovIndex:]
			return true, overflowBuffer
		}
	}
	return false, nil
	//b.ttlSeconds > int(time.Now().Sub(b.timeCreated).Seconds())
}

func (b *TSOBuffer) isBufferReady() bool {
	if b.segmentationInProgress && b.bytesNeeded == 0 && b.currentPayloadLen == int(b.expectedPayloadLen) {
		return true
	} else {
		return false
	}
}

func IsSegmented(buffer []byte) bool {
	if len(buffer) < 12 {
		return false
	}
	pLen := uint(buffer[4])<<24 + uint(buffer[5])<<16 + uint(buffer[6])<<8 + uint(buffer[7])
	if len(buffer) < int(pLen)+12 {
		return true
	} else {
		return false
	}
}
