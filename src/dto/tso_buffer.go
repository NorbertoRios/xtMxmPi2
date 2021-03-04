package dto

import (
	"interfaces"
	"time"
)

type TSOBuffer struct {
	header                 interfaces.AbstractHeader
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

func (b *TSOBuffer) initBuffer(buffer []byte, h interfaces.AbstractHeader) {
	prefPresent := b.segmentedHeaderPrefix != nil
	b.buffer = make([]byte, 0)
	if !prefPresent {
		b.resetBuffer()
	} else {
		b.buffer = append(b.buffer, b.segmentedHeaderPrefix...)
		b.segmentedHeaderPrefix = nil
	}
	b.buffer = append(b.buffer, buffer...)
	b.header = h.FillHeaderFromPackage(b.buffer)
	b.timeCreated = time.Now()
	b.expectedPayloadLen = b.header.GetPayloadLen()
	b.currentPayloadLen = len(b.buffer) - 12
	b.bytesNeeded = int(b.header.GetPayloadLen()) - b.currentPayloadLen
	b.segmentationInProgress = true
}

func (b *TSOBuffer) initBufferWithPartialHeader(buffer []byte) {
	b.resetBuffer()
	b.segmentationInProgress = true
	b.segmentedHeaderPrefix = make([]byte, 0)
	b.segmentedHeaderPrefix = append(b.segmentedHeaderPrefix, buffer...)
}

func (b *TSOBuffer) addSegment(segment []byte, h interfaces.AbstractHeader) (bool, []byte) {
	if prefPresent := b.segmentedHeaderPrefix != nil; prefPresent {
		leftForHeader := 12 - len(b.segmentedHeaderPrefix)
		b.initBuffer(segment[:leftForHeader], h)
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

func HandlePackageWithSO(tso *TSOBuffer, buffer []byte, c interfaces.IChannel, h interfaces.AbstractHeader, handler func(interfaces.IChannel, []byte)) []byte {
	if tso.segmentationInProgress {
		//fmt.Println("SEG IN PROGRESS")

		//fmt.Println(tso.bytesNeeded)
		//fmt.Printf(" needed, now size is %v", len(tso.buffer))
		if added, overflow := tso.addSegment(buffer, h); added {
			if tso.isBufferReady() {
				//fmt.Println("buffer is ready to release TSO")
				handler(c, tso.buffer)
				tso.resetBuffer()
			}
			if overflow != nil && len(overflow) > 0 {
				return overflow
			}
		} else {
			handler(c, buffer)
		}
	} else if h.IsSegmented(buffer) {
		//fmt.Println()
		//fmt.Print("SEGMENTATION FOUND IN PACKAGE :")
		//fmt.Printf(" %x", buffer[:12])
		tso.initBuffer(buffer, h)
	} else if segmented, first, overflow := h.ContainsAdditionalTCPSegment(buffer); segmented { //if merged
		handler(c, first)
		return overflow
	} else if len(buffer) < 12 {
		tso.initBufferWithPartialHeader(buffer)
	} else {
		handler(c, buffer)
	}
	return nil
}
