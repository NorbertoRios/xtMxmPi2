package dto

type GeneralPackageHeader struct {
	V            uint8  // 1,2 bits
	P            bool   // 3 bit
	M            bool   // 4 bit
	CsrcCount    uint8  //5-8bit
	PayloadType  uint8  //9-16bit
	Ssrc         uint16 //17-32bit
	PayloadLen   uint   //33-64 bit
	Reserve      uint   //64-96
	Csrc         []uint //max 15 elements of 32 bits each
	PayloadBody  []byte
	ExtendedPart []byte
}

func (p *GeneralPackageHeader) FillGeneralPackageHeaderFromPackage(buffer []byte) *GeneralPackageHeader {
	if len(buffer) < 12 {
		return nil
	}
	var headerLen uint = 0
	p.V = buffer[0] & 0xC0 >> 6
	p.P = buffer[0]&0x20>>5 != 0
	p.M = buffer[0]&0x10>>4 != 0
	p.CsrcCount = buffer[0] >> 4
	headerLen = (uint(p.GetCsrcCount())*32 + 96) / 8
	p.PayloadType = buffer[1]
	p.Ssrc = uint16(buffer[2])<<8 + uint16(buffer[3])
	p.PayloadLen = uint(buffer[4])<<24 + uint(buffer[5])<<16 + uint(buffer[6])<<8 + uint(buffer[7])
	p.Reserve = uint(buffer[8])<<24 + uint(buffer[9])<<16 + uint(buffer[10])<<8 + uint(buffer[11])
	if int(p.PayloadLen+headerLen) > len(buffer) {
		// next package contains same json
	} else {
		p.PayloadBody = buffer[headerLen : headerLen+p.PayloadLen]
		p.ExtendedPart = buffer[p.PayloadLen+headerLen:]
	}
	return p
}

func (p *GeneralPackageHeader) toHeaderBytes(payloadLen uint) []byte {
	r := [12]byte{}
	r[0] = p.V<<6 + bool2int(p.P)<<5 + p.CsrcCount
	r[1] = p.PayloadType
	r[2] = 0 //Ssrc
	r[3] = 0 //Ssrc
	r[4] = byte(payloadLen >> 24)
	r[5] = byte(payloadLen >> 16)
	r[6] = byte(payloadLen >> 8)
	r[7] = byte(payloadLen)
	r[8] = 0 //start of reserve
	r[9] = 0
	r[10] = 0
	r[11] = 0
	return r[:]
}

func (p *GeneralPackageHeader) GetCsrcCount() uint8 {
	return 0 //set to 0 because of missing Csrc when Csrc count not null
	//return p.CsrcCount
}

func bool2int(a bool) byte {
	if a {
		return 1
	} else {
		return 0
	}
}
