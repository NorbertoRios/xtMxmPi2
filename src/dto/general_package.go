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

func (p GeneralPackageHeader) FillGeneralPackageHeaderFromPackage(buffer []byte) *GeneralPackageHeader {
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
	p.Ssrc = uint16(buffer[2]) + uint16(buffer[3])<<8
	p.PayloadLen = uint(buffer[4]) + uint(buffer[5])<<8 + uint(buffer[6])<<16 + uint(buffer[7])<<24
	p.Reserve = uint(buffer[8]) + uint(buffer[9])<<8 + uint(buffer[10])<<16 + uint(buffer[11])<<24
	p.PayloadBody = buffer[headerLen : headerLen+p.PayloadLen]
	if p.PayloadLen+headerLen > uint(len(buffer)) {
		p.ExtendedPart = buffer[p.PayloadLen+headerLen:]
	}
	return &p
}

func (p GeneralPackageHeader) GetCsrcCount() uint8 {
	return 0 //set to 0 because of missing Csrc when Csrc count not null
	//return p.CsrcCount
}
