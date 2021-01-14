package dto

type ChannelNumberAlarmParameter struct {
	Channel     uint `json:"CHANNEL,omitempty"`     // position of bit is considered a channel number starts from 1
	ChannelMask uint `json:"CHANNELMASK,omitempty"` // position of bit is considered a channel alarm valid starts from 1
	//LCH[N] int
	PUSH      uint   // position of bit channel needs to push
	ALARMNAME string //Indicates the name of the alarm, 32 bytes
	SER       string //Abbreviation of the name of the alarm
}

type MemoryAbnormalAlarmParameter struct {
	STORAGETYPE      byte //0 - hard drive 1-U disk 2-SD card
	STORAGEINDEX     int  //1-32 (logical number of memory)
	ERRORCODE        int
	ERRORDESCRIPTION string
	PUSH             byte // 1 need to push, 0 don't
}

type use int

var useAlarmEnum = struct {
	No              use
	EmergencyCall   use
	Call            use
	Neutral         use
	AirConditioning use
	PressureTable   use
	High            use
	GreenCover      use
	FrontDoor       use
	FrontDoorOpen   use
	FrontDoorOff    use
	Gate1           use
	//Back Door 1 use
	//Back Door 1 use
	//Back Door 1 open use
	//Back Door 1 off use
	NearLight          use
	FarLight           use
	RightTurn          use
	LeftTurn           use
	SecurityCommitment use
	ForgottenAlarm     use
}{
	No:                 0,
	EmergencyCall:      1,
	Call:               2,
	Neutral:            3,
	AirConditioning:    4,
	PressureTable:      5,
	High:               6,
	GreenCover:         7,
	FrontDoor:          8,
	FrontDoorOpen:      9,
	FrontDoorOff:       10,
	Gate1:              11,
	NearLight:          17,
	FarLight:           18,
	RightTurn:          19,
	LeftTurn:           20,
	SecurityCommitment: 22,
	ForgottenAlarm:     23,
}

type UserDefinedAlarmParameter struct {
	SNO       byte
	ALARMNAME string
	SER       string //Abbreviation of the name of the alarm
	//LCH[N] int // SNO indicates a valid alarm IO number, starting from 1,
	//and each LCH element represents the bit set corresponding to the channel number
	//associated with the Nth IO alarm, which is a bit set. For example: IO1 motion detection alarm,
	//linkage 0 channel, 1 channel, 2 channel video, said LCH [1] = 15 (decimal).
	PUSH uint // 0, 1

	//USE indicates IO usage
	//Usage:
	//0: no;
	//1: emergency call;
	//2: call;
	//3: neutral;
	//4: air conditioning;
	//5: pressure table;
	//6: high;
	//7: green cover;
	//8: front door;
	//9: front door open;
	//10: front door off;
	//11: gate 1;
	//13: back door 1;
	//15: back door 1;
	//15: back door 1 open;
	//16: back door 1 off;
	//17: near light;
	//18: far light;
	//19: right turn;
	//20: left turn;
	//22: security commitment;
	//23: forgotten alarm
	USE use
	L   string
	PC  uint
	//DN[N]
	//MN[M]
}

type SentryInspectionAlarmParameter struct {
	SNO     byte //Custom alarms are numbered by the device to distinguish between multiple custom alarms.
	CHANNEL byte //linked channel number, starting from 0
}

type EmergencyAlarmParameter struct {
	ALARMNAME string
	SER       string //Abbreviation of the name of the alarm
	LCH       uint   //The channel number of the linked video, by bit address
	PUSH      uint   // 0, 1
}
type aType int

var SpeedAlarmParameterEnum = struct {
	StandardLowSpeedAlarm      aType
	StandardHighSpeedAlarm     aType
	OverSpeedPreAlarm          aType
	RealTimeHighSpeedAlarm     aType
	EmergencyDecelerationAlarm aType
	SpeedRecoveryNormalAlarm   aType
	Parking                    aType
	ThresholdAlarm             aType
	StartAlarm                 aType
}{
	StandardLowSpeedAlarm:      0,
	StandardHighSpeedAlarm:     1,
	OverSpeedPreAlarm:          2,
	RealTimeHighSpeedAlarm:     3,
	EmergencyDecelerationAlarm: 4,
	SpeedRecoveryNormalAlarm:   5,
	Parking:                    6,
	ThresholdAlarm:             7,
	StartAlarm:                 8,
}

type at byte

var atEnum = struct {
	NoSpecificLocation at
	CircularArea       at
	RectangularArea    at
	PolygonArea        at
	RoadSection        at
}{
	NoSpecificLocation: 0,
	CircularArea:       1,
	RectangularArea:    2,
	PolygonArea:        3,
	RoadSection:        4,
}

type SpeedAlarmParameter struct {
	ATYPE     aType //Alarm category.
	CSP       int   //unit 1%/km/h
	MINSP     int   //unit 1%/km/h
	MAXSP     int   //unit 1%/km/h
	MINS      int   //unit 1%/km/h
	MAXS      int   //unit 1%/km/h
	ALARMNAME string
	SER       string //Abbreviation of the name of the alarm
	LCH       uint   //The channel number of the linked video, by bit address
	PUSH      uint   // 0, 1
	AT        at     //Position type
	I         int
	LD        int // from 0 to 5
}

type LowVoltageAlarmParameter struct {
	V    int  // 1/100 Volt
	PUSH uint // 0, 1
}
