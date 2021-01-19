package dto

import (
	"reflect"
)

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

type DayNightSwitchAlarm struct{}
type ProhibitDrivingAlarm struct{}
type IllegalIgnitionAlarm struct{}
type IllegalShutdownAlarm struct{}
type PhoneCallAlarm struct{}

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

type SerialAlarmParameter struct {
	ALARMNAME string   //Indicates the name of the alarm, 32 bytes
	SER       string   //Abbreviation of the name of the alarm
	LCH       uint     //The channel number, by bit address
	SUM       int      //Total number of serial alarms
	S         []string //Serial alarm information description
}

type direction byte

var DirectionEnum = struct {
	XDirection        direction
	YDirection        direction
	ZDirection        direction
	Collision         direction
	RollOver          direction
	Bumps             direction
	HarshAcceleration direction
	HardBrake         direction
	G4                direction
	G5                direction
	HarshLeftTurn     direction
	HarshRightTurn    direction
}{
	XDirection:        0,
	YDirection:        1,
	ZDirection:        2,
	Collision:         3,
	RollOver:          4,
	Bumps:             5,
	HarshAcceleration: 6,
	HardBrake:         7,
	G4:                8,
	G5:                9,
	HarshLeftTurn:     10,
	HarshRightTurn:    11,
}

type AccAlarmParameter struct {
	ALARMNAME string //Indicates the name of the alarm, 32 bytes
	SER       string //Abbreviation of the name of the alarm
	LCH       uint   //The channel number, by bit address
	D         direction
	X         int // *1000
	Y         int
	Z         int
	V         int //G4, G5 effective alarm
}

type fenceEvent byte

var FenceEventEnum = struct {
	IntoFence            fenceEvent
	OutOfFence           fenceEvent
	IncomingLine         fenceEvent
	OutOfLine            fenceEvent
	LineDeviation        fenceEvent
	DrivingTimeIsTooLong fenceEvent
}{
	IntoFence:            1 << 0,
	OutOfFence:           1 << 1,
	IncomingLine:         1 << 2,
	OutOfLine:            1 << 3,
	LineDeviation:        1 << 4,
	DrivingTimeIsTooLong: 1 << 5,
}

type fenceType byte

var fenceTypeEnum = struct {
	round     fenceType
	rectangle fenceType
	polygon   fenceType
	line      fenceType
}{
	round:     1,
	rectangle: 2,
	polygon:   3,
	line:      4,
}

type GeoFenceAlarmParameter struct {
	E fenceEvent

	//Out of the fence action
	//N represents the subscript of the fence event type bit,
	// AC[N] uint
	AT fenceType
	ID int    //fence ID
	AN string //fence name
}

type stopAnnouncement byte

var stopAnnouncementEnum = struct {
	RetentionInStop         stopAnnouncement
	PassStopWithoutStopping stopAnnouncement
}{
	RetentionInStop:         0,
	PassStopWithoutStopping: 1,
}

type PeripheralDroppedAlarmParameter struct {
	AN  string // name
	SER string //Abbreviation of the name of the alarm
	ID  int    //Peripheral ID
}

type StopAnnouncementAlarmParameter struct {
	AN  string           // name
	SER string           //Abbreviation of the name of the alarm
	S   stopAnnouncement //Stop announcement alarm subtype
}

type FatigueDrivingSubtype byte

var FatigueDrivingSubtypeEnum = struct {
	FatigueAlarm                FatigueDrivingSubtype
	FatiguePreAlarm             FatigueDrivingSubtype
	DailyCumulativeFatigueAlarm FatigueDrivingSubtype
}{
	FatigueAlarm:                0,
	FatiguePreAlarm:             1,
	DailyCumulativeFatigueAlarm: 2,
}

type FatigueDrivingAlarmParameter struct {
	AN string //name
	AS int    //Fatigue driving subtype
	C  int    //Current driving duration length in seconds
	D  uint   //Fatigue level, 0-100, use non-negative integer to represent
}

type TimeoutParkingAlarmParameter struct {
	AN string //name
	C  int    //The length of the current parking duration, in seconds
}

type GestureAlarmType byte

var GestureAlarmTypeEnum = struct {
	Brakes                             GestureAlarmType
	AnxiousToSpeedUp                   GestureAlarmType
	SharpTurn                          GestureAlarmType
	NeutralGlide                       GestureAlarmType
	EngineOverturned                   GestureAlarmType
	IdleTimeIsTooLong                  GestureAlarmType
	LaneDeviationLeftLine              GestureAlarmType
	ForwardCollision                   GestureAlarmType
	CarDistanceTooClose                GestureAlarmType
	Rollover                           GestureAlarmType
	Collision                          GestureAlarmType
	LaneDepartureFromPreAlarmRightLine GestureAlarmType
	VideoOcclusion                     GestureAlarmType
	FatigueAlarmFromPeripherals        GestureAlarmType
	AntiCollisionFromPeripherals       GestureAlarmType
	PhoneCallMakingFromPeripherals     GestureAlarmType
	SmokingFromPeripherals             GestureAlarmType
	WatchingTheFront                   GestureAlarmType
	SilentHelpRequest                  GestureAlarmType
}{
	Brakes:                             0,
	AnxiousToSpeedUp:                   1,
	SharpTurn:                          2,
	NeutralGlide:                       3,
	EngineOverturned:                   4,
	IdleTimeIsTooLong:                  5,
	LaneDeviationLeftLine:              6,
	ForwardCollision:                   7,
	CarDistanceTooClose:                8,
	Rollover:                           9,
	Collision:                          10,
	LaneDepartureFromPreAlarmRightLine: 11,
	VideoOcclusion:                     12,
	FatigueAlarmFromPeripherals:        13,
	AntiCollisionFromPeripherals:       14,
	PhoneCallMakingFromPeripherals:     15,
	SmokingFromPeripherals:             16,
	WatchingTheFront:                   17,
	SilentHelpRequest:                  18,
}

type GestureDataSRC string

var GestureDataSRCEnum = struct {
	CanInfo   GestureDataSRC
	SixAxis   GestureDataSRC
	TirePress GestureDataSRC
	TruckLoad GestureDataSRC
}{
	CanInfo:   "CANINFO（OBD）",
	SixAxis:   "6-AXIS",
	TirePress: "TIREPRESS",
	TruckLoad: "TRUCKLOAD",
}

var GestureDataSRCMap = map[string]reflect.Type{
	string(GestureDataSRCEnum.CanInfo):   reflect.TypeOf(OBDOriginalStatusParameter{}),
	string(GestureDataSRCEnum.SixAxis):   reflect.TypeOf(AXISOriginalStatusParameter{}),
	string(GestureDataSRCEnum.TirePress): reflect.TypeOf(TyrePressureOriginalStatusParameter{}),
	string(GestureDataSRCEnum.TruckLoad): reflect.TypeOf(LoadOriginalStatusParameter{}),
}

type GestureAlarmParameter struct {
	ALARMNAME string
	SER       string // abbreviation
	LCH       uint   //The channel number of the linked video, bit representation
	S         GestureAlarmType
	DATSRC    GestureDataSRC
	DATA      interface{}
}

type OBDOriginalStatusParameter struct {
	REFVALUE    int     //Threshold
	RPM         int     //Engine RPM
	COOLANTTEMP int     //Coolant temperature
	OILPRESSURE float32 //Engine oil pressure
	KEYSTATUS   bool    //Car key switch
	ACCEL       int     //The accelerator pedal depresses the percentage
	SPEED       int     //Speed
	BATTERY     float32 //Battery voltage
	OILWEAR     float32 //Total fuel consumption
	OILTEMP     float32 //Engine oil temperature
	ENGINENUM   string  //Engine number
}

type AXISOriginalStatusParameter struct {
	REFVALUE         int     //Threshold values to trigger an event.
	DATSRC           string  //Raw data source
	ACCELX           float32 //Acceleration in X axis. Unit:G.；
	ACCELY           float32 //Acceleration in Y axis. Unit:G.
	ACCELZ           float32 //Acceleration in Z axis. Unit:G.
	ANGULARVELOCITYX float32 //Rotational speed in X axis. Unit:degree/second.
	ANGULARVELOCITYY float32 //Rotational speed in Y axis. Unit:degree/second.
	ANGULARVELOCITYZ float32 //Rotational speed in Z axis. Unit:degree/second.
}

type TyrePressureOriginalStatusParameter struct {
	REFVALUE string //Threshold values to trigger an event.
	DATSRC   string //Raw data source
	TNUM     uint   //The quantity of tire
	TIRE     []uint //Tire data array
}

type TyreParameter struct {
	ID    string  //Serial ID
	NAME  string  //Name
	PRESS float32 //Pressure (float)
	TEMP  float32 //Temperature (float)
}

type LoadOriginalStatusParameter struct {
	REFVALUE  string //Threshold values to trigger an event.
	DATSRC    string //Raw data source
	AXLESNUM  int    //Number of axles
	LOADLIMIT int    //Load limit
	LEFT      int    //Left load
	RIGHT     int    //Right load
	TOTAL     int    //Total load
}

type GreenDrivingAlarmParameter struct {
	DGSL GreenDrivingDataParameter
}

type GreenDrivingDataParameter struct {
	B []interface{}
}

type CustomExternalInputAlarm struct {
	SNO       byte //alarm id
	ALARMNAME string
	SER       string //Abbreviation of the name of the alarm
	//SNO indicates valid alarm IO number, starting from 1, and each LCH element
	//represents the bit set corresponding to the linkage channel number associated
	//with the Nth IO alarm, this set is a bit set. For example: IO1 occurs motion
	//detection alarm, linked to video from channel 0, channel 1, channel 2,
	//means LCH [1] = 15 (decimal).
	//LCH[N] uint
	PUSH byte // 1 need to push, 0 don't
	FROM uint
	SUB  CustomExternalInputAlarmSubtype
}

type CustomExternalInputAlarmSubtype byte

var CustomExternalInputAlarmSubtypeEnum = struct {
	NoAlarm           CustomExternalInputAlarmSubtype
	AlarmOccurred     CustomExternalInputAlarmSubtype
	TrimmingOccurs    CustomExternalInputAlarmSubtype
	ShortCircuitAlarm CustomExternalInputAlarmSubtype
}{
	NoAlarm:           0,
	AlarmOccurred:     1,
	TrimmingOccurs:    2,
	ShortCircuitAlarm: 3,
}

type BusLaneOccupationAlarmParameter struct {
	F         byte //Bus lane occupation detection type 0: to IPC intelligent module active detection.
	ALARMNAME string
	SER       string //Abbreviation of the name of the alarm
	PUSH      byte   // 1 need to push, 0 don't
	NO        []interface{}
	C         []interface{}
}

type OilVolumeAlarmParameter struct {
	ALARMNAME string //Indicates the name of the alarm, 32 bytes
	SER       string //Abbreviation of the name of the alarm
	LCH       byte   //he channel number of the linked video, bit represent
}

type SpecialCustomerMalfunctionAlarmParameter struct {
	AS SpecialCustomerMalfunctionType
}

type SpecialCustomerMalfunctionType byte

var SpecialCustomerMalfunctionTypeEnum = struct {
	LvLowVoltageMalfunction      SpecialCustomerMalfunctionType
	HvHighVoltageMalfunction     SpecialCustomerMalfunctionType
	HtHighTemperatureMalfunction SpecialCustomerMalfunctionType
	HeaterFailure                SpecialCustomerMalfunctionType
	WatchdogReset                SpecialCustomerMalfunctionType
	OverCurrent                  SpecialCustomerMalfunctionType
	TimeInvalid                  SpecialCustomerMalfunctionType
	SDFull                       SpecialCustomerMalfunctionType
	USBFull                      SpecialCustomerMalfunctionType
	FirmwareUpdate               SpecialCustomerMalfunctionType
	ConfigLoad                   SpecialCustomerMalfunctionType
	TimeJumpMoreThan24Hours      SpecialCustomerMalfunctionType
}{
	LvLowVoltageMalfunction:      0,
	HvHighVoltageMalfunction:     1,
	HtHighTemperatureMalfunction: 2,
	HeaterFailure:                3,
	WatchdogReset:                4,
	OverCurrent:                  5,
	TimeInvalid:                  6,
	SDFull:                       7,
	USBFull:                      8,
	FirmwareUpdate:               9,
	ConfigLoad:                   10,
	TimeJumpMoreThan24Hours:      11,
}

type TemperatureAbnormallyAlarmParameter struct { //3.4.1.5.27
	ALARMNAME string
	S         byte //0 high temperature, 1 low temperature
	CU        int  //Current temperature
	UP        int  //High temperature range
	LOW       int  //Low temperature range
	ID        int  //Sensor ID
	TH        int  //Alarm threshold
	TN        int  //Number of sensors
}

type AbnormalTemperatureChangeAlarmParameter struct {
	ALARMNAME string
	SER       string //Abbreviation of the name of the alarm
	LCH       byte   //channel number of the linked video, bit represent
	TU        int    //Temperature unit 0: (1/100) degrees Celsius 1: (1/100) degrees Fahrenheit
	T         int    //Temperature value
	AT        fenceType
	ID        int    //Fence ID
	AN        string //fence name
	SID       int    //The ID of the sensor, starts from 1
}

type SmokeAlarmParameter struct {
	ALARMNAME string
	SER       string //Abbreviation of the name of the alarm
	LCH       byte   //channel number of the linked video, bit represent
	SDU       int    //Smoke concentration unit 0：(1/10000)mg/m3 ; 1：(1/100)ppb
	SD        int    //Smoke concentration value
	AT        fenceType
	ID        int    //Fence ID
	AN        string //fence name
	SID       int    //The ID of the sensor, starts from 1
}

type GBoxAlarmParameter struct {
	AS GBoxASType
}

type GBoxASType byte

var GBoxASTypeEnum = struct {
	RapidDeceleration                        GBoxASType
	RapidAcceleration                        GBoxASType
	OverSpeed                                GBoxASType
	FatigueEvent                             GBoxASType
	EngineOverturned                         GBoxASType
	UseAirConditionerWhenIdling              GBoxASType
	NeutralSliding                           GBoxASType
	TooLongIdle                              GBoxASType
	ColdStart                                GBoxASType
	OpenTheDoorWhenVehicleIsMovingFrontDoor  GBoxASType
	OpenTheDoorWhenVehicleIsMovingMiddleDoor GBoxASType
}{
	RapidDeceleration:                        1,
	RapidAcceleration:                        2,
	OverSpeed:                                3,
	FatigueEvent:                             4,
	EngineOverturned:                         5,
	UseAirConditionerWhenIdling:              6,
	NeutralSliding:                           7,
	TooLongIdle:                              8,
	ColdStart:                                9,
	OpenTheDoorWhenVehicleIsMovingFrontDoor:  10,
	OpenTheDoorWhenVehicleIsMovingMiddleDoor: 11,
}

type LicensePlateRecognitionAlarmParameter struct {
	//CAR[N] string //Identified license plate number,
}

type speedUnit byte

var speedUnitEnum = struct {
	KMpH speedUnit
	MpH  speedUnit
}{
	KMpH: 0,
	MpH:  1,
}

type SwitchDoorWhenVehicleIsMovingAlarmParameter struct { //3.4.1.5.32
	ALARMNAME string
	SER       string //Abbreviation of the name of the alarm
	DR0       byte   //front door
	DR1       byte   //door
	DR2       byte   //back door
	D         speedUnit
	V         int // speed/100
	X         int //speed limit, speed/100
}

type WirelessSignalAbnormalityAlarmParameter struct {
	ID byte //0: Communication module 1 ; 1: Communication module 2
	MS byte //Communication module status 0: module does not exist ; 1: module exists
	SS byte //SIM card status 0: SIM card does not exist ; 1: SIM card is valid ; 2: SIM card is invalid
	S  byte // Signal strength, 0 ~ 5 grid number
}

type ArmingAlarmParameter struct {
	TH int //Arming alarm time threshold, unit: seconds
}

type GPSMalfunctionAlarmParameter struct {
	CT int //Malfunction duration, in seconds
	//GPS system time when the last valid time,
	//for example: 20160815144530 means August 15th, 2016 14:45:30
	HT string
	HP LocationInformationParameter
}

type GPSAntennaAlarmParameter struct {
	S         int //0: normal 1: open circuit 2: short circuit
	ALARMNAME string
	SER       string //Abbreviation of the name of the alarm
}

type LocationStatusType byte

var locationStatusTypeEnum = struct {
	Valid       LocationStatusType
	Invalid     LocationStatusType
	NoGpsModule LocationStatusType
}{
	Valid:       0,
	Invalid:     1,
	NoGpsModule: 2,
}

type LocationInformationParameter struct {
	V LocationStatusType
	//longitude.
	//Character string description of float number, 6 bit after decimal point
	//in the format of dddd.mmmmmm(all the zeros need keep)
	//dddd Range:-179~179
	//Mmmmmmmm Range:0~999999
	//Positive value means east longitude, Negative value means west longitude.
	//E.g ’-080.092000’ stand for west longitude 80.092.
	J string
	//Latitude.
	//Character string described float number, 6 bit after decimal point
	//in the format of ddd.mmmmmm (all the zeros need keep).
	//dddd Range:-89~89
	//mmmmmm Range:0~999999
	//Positive value means North latitude.
	//Negative value means South latitude.
	//E.g ‘-80.590920’ stand for N 80.590920
	W string
	S int // Ground speed. Unit:0.01Km/h

	//Ground course range:0~35999
	// Unit.0.01 degree offset of the North in the clock wise.
	C int
	//14 Bytes
	//Local time in yyyymmddhhmmss format.
	//E.g :20120928121212 stand for 12:12:12 at 28 of September 2012..
	T string
	H int //altitude by sea level
}

type DSMAlarmParameter struct {
	ALARMNAME string
	SER       string //Abbreviation of the name of the alarm
	ST        SubAlarmTypeDSM
	LEV       byte //Alarm level 1. First Level Alarm 2. Second Level Alarm
	LCH       uint //Alarm linkage channel, Bit0~31 for CH1~CH32, when bit==1 will be valid.
	//When Lane departure warning was triggered,
	//0—means left Lane departure;
	//1—means right lane departure.
	LDWTYPE byte
	KW      string //Reserved
}

type SubAlarmTypeDSM uint

var SubAlarmTypeDSMEnum = struct {
	FatigueDriving          SubAlarmTypeDSM
	NoDriver                SubAlarmTypeDSM
	DriverMakingPhoneCall   SubAlarmTypeDSM
	DriverSmoking           SubAlarmTypeDSM
	DriverDistraction       SubAlarmTypeDSM
	LaneDeviation           SubAlarmTypeDSM
	FrontCarCollision       SubAlarmTypeDSM
	PreAlarmForOverSpeed    SubAlarmTypeDSM
	LicensePlateRecognition SubAlarmTypeDSM
	FollowDistanceWarning   SubAlarmTypeDSM
	Yawning                 SubAlarmTypeDSM
	PedestrianDetection     SubAlarmTypeDSM
	NoSeatbeltAlarm         SubAlarmTypeDSM
	FacialRecognitionFailed SubAlarmTypeDSM
	StopSignDetectionAlarm  SubAlarmTypeDSM
}{
	FatigueDriving:          0,
	NoDriver:                1,
	DriverMakingPhoneCall:   2,
	DriverSmoking:           3,
	DriverDistraction:       4,
	LaneDeviation:           5,
	FrontCarCollision:       6,
	PreAlarmForOverSpeed:    7,
	LicensePlateRecognition: 8,
	FollowDistanceWarning:   9,
	Yawning:                 10,
	PedestrianDetection:     11,
	NoSeatbeltAlarm:         16,
	FacialRecognitionFailed: 17,
	StopSignDetectionAlarm:  31,
}

type FireBoxAlarmParameter struct {
	//Fire Box Alarm Subtype
	//0-31: Detector 1 - detector 32 alarm
	//32: fire box short circuit alarm
	//33: fire box open circuit alarm
	SNO uint
	LCH uint                    //The channel number of the linked video, bit represent, bit value is 1, valid when field SNO is 0-31
	SP  FireBoxDetectorPosition //Detector position, field SNO is valid at 0-31
	STY FireBoxDetectorType     //Detector type, field SNO is valid at 0-31
	STE FireBoxDetectorStatus   //Detector status, field SNO is valid at 0-31
}

type FireBoxDetectorStatus byte

var FireBoxDetectorStatusEnum = struct {
	Offline   FireBoxDetectorStatus
	Normal    FireBoxDetectorStatus
	Failure   FireBoxDetectorStatus
	Pollution FireBoxDetectorStatus
	Alarm     FireBoxDetectorStatus
	Isolated  FireBoxDetectorStatus
}{
	Offline:   0,
	Normal:    1,
	Failure:   2,
	Pollution: 3,
	Alarm:     4,
	Isolated:  5,
}

type FireBoxDetectorType byte

var FireBoxDetectorTypeEnum = struct {
	NotInstalled                 FireBoxDetectorType
	MixedWithSmokeAndTemperature FireBoxDetectorType
	HighTemperature              FireBoxDetectorType
	Flame                        FireBoxDetectorType
	TemperatureSensitiveCable    FireBoxDetectorType
	SmokeSense                   FireBoxDetectorType
	SystemReserved               FireBoxDetectorType
}{
	NotInstalled:                 0,
	MixedWithSmokeAndTemperature: 1,
	HighTemperature:              2,
	Flame:                        3,
	TemperatureSensitiveCable:    4,
	SmokeSense:                   5,
	SystemReserved:               6,
}

type FireBoxDetectorPosition byte

var FireBoxDetectorPositionEnum = struct {
	None                           FireBoxDetectorPosition
	TheTopOfTheFristEndDriverRoom  FireBoxDetectorPosition
	TheTopOfTheSecondEndDriverRoom FireBoxDetectorPosition
	FristEndForTheCabinetInside    FireBoxDetectorPosition
	SecondEndForTheCabinetInside   FireBoxDetectorPosition
	TopOfMechanicalRoom1           FireBoxDetectorPosition
	TopOfMechanicalRoom2           FireBoxDetectorPosition
	TopOfMechanicalRoom3           FireBoxDetectorPosition
	TopOfMechanicalRoom4           FireBoxDetectorPosition
	CorridorFloorTrunking          FireBoxDetectorPosition
	TopOfTheElectricRoom           FireBoxDetectorPosition
	TopOfTheElectricRoom1          FireBoxDetectorPosition
	TopOfTheElectricRoom2          FireBoxDetectorPosition
	TheTopOfTheElectricRoom        FireBoxDetectorPosition
	TheTopOfTheElectricRoom1       FireBoxDetectorPosition
	TheTopOfTheElectricRoom2       FireBoxDetectorPosition
}{
	None:                           0,
	TheTopOfTheFristEndDriverRoom:  1,
	TheTopOfTheSecondEndDriverRoom: 2,
	FristEndForTheCabinetInside:    3,
	SecondEndForTheCabinetInside:   4,
	TopOfMechanicalRoom1:           5,
	TopOfMechanicalRoom2:           6,
	TopOfMechanicalRoom3:           7,
	TopOfMechanicalRoom4:           8,
	CorridorFloorTrunking:          9,
	TopOfTheElectricRoom:           10,
	TopOfTheElectricRoom1:          11,
	TopOfTheElectricRoom2:          12,
	TheTopOfTheElectricRoom:        13,
	TheTopOfTheElectricRoom1:       14,
	TheTopOfTheElectricRoom2:       15,
}
