package dto

const (
	ImportantEventAI AlarmImportance = iota
	GeneralAlarmAI
	EmergencyAlarmAI
)

const (
	SimplifiedChinese Lang = iota
	English
	Korean
	Italian
	German
	Thai
	Turkey
	Portugal
	Spain
	Romania
	Greece
	French
	Russian
	Dutch
	Hebrew
	ChineseTraditional
)
const SUB_STREAM_N9M = 0
const MAIN_STREAM_N9M = 1
const SUBSTREAM = "substream"
const STREAM = "stream"
const SCREENSHOT = "screenshot"
