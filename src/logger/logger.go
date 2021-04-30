package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	//Error key
	Error = "error"

	//Info key
	Info = "info"

	//Fatal key
	Fatal = "fatal"
)

var staticLogger *logger

//Logger ...
func Logger() *logger {
	if staticLogger == nil {
		staticLogger = buildLogger()
	}
	return staticLogger
}

func buildLogger() *logger {
	log.SetFlags(log.Ldate | log.Ltime)

	errorOut := &lumberjack.Logger{
		Filename:   logFilePath("error.log"),
		MaxSize:    500, // megabytes
		MaxBackups: 6,
		MaxAge:     7,    //days
		Compress:   true, // disabled by default
	}
	infoOut := &lumberjack.Logger{
		Filename:   logFilePath("info.log"),
		MaxSize:    500, // megabytes
		MaxBackups: 6,
		MaxAge:     7,    //days
		Compress:   true, // disabled by default
	}
	fatalOut := &lumberjack.Logger{
		Filename:   logFilePath("fatal.log"),
		MaxSize:    500, // megabytes
		MaxBackups: 6,
		MaxAge:     7,    //days
		Compress:   true, // disabled by default
	}
	return &logger{
		ErrorOutput: errorOut,
		InfoOutput:  infoOut,
		FatalOutput: fatalOut,
	}
}

func logFilePath(_type string) string {
	dir := filepath.Dir(os.Args[0])
	absPath, _ := filepath.Abs(dir + "/logs/" + _type)
	return absPath
}

//Logger logger
type logger struct {
	ErrorOutput *lumberjack.Logger
	InfoOutput  *lumberjack.Logger
	FatalOutput *lumberjack.Logger
}

//WriteToLog write content to log
func (l *logger) WriteToLog(logType string, content ...interface{}) {
	logContent := fmt.Sprintln(content...)
	switch logType {
	case Error:
		log.SetOutput(l.ErrorOutput)
	case Info:
		log.SetOutput(l.InfoOutput)
	default:
		{
			log.SetOutput(l.FatalOutput)
			log.Fatalf(fmt.Sprintf("[%v]", logType), logContent)
			return
		}
	}
	log.Print(fmt.Sprintf("[%v]", logType), logContent)
}
