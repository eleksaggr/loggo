package loggo

import (
	"io"
	"log"
	"time"
)

type Level uint

const (
	INFO Level = iota
	DEBUG
	WARNING
	ERROR
)

var levelNames map[Level]string = map[Level]string {
	INFO: "INFO",
	DEBUG: "DEBUG",
	WARNING: "WARN",
	ERROR: "ERROR",
}

var levelColors map[Level]string = map[Level]string {
	INFO: "\033[0m",
	DEBUG: "\033[94m",
	WARNING: "\033[93m",
	ERROR: "\033[91m",
}

type Log struct {
	Logger *log.Logger
}

type Record struct {
	Message string
	Time time.Time
	Level Level
}

func NewLog(out io.Writer) *Log {
	return &Log{Logger : log.New(out, "", 0)}
}

func (l *Log) log(record *Record) {
	time := record.Time.Format("15:04:05")
	l.Logger.Printf("%s%s [%s] %s%s",levelColors[record.Level], time, levelNames[record.Level], record.Message, "\033[0m")
}

func (l *Log) Info(message string) {
	l.log(&Record{Message: message, Time: time.Now(), Level: INFO})
}

func (l *Log) Debug(message string) {
	l.log(&Record{Message: message, Time: time.Now(), Level: DEBUG})
}

func (l *Log) Warning(message string){
	l.log(&Record{Message: message, Time: time.Now(), Level: WARNING})
}

func (l *Log) Error(message string){
	l.log(&Record{Message: message, Time: time.Now(), Level: ERROR})
}
