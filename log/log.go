package log

import (
	"log"
	"os"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

var Level = INFO

func Print(level int, v ...interface{}) {
	if level >= Level {
		log.Println(v...)
	}
}

func Printf(level int, fmt string, v ...interface{}) {
	if level >= Level {
		log.Printf(fmt, v...)
	}
}

func Debug(v ...interface{})              { Print(DEBUG, v...) }
func Debugf(fmt string, v ...interface{}) { Printf(DEBUG, fmt, v...) }

func Info(v ...interface{})              { Print(INFO, v...) }
func Infof(fmt string, v ...interface{}) { Printf(INFO, fmt, v...) }

func Warn(v ...interface{})              { Print(WARN, v...) }
func Warnf(fmt string, v ...interface{}) { Printf(WARN, fmt, v...) }

func Error(v ...interface{})              { Print(ERROR, v...) }
func Errorf(fmt string, v ...interface{}) { Printf(ERROR, fmt, v...) }

func Fatal(v ...interface{}) {
	Print(FATAL, v...)
	os.Exit(1)
}

func Fatalf(fmt string, v ...interface{}) {
	Printf(FATAL, fmt, v...)
	os.Exit(1)
}
