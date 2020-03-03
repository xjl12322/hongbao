package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type LogLevel uint16

const (
	UNKNOWN  LogLevel=iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL

)

type Logger struct {
	Level LogLevel

}

func parseLogLevel(s string)(LogLevel,error){
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG,nil
	case "trace":
		return TRACE,nil
	case "info":
		return INFO,nil
	case "warning":
		return WARNING,nil
	case "error":
		return ERROR,nil
	case "fatal":
		return FATAL,nil
	default:
		err := errors.New("无效的级别日志")
		return UNKNOWN,err
	}
}


func NewLog(levelStr string) Logger{
	level,err := parseLogLevel(levelStr)
	if err != nil{
		panic(err)
	}
	return Logger{Level:level,}
}
func (l Logger)enable(logLevel LogLevel) bool {
	return l.Level<=logLevel

}


func (l Logger)log(lv LogLevel,format string, a ...interface{})  {
	if l.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		funcName, fileName, lineNo := getInfo(3)
		now := time.Now().Format("2016-01-02 15:04:05")
		fmt.Printf("[%s] [%s][%s  %s-%d] %s\n", now, getLogString(lv), fileName, funcName, lineNo, msg)
	}
}

func (l Logger)Debug(format string, a ...interface{})  {
	l.log(DEBUG,format,a...)
}

func (l Logger)Info(format string, a ...interface{})  {
	l.log(INFO,format,a...)
}
func (l Logger)Warning(format string, a ...interface{})  {
	l.log(WARNING,format,a...)
}
func (l Logger)Error(format string, a ...interface{})  {
	l.log(ERROR,format,a...)
}
func (l Logger)Fatal(format string, a ...interface{})  {
	l.log(FATAL,format,a...)

}