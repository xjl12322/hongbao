package main

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

func getLogString(lv LogLevel)string {
	switch lv {

	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return ""

}

func getInfo(skip int) (funcName,fileName string,lineNo int) {
	pc,file,lineNo,ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.caller() failed\n")
		return
	}
	fileName = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(fileName,".")[0]
	return

}





