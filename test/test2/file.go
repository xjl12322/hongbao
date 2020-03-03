package main

import (
	"os"
	"path"
	"time"
)
import "fmt"

type FileLogger struct {
	Level LogLevel
	filePath string //
	fileName string
	fileObj *os.File
	errfileObj *os.File
	maxFileSize int64
}

func NewFileLogger(levelStr,fp,fn string,maxSize int64)*FileLogger  {
	logLevel,err := parseLogLevel(levelStr)
	if err != nil{
		panic(err)
	}
	fl := &FileLogger{
		Level:logLevel,
		filePath:fp,
		fileName:fn,
		maxFileSize:maxSize,
	}
	err= fl.initFile() //初始化打开文件
	if err != nil {
		panic(err)
	}
	return fl

}
func (l *FileLogger)initFile()(error){
	fullFileName := path.Join(l.filePath,l.fileName)
	fileObj,err := os.OpenFile(fullFileName,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)

	if err != nil{
		fmt.Printf("open log file failed,err%v\n",err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed, err:%v\n", err)
		return err
	}
	l.fileObj = fileObj
	l.errfileObj = errFileObj
	return nil
}



func (l *FileLogger)enable(logLevel LogLevel) bool {
	return l.Level<=logLevel

}
func (l *FileLogger)checkSize(file *os.File) bool {
	fileInfo,err := file.Stat()
	if err != nil{
		fmt.Printf("get file info failed,err:%v\n",err)
		return false

	}
	return fileInfo.Size() >= l.maxFileSize

}
// 切割文件
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 需要切割日志文件
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())      // 拿到当前的日志文件完整路径
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr) // 拼接一个日志文件备份的名字
	// 1. 关闭当前的日志文件
	file.Close()
	// 2. 备份一下 rename  xx.log  ->  xx.log.bak201908031709
	os.Rename(logName, newLogName)
	// 3. 打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed, err:%v\n", err)
		return nil, err
	}
	// 4. 将打开的新日志文件对象赋值给 f.fileObj
	return fileObj, nil
}
func (l *FileLogger)log(lv LogLevel,format string, a ...interface{})  {
	if l.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		funcName, fileName, lineNo := getInfo(3)
		now := time.Now().Format("2016-01-02 15:04:05")
		if l.checkSize(l.fileObj) {
			newFile, err := l.splitFile(l.fileObj) // 日志文件
			if err != nil {
				return
			}
			l.fileObj = newFile
		}
		fmt.Fprintf(l.fileObj,"[%s] [%s][%s  %s-%d] %s\n", now, getLogString(lv), fileName, funcName, lineNo, msg)
		if lv >= ERROR{ // 如果要记录的日志大于等于ERROR级别,我还要在err日志文件中再记录一遍
			fmt.Fprintf(l.fileObj,"[%s] [%s][%s  %s-%d] %s\n", now, getLogString(lv), fileName, funcName, lineNo, msg)
		}

	}
}

func (l *FileLogger)Debug(format string, a ...interface{})  {
	l.log(DEBUG,format,a...)
}

func (l *FileLogger)Info(format string, a ...interface{})  {
	l.log(INFO,format,a...)
}
func (l *FileLogger)Warning(format string, a ...interface{})  {
	l.log(WARNING,format,a...)
}
func (l *FileLogger)Error(format string, a ...interface{})  {
	l.log(ERROR,format,a...)
}
func (l *FileLogger)Fatal(format string, a ...interface{})  {
	l.log(FATAL,format,a...)

}


func (l *FileLogger)Close(){

	l.fileObj.Close()
	l.errfileObj.Close()
}
















