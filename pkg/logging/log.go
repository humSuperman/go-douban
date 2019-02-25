package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPreFix = ""
	DefaultCallerDepth = 2

	logger *log.Logger
	logPrefix = ""
	levelFlags = []string{"DEBUG","INFO","WARN","ERROR","FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

func Setup(){
	/**
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	logger = log.New(F,DefaultPreFix,log.LstdFlags)*/
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F,err = openLogFile(fileName,filePath)
	if err != nil {
		log.Fatalln(err)
	}
	logger = log.New(F,DefaultPreFix,log.LstdFlags)
}

func Debug(v ...interface{}){
	setpreFix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}){
	setpreFix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}){
	setpreFix(WARN)
	logger.Println(v)
}

func Error(v ...interface{}){
	setpreFix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}){
	setpreFix(FATAL)
	logger.Println(v)
}

func setpreFix(level Level){
	_,file,line,ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]",levelFlags[level],filepath.Base(file),line)
	}else{
		logPrefix = fmt.Sprintf("[%s]",levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
