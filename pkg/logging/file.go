package logging

import (
	"admin/pkg/file"
	"admin/pkg/setting"
	"fmt"
	"os"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt = "log"
	TimeFormat = "20060102"
)

func getLogFilePath() string {
	//dir, _ := os.Getwd()
	//path := dir + "/src/admin/" + LogSavePath
	return fmt.Sprintf("%s%s",setting.AppSetting.RuntimeRootPath,setting.AppSetting.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt)
}

func getLogFileFullPath() string{
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s",LogSaveName,time.Now().Format(TimeFormat),LogFileExt)

	return fmt.Sprintf("%s%s",prefixPath,suffixPath)
}

func openLogFile(filename,filePath string) (*os.File,error){
	dir,err := os.Getwd()
	if err != nil {
		return nil,fmt.Errorf("os.Getwd() error:%v\n",err)
	}
	src := dir + "/src/admin/" + filePath
	perm := file.CheckPermission(src)
	if perm == true {
		return nil,fmt.Errorf("file.CheckPremission premission denied src:%s\n",src)
	}

	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil,fmt.Errorf("file.IsNotExistDir src:%s; error:%v\n",src,err)
	}

	f,err := file.Open(src+filename,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil {
		return nil,fmt.Errorf("Fail to OpenFile :%v\n",err)
	}
	return f,nil
}

func mkDir(){
	dir:= getLogFilePath()//php中的getcwd(),获取main文件的根目录地址
	err := os.MkdirAll(dir,os.ModePerm)
	if err != nil {
		panic(err)
	}
}
