package file

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"log"
)
//获取文件大小
func GetSize(f multipart.File)(int,error){
	content,err := ioutil.ReadAll(f)
	return len(content),err
}
//获取文件后缀
func GetExt(src string) string {
	return path.Ext(src)
}
//检查文件权限 true-无权限 false-有权限
func CheckPermission(src string) bool {
	_,err := os.Stat(src)
	return os.IsPermission(err)
}
//路径是否存在 true-不存在 false-存在
func CheckExist(src string) bool {
	_,err := os.Stat(src)

	return os.IsNotExist(err)
}
//检查路径是否存在，不存在创建
func IsNotExistMkDir(src string) error {
	if exist := CheckExist(src); exist==false {
		if err := MkDir(src);err != nil{
			return err
		}
	}
	return nil
}
//递归创建文件夹
func MkDir(src string) error{
	err := os.MkdirAll(src,os.ModePerm)//0777
	if err != nil {
		return err
	}
	return nil
}
//打开文件
func Open(name string,flag int, mode os.FileMode)(*os.File,error){
	f,err := os.OpenFile(name,flag,mode)
	if err != nil{
		return nil,err
	}
	return f,nil
}

func WriteFile(path string,content []byte) bool {
	file,err := os.Create(path)
	if err != nil {
		log.Fatalf("file create error:%v",err)
		return false
	}
	if _,err := file.Write(content); err != nil {
		log.Fatalf("file write error:%v",err)
		return false
	}
	return true
}


