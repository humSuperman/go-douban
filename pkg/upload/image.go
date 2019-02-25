package upload

import (
	"admin/pkg/setting"
	"path"
	"strings"
	"admin/pkg/util"
	"mime/multipart"
	"admin/pkg/file"
	"fmt"
	"admin/pkg/logging"
	"os"
)
//获取图片完整访问URL
func GetImageFullUrl(name string) string{
	return setting.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + name
}
//获取图片路径
func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}
//获取图片名称
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name,ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}
//获取图片完整路径
func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}
//检查图片后缀
func CheckImageExt(name string) bool {
	ext := file.GetExt(name)
	ext = strings.ToLower(ext)
	AppExt := setting.AppSetting.ImageAllowExts
	for _,allowExt := range AppExt {
		if ext == allowExt {
			return true
		}
	}
	return false
}
//检查图片大小 true-符合系统设置 false-图片太大
func CheckImageSize(f multipart.File) bool {
	size,err := file.GetSize(f)
	if err != nil {
		fmt.Println(err)
		logging.Warn(err)
		return false
	}
	return size < setting.AppSetting.ImageMaxSize
}
//检查图片
func CheckImage(src string) error{
	dir,err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err:%v\n",err)
	}
	
	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err:%v\n",err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}