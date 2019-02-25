package models

import (
	"admin/pkg/logging"
	"admin/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var db *gorm.DB

func Setup(){
	var err error
	conf := setting.DatabaseSetting
	db,err = gorm.Open(conf.Type,fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Name))

	if err!= nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB,defaultTableName string) string{
		return conf.TablePrefix + defaultTableName
	}

	//db.Callback().Create().Replace("gorm:update_time_stamp", createTimeStampForCreateCallBack)
	//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallBack)
	//db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	db.LogMode(true)
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
//数据库创建数据回调
func createTimeStampForCreateCallBack(scope *gorm.Scope){
	logging.Info("create call back")
	if !scope.HasError() {
		nowTime := time.Now().Unix()

		if createTimeField,ok := scope.FieldByName("CreateTime");ok {
			if createTimeField.IsBlank{
				createTimeField.Set(nowTime)
			}
		}
	}
}
//数据库修改数据回调
func updateTimeStampForUpdateCallBack(scope *gorm.Scope){
	logging.Info("update call back")
	if !scope.HasError() {
		nowTime := time.Now().Unix()

		if createTimeField,ok := scope.FieldByName("UpdateTime");ok {
			if createTimeField.IsBlank{
				createTimeField.Set(nowTime)
			}
		}
	}
}
//删除回调
func deleteCallback(){
	//数据软删除使用
	logging.Info("update call back")
}