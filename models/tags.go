package models

import (
	"admin/pkg/e"
	"fmt"
)

type Tags struct{
	TagsId int `gorm:"primary_key,AUTO_INCREMENT" json:"tags_id"`
	TagsName string `json:"tags_name"`
	CreateTime int `json:"create_time"`
}
//添加标签tag
func AddTag(name string) (TagId,err int){
	id := GetTagsByName(name)
	if id > 0{
		return 0,e.TAGS_IS_EXIST
	}
	TagData := Tags{TagsName:name}
	db.Create(&TagData)
	if TagData.TagsId > 0 {
		return TagData.TagsId,0
	}else{
		return 0,e.TAGS_INSERT_FAIL
	}
}
//通过name获取tag_id
func GetTagsByName(name string) int{
	tags := Tags{}
	db.Where("tags_name = ?",name).First(&tags)
	fmt.Println(tags)
	if tags.TagsId > 0 {
		return tags.TagsId
	}
	return 0
}

func DelTag(id int){
	db.Where("tags_name = ?",id).Delete(Tags{})
}
