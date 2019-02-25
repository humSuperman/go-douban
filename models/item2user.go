package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"math"
)

type Item2member struct {
	Id int `gorm:"AUTO_INCREMENT,primary_key" json:"id"`
	MemberId int `json:"member_id"`
	ItemId int `json:"item_id"`
	UseNum int `json:"use_num"`
	CreateTime int `json:"create_time"`
}
//添加 user对应item关系
func AddUserItem(items Item2member) int{
	if id := CheckUserItemExist(items.MemberId,items.ItemId); id > 0 {
		db.Model(&Item2member{}).Where("id = ?",id).Update("use_num", gorm.Expr("use_num + ?", 1))
		return id
	}else{
		db.Create(&items)
	}
	return items.ItemId
}
//检查 user对应item是否存在
func CheckUserItemExist(m_id,i_id int) int{
	var items Item2member
	db.Where("member_id = ? AND item_id = ?",m_id,i_id).First(&items)
	return items.Id
}

func GetUserItemListByMovie(movie_id int,rate float64){
	var item_to_member []Item2member
	var ViewsList []MovieComment
	var MemberList []int
	db.Where("movie_id = (?)",movie_id).Select("member_id,score").Find(&ViewsList)
	for _,val := range ViewsList{
		if math.Abs(val.Score*2-rate) <= 3 {
			MemberList = append(MemberList,val.MemberId)
		}
	}
	db.Where("member_id in (?) AND use_num >=5",MemberList).Select("item_id,use_num").Find(&item_to_member)
	fmt.Println(item_to_member)
}
