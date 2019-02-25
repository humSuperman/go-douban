package models

import "fmt"

type Item struct {
	ItemId int `gorm:"primary_key,auto_increment" json:"item_id"`
	ItemName string `json:"item_name"`
	ItemQuota int `json:"item_quota"`
	CreateTime int `json:"create_time"`
}
//添加item
func AddItem(item Item) int{
	if item_id := GetItemIdByName(item.ItemName); item_id > 0 {
		return item_id
	}
	db.Create(&item)
	return item.ItemId
}
//检查item是否存在
func IsExistItem(ItemName string) bool {
	count := 0
	db.Model(Item{}).Where("item_name = ?",ItemName).Count(&count)
	return count > 0
}
//根据item名获取item_id
func GetItemIdByName(ItemName string) int {
	var item = Item{}
	db.Where("item_name = ?",ItemName).First(&item)
	return item.ItemId
}

func GetItemName(InWhere []int){
	var items []Item
	db.Where("item_id in (?)",InWhere).Find(&items)
	for _,val := range items {
		fmt.Println(val)
	}
}

