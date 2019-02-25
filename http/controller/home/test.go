package home

import (
	"admin/models"
	"admin/pkg/file"
	"fmt"
)

func UpdateMember(){
	List := models.GetUseMemberLink()
	for _,val := range List {
		path := "D:/Go/gin-admin/src/admin/html/" + val.MemberLink +"_0.html"
		fmt.Println(val.Id,val.MemberLink)
		if file.CheckExist(path) {
			fmt.Println(val.Id,val.MemberLink)
			models.SetMemberIsNotUse(val.Id)
		}
	}
}
