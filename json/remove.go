package main

import (
	"admin/models"
	"admin/pkg/file"
	"admin/pkg/setting"
	"fmt"
	"os"
)

func main() {
	setting.Setup()
	models.Setup()
	user := models.GetMemberList()
	for _,val := range user {
		for i:=0;i<=20000;i++ {
			path := fmt.Sprintf( "D:/json_movie/%s_%d.json",val.MemberLink,i)
			if file.CheckExist(path){
				break
			}
			fmt.Println(path)
			os.Remove(path)
		}
	}
}
