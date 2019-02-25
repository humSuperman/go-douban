package main

import (
	"admin/crawl/douban"
	"admin/models"
	"admin/pkg/setting"
)

func main() {
	setting.Setup()
	models.Setup()
	douban.MdxRequest()
	//douban.WapRequest()
	//douban.MemberMovies()
}
