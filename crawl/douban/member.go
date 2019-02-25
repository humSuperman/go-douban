package douban

import (
	"admin/models"
	"fmt"
	"admin/pkg/http"
	"os"
	"regexp"
	"bytes"
	"time"
	"admin/pkg/file"
)

func MemberMovies(){
	PublicUrl := "https://movie.douban.com/people/%s/collect?start=%d"
	hreader := map[string]string{
		"Accept":"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
		"Cache-Control":"max-age=0",
		"Connection":"keep-alive",
		"Host":"movie.douban.com",
		"User-Agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36 OPR/57.0.3098.116",

	}
	GoPath := os.Getenv("GOPATH")
	for {
		memberLink := models.GetNotUseMemberLink()
		if memberLink == "" {
			break
		}
		p := 0
		for {
			// /data/golang/src/admin/data/html/member_movie
			// D:/Go/gin-admin/html/member_movie
			path := fmt.Sprintf("%s/src/admin/data/html/member_movie/%s_%d.html",GoPath,memberLink,p)
			if !file.CheckExist(path){
				p++
				continue
			}
			data := http.HttpHandle(fmt.Sprintf(PublicUrl,memberLink,p*15),hreader)
			res := MemberMoviesHtml(data)
			if bytes.Count([]byte(res),nil) <= 1 {
				models.SetMemberIsUse(memberLink)
				break
			}

			file.WriteFile(path,res)
			p++
			time.Sleep(50*time.Second)
		}

	}
	fmt.Println("SUCCESS")
}

func MemberMoviesHtml(data []byte) []byte{
	content := `<div class="grid-view">[\s\S[:^ascii:]]+<div class="aside">`
	regx := regexp.MustCompile(content)
	match := regx.Find(data)

	item := regexp.MustCompile(`<div class="item"`)
	count := item.FindAll(match,1)
	for _,val := range count {
		if bytes.Count([]byte(val),nil) > 1 {
			return match
		}
	}
	return nil
}
