package douban

import (
	"admin/models"
	"admin/pkg/file"
	"admin/pkg/http"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type MemberMovie struct {
	Count int `json:"count"`
	Total int `json:"total"`
	Start int `json:"start"`
}

func WapRequest(){

	header := map[string]string{
		"Accept":"application/json",
		"User-Agent":"Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Mobile Safari/537.36",
		"X-Requested-With":"XMLHttpRequest",
		"Referer":"",
	}
	GoPath := os.Getenv("GOPATH")
	for {
		Link := models.GetNotUseMemberLink()
		p := 0
		header["Referer"] = fmt.Sprintf("https://m.douban.com/people/%s/movie/done",Link)
		for {
			url := fmt.Sprintf("https://m.douban.com/rexxar/api/v2/user/%s/interests?type=movie&status=done&start=%d&count=50&ck=aegX&for_mobile=1",Link,p*50)
			fmt.Println(url)
			path := fmt.Sprintf( "%s/src/admin/json_movie/%s_%d.json",GoPath,Link,p)
			if !file.CheckExist(path){
				p++
				continue
			}
			res := http.HttpHandle(url,header)
			//fmt.Printf("%s\n",res)

			var movie MemberMovie
			json.Unmarshal(res, &movie)
			fmt.Println(movie)
			if movie.Total == 0 {
				break
			}
			file.WriteFile(path,res)
			if movie.Start+movie.Count >= movie.Total {
				break
			}
			p++
			time.Sleep(1*time.Second)
		}

	}
	fmt.Println("Success")
}

//想看
//https://m.douban.com/people/48475163/movie/todo

//看过
//https://m.douban.com/people/auringonpaiste/movie/done

//电影详情
//https://m.douban.com/movie/subject/27109679/

//电影评论列表
//https://m.douban.com/rexxar/api/v2/movie/27109679/interests?count=20&order_by=hot&start=25&ck=&for_mobile=1

//电影演职人员
//https://m.douban.com/rexxar/api/v2/movie/26894670/credits