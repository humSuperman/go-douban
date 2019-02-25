package main

import (
	"admin/models"
	"admin/pkg/file"
	"admin/pkg/setting"
	"admin/pkg/times"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)
type Bodys struct {
	Count int `json:"count"`
	Start int `json:"start"`
	Total int `json:"total"`
	Interests []Interests `json:"interests"`
	User string
}
type Interests struct {
	Comment string `json:"comment"`
	Rating Rat `json:"rating"`
	Tags []string `json:"tags"`
	Platforms []string `json:"platforms"`
	VoteCount int `json:"vote_count"`
	CreateTime string `json:"create_time"`
	Status string `json:"status"`
	Id int `json:"id"`
	IsPrivate bool `json:"is_private"`
	Subject Movie `json:"subject"`
}
type Movie struct {
	Rating Rat `json:"rating"`
	Genres []string `json:"genres"`
	Pubdate []string `json:"pubdate"`
	HasLinewatch bool `json:"has_linewatch"`
	Url string `json:"url"`
	Title string `json:"title"`
	ReleaseDate string `json:"release_date"`
	Pic pic `json:"pic"`
	Subtype string `json:"subtype"` //子类型 moive-电影 tv-电视剧
	Directors []Name  `json:"directors"` //导演
	IsShow bool `json:"is_show"`
	Actors []Name `json:"actors"`
	IsReleased bool `json:"is_released"`
	Year string `json:"year"`
	Type string `json:"type"`
	Id string `json:"id"`
}
type Name struct {
	Name string `json:"name"`
}
type pic struct {
	Large string `json:"large"`		//大图
	Normal string `json:"normal"`	//小图
}
type Rat struct {
	Count int `json:"count"`
	Max float64 `json:"max"`
	StarCount float64 `json:"star_count"`
	Value float64 `json:"value"`
}

type Items struct{
	ItemId int
	ValId int
	ValType int
}

func main() {
	setting.Setup()
	models.Setup()
	user := models.GetUseMemberLink()
	//初始化channel
	var UserChannel = make(chan string,10)
	//var JsonChannel = make(chan []byte,10)
	var ArrayChannel = make(chan Bodys,10)

	//var ItemChannel = make (chan Items,10)

	go CreateUserChann(user,UserChannel)
	//1.读取文件
	go ReadUserJson(UserChannel,ArrayChannel)
	//2.将文件内容加入 转换json goruntine
	//go JsonToArray(JsonChannel,ArrayChannel)
	//3.读取ArrayChannel
	go ReadArrayChann(ArrayChannel)

	//go ItemsRelational(ItemChannel)
	//4.将电影详情加入mysql数据库 并将平台id与电影评价 标签信息分别加入goruntine
		//电影评价goruntine
		//电影标签goruntine
		//演职人员goruntine
	time.Sleep(186400*time.Second)

	//Test()
}
//创建 user channel
func CreateUserChann(user []models.Member,UserChannel chan string){
	for _,val := range user {
		UserChannel <- val.MemberLink
	}

}
//读取json文件
func ReadUserJson(UserChannel chan string,ArrayChannel chan Bodys){
	for {
		user := <- UserChannel
		for i:=0; i>=0; i++ {
			path := fmt.Sprintf("F:/json_movie/%s_%d.json",user,i)
			//fmt.Println(path)
			if !file.CheckExist(path) {
				f,err := ReadFile(path)
				if err == nil {

					JsonToArray(f,user,ArrayChannel)
					//jsonString <- f
				}
			}else{
				break
			}
		}

	}

}
//读取文件
func ReadFile(path string) ([]byte,error) {
	f,err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Errorf("%s",err)
		return nil,err
	}
	//fmt.Printf("%s\n", f)
	//JsonToArray(f)
	return f,nil
}
//将json字符串转为数组
func JsonToArray(f []byte,user string,ArrayChannel chan Bodys){
	//for {
		var Array Bodys
		//
		// f := <- StringChann
		//fmt.Printf("%s\n", f)
		err := json.Unmarshal(f,&Array)
		if err != nil {
			fmt.Errorf("%s",err)
		}
		Array.User = user
		ArrayChannel <- Array
	//}
	/**
	for _,val := range Array.Interests {
		fmt.Println(val.Subject.Year)
		//fmt.Println(val.Tags)
	}*/
	//return Array,nil
	//fmt.Println(Array.Interests)
}
//读取json数组channel
//并完成数据入库
func ReadArrayChann(ArrayChannel chan Bodys){
	for {
		array := <-ArrayChannel
		user_id := models.GetMemberInfoByLink(array.User)
		for _,val := range array.Interests {
			douID,_ := strconv.Atoi(val.Subject.Id)

			MovieType := 2
			if val.Subject.Subtype== "movie" {
				MovieType = 1
			}
			//电影数据入库
			movie := models.Movies{
				DouId:douID,
				MovieName:val.Subject.Title,
				MovieType:MovieType,
				MovieUrl:val.Subject.Pic.Large,
				MovieRate:val.Subject.Rating.Value,
				RateNum:val.Subject.Rating.Count,
			}
			movie_id := models.AddMove(movie)

			//演员入库
			//演员 电影关系入库
			//二者在另一个接口中获取并写入 演员/导演的其他信息

			//观影记录入库 [暂无想看这个版块，移放在下一版本]

			//用户评价标签数据入库
			if !models.MemberIsComment(user_id,movie_id) {
				comment_item := ""
				for _,val := range val.Tags {
					item := models.Item{
						ItemName:val,
					}
					item_id := models.AddItem(item)
					user_item := models.Item2member{
						ItemId:item_id,
						MemberId:user_id,
						UseNum:1,
					}
					id := models.AddUserItem(user_item)
					comment_item = fmt.Sprintf("%s%d,",comment_item,id)
					movie_item := models.Item2movie{
						ItemId:item_id,
						MovieId:movie_id,
						UseNum:1,
						ItemType:2,
					}
					models.AddMovieItem(movie_item)
					fmt.Println(id,2)
				}

				//评价信息入库 需要带用户评价item
				comment := models.MovieComment{
					MovieId:movie_id,
					MemberId:user_id,
					Content:val.Comment,
					Score:val.Rating.Value,
					Likes:val.VoteCount,
					Items:strings.TrimRight(comment_item,","),
					CreateTime:times.StrToIntTime(val.CreateTime,"s"),
				}
				models.AddComment(comment)
				//电影标签数据入库
				for _,val := range val.Subject.Genres {
					item := models.Item{
						ItemName:val,
					}
					item_id := models.AddItem(item)
					movie_item := models.Item2movie{
						ItemId:item_id,
						MovieId:movie_id,
						UseNum:1,
						ItemType:1,
					}
					id := models.AddMovieItem(movie_item)
					fmt.Println(id,1)
				}
			}
			//电影评分数据入库

			/*
			if movie_id <= 0 {
				fmt.Printf("movie rate install error: movie=%d",douID)
				continue
			}
			rate := models.MovieRate{
				DouId:douID,
				MovieId:movie_id,
				CommentTotal:val.Subject.Rating.Count,
				CommentValue:val.Subject.Rating.Value,
			}
			if models.AddRate(rate) <= 0 {
				fmt.Printf("movie rate install error: movie=%d",douID)
				continue
			}*/
			//fmt.Println("success")
		}
	}
}
func Test(){
	//4.
	path := "D:/Go/gin-admin/src/admin/json_movie/y328675858_86.json"
	f,err := ioutil.ReadFile(path)//os.Open(path)
	if err != nil {
		fmt.Println(err)
		panic("error")
	}
	//JsonToArray(f)
	fmt.Printf("bytes: %s\n", f)
}

