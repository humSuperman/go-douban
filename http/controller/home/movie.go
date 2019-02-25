package home

import (
	"admin/models"
	"admin/pkg/gredis"
	"admin/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
	"time"
)

//Item CF
func MovieRecommend(c *gin.Context){

	t := time.Now()
	gredis.ZSet("zadd","value-2",1,20)
	movie_id := c.Query("movie")
	id,_ := strconv.Atoi(movie_id)

	//1.获取电影详情及关联的item 用户添加&豆瓣添加
	//info := models.GetMovieInfo(id)
	ItemSlice := models.GetMovie2Item(id)
	//fmt.Println(ItemSlice)

	//获取item名
	//models.GetItemName(ItemSlice)

	//2.获取item集合关联的所以电影
	MoviesToItem := models.GetMoviesListByItems(ItemSlice,id)

	BaseItemMap := make(map[int]int)
	for _,val := range ItemSlice{
		BaseItemMap[val] = 1
	}
	//3.求item相似度
	LikeScore := util.ItemCF(MoviesToItem,BaseItemMap)

	//4.相似度正序排列 返回相似度最高的30
	var ReScore util.PersonSlice
	for key,val := range LikeScore {
		rs := util.Person{
			Keys:key,
			Val:val,
		}
		ReScore = append(ReScore,rs)
		gredis.RedisConn.Get()
	}
	sort.Stable(ReScore)
	res := ReScore[len(ReScore)-30:]

	//5.获取推荐电影id集合
	var rs_mid []int
	for _,val := range res {
		rs_mid = append(rs_mid,val.Keys)
	}
	//6.获取相似度 item 的所以电影
	rs := models.GetMoviesList(rs_mid,"id,movie_name,movie_rate")
	fmt.Println(time.Since(t))
	c.JSON(http.StatusOK,gin.H{
		"code":"200",
		"msg":"success",
		"data":rs,
	})
}
//User CF
func RecommendUser(c *gin.Context){
	movie_id := c.Query("movie")
	id,_ := strconv.Atoi(movie_id)
	//1.获取电影详情及关联的item 用户添加&豆瓣添加
	info := models.GetMovieInfo(id)
	ItemSlice := models.GetMovie2Item(id)
	fmt.Println(ItemSlice)

	//获取item名
	//models.GetItemName(ItemSlice)

	//2.获取观看过此电影的所有item
	models.GetUserItemListByMovie(id,info.MovieRate)

	//3.求item相似度

	//4.获取相似度 item 的所以电影
	//models.GetMoviesByMember2Item(users)

	//fmt.Println(users)
	c.JSON(http.StatusOK,gin.H{
		"code":"200",
		"msg":"success",
		"data":"",
	})
}