package models

import (
	"github.com/jinzhu/gorm"
)

type Item2movie struct {
	Id int `gorm:"AUTO_INCREMENT,primary_key" json:"id"`
	MovieId int `json:"movie_id"`
	ItemId int `json:"item_id"`
	UseNum int `json:"use_num"`
	ItemType int `json:"item_type"`
	CreateTime int `json:"create_time"`
}
//添加movie对应item数据
func AddMovieItem(items Item2movie) int{
	if id := CheckMovieItemExist(items.MovieId,items.ItemId); id > 0 {
		db.Model(&Item2movie{}).Where("id = ?",id).Update("use_num", gorm.Expr("use_num + ?", 1))
		return id
	}else{
		db.Create(&items)
	}
	return items.ItemId
}
//检查 movie与item关系是否存在
func CheckMovieItemExist(m_id,i_id int) int{
	var items Item2movie
	db.Where("movie_id = ? AND item_id = ?",m_id,i_id).First(&items)
	return items.Id
}

//获取 movie对应item集合
func GetMovie2Item(movie_id int) []int{
	var item []Item2movie
	db.Where("movie_id = ?",movie_id).Select("item_id,use_num").Find(&item)
	var ItemSlice []int
	for _,val := range item {
		//item使用次数小于5 不做统计
		if val.ItemId > 0 && val.UseNum >= 5 {
			ItemSlice = append(ItemSlice,val.ItemId)
		}
	}
	return ItemSlice
}

//获取所有电影分别对应item集合
func GetMoviesListByItems(BaseItem []int,movie_id int) map[int][]int {
	//1.获取基础item集合关联的电影
	var RsMovies []Item2movie
	db.Where("item_id IN (?) AND use_num >= 5",BaseItem).Select("movie_id").Find(&RsMovies)

	//2.获取关联电影集合对应的的item
	var movies_id []int
	map_movies_id := make(map[int]int)
	for _,val := range RsMovies {
		//电影id与条件id相同 不做统计
		if val.MovieId == 0 || val.MovieId == movie_id {
			continue
		}
		if _,ok := map_movies_id[val.MovieId]; !ok {
			movies_id = append(movies_id,val.MovieId)
			map_movies_id[val.MovieId] = 0
		}
	}
	var ItemMovie []Item2movie
	db.Where("movie_id IN (?) AND use_num >= 5",movies_id).Select("movie_id,item_id").Find(&ItemMovie)
	MoviesToItem := make(map[int][]int)
	for _,val := range ItemMovie {
		MoviesToItem[val.MovieId] = append(MoviesToItem[val.MovieId],val.ItemId)
	}

	return MoviesToItem
}
