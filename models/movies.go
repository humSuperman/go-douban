package models

type Movies struct{
	ID int `gorm:"primary_key,AUTO_INCREMENT" json:"id"`
	DouId int `json:"dou_id"`
	MovieName string `json:"movie_name"`
	MovieUrl string `json:"movie_url"`
	MovieType int `json:"movie_type"`
	MovieRate float64 `json:"movie_rate"`
	RateNum int `json:"rate_num"`
	CreateTime int `json:"create_time"`
}

//添加电影
func AddMove(movie Movies) int {
	if MoveExist(movie.DouId){
		return GetMovieIdByDouId(movie.DouId)
	}
	db.Create(movie)
	return movie.ID
}

//检查电影是否存在
func MoveExist(dou_id int) bool {
	count := 0
	db.Model(&Movies{}).Where("dou_id = ?",dou_id).Count(&count)
	return count > 0
}

//通过豆瓣电影id获取数据库自定义id
func GetMovieIdByDouId(dou_id int) int {
	var movie Movies
	db.Where("dou_id = ?",dou_id).First(&movie)
	return movie.ID
}

//获取电影详情
func GetMovieInfo(id int) Movies{
	var movie Movies
	db.Where("",id).First(&movie)
	return movie
}

//获取电影列表
func GetMoviesList(InWhere []int,field string) []Movies{
	var Movie []Movies
	db.Where("id in (?)",InWhere).Select(field).Find(&Movie)
	return Movie
}

