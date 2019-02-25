package models

type MovieRate struct{
	RateId int `gorm:"primary_key,AUTO_INCREMENT" json:"rate_id"`
	DouId int `gorm:"unique" json:"dou_id"`
	MovieId int `gorm:"unique" json:"movie_id"`
	CommentTotal int `json:"comment_total"`
	CommentValue float64 `json:"comment_value"`
}

func AddRate(rate MovieRate) int {
	if id := MovieIsComment(rate.DouId); id> 0 {
		return id
	}
	db.Create(rate)
	return rate.DouId
}

func MovieIsComment(dou_id int) int {
	rate := MovieRate{}
	db.Where("dou_id = ?",dou_id).First(&rate)
	return rate.RateId
}