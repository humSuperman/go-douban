package models

type MovieRateUser struct{
	Id int `gorm:"primary_key,AUTO_INCREMENT" json:"id"`
	UserId int `json:"user_id"`
	DouId int `json:"dou_id"`
	MovieId int `json:"movie_id"`
	CommentValue float64 `json:"comment_value"`
	CommentDesc string `json:"comment_desc"`
	CreateTime int `json:"create_time"`
}

func AddRateUser(rate MovieRateUser) int {
	if id := UserIsComment(rate.DouId,rate.UserId); id> 0 {
		return id
	}
	db.Create(rate)
	return rate.Id
}

func UserIsComment(dou_id,user_id int) int {
	rate := MovieRateUser{}
	db.Where("user_id = ? and dou_id = ?",user_id,dou_id).First(&rate)
	return rate.Id
}