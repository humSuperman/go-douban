package models

type MovieComment struct {
	CId int `gorm:"primary_key,AUTO_INCREMENT" json:"c_id"`
	MovieId int `json:"movie_id"`
	MemberId int `json:"member_id"`
	Content string `json:"content"`
	Score float64 `json:"score"`
	Likes int `json:"likes"`
	Items string `json:"items"`
	CreateTime int `json:"create_time"`
}

//添加电影评论 -1-已评价
func AddComment(c MovieComment) int{
	//if MemberIsComment(c.MemberId,c.MovieId){
	//	return -1
	//}
	db.Create(c)
	return c.CId
}

//用户是否评价该电影 true-评价 false-未评价
func MemberIsComment(member_id,movie_id int) bool{
	count := 0
	db.Model(MovieComment{}).Where("member_id = ? AND movie_id = ?",member_id,movie_id).Count(&count)

	return count > 0
}

//获取电影评论列表
func CommentList(where map[string]interface{},page,limit int) []MovieComment{
	var comment []MovieComment
	db.Where(where).Offset((page-1)*limit).Limit(limit).Find(comment)
	return comment
}
