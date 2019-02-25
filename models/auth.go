package models

type Auth struct{
	ID int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username,password string) bool{
	var auth Auth
	db.Select("id,password").Where(Auth{Username:username}).First(&auth)
	if auth.ID>0 && auth.Password==password {
		return true
	}
	return false
}
