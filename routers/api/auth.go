package api

import (
	"admin/models"
	"admin/pkg/app"
	"admin/pkg/e"
	"admin/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct{
	Username string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required;MaxSize(50)"`
}

func GetAuth(c *gin.Context){
	appG := app.Gin{C:c}
	username := c.Query("username")
	password := c.Query("password")
	data := make(map[string]interface{})

	valid := validation.Validation{}
	valid.Required(username,"username").Message("请输入用户名")
	valid.Required(password,"password").Message("请输入密码")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, data)
		return
	}

	isExist := models.CheckAuth(username,password)
	if !isExist {
		appG.Response(http.StatusOK, e.ERROR_AUTH, data)
		return
	}

	token,err := util.GenerateToken(username,password)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_AUTH_TOKEN, data)
		return
	}
	data["token"] = token
	appG.Response(http.StatusOK, e.SUCCESS, data)
	return
}