package admin

import (
	"admin/models"
	"admin/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Tags struct {

}

func AddTag(c *gin.Context){
	name := c.Query("name")
	_,err := models.AddTag(name)
	data := make(map[string]interface{})
	code := e.SUCCESS
	msg := ""
	if err > 0 {
		code = err
		msg = e.GetMsg(code)
	}else{
		msg = "新增标签成功"
	}
	c.JSON(http.StatusOK,gin.H{
		"code" : code,
		"msg" : msg,
		"data" : data,
	})
}

func (t *Tags)EditTag(){

}

func (t *Tags)DelTag(){

}
