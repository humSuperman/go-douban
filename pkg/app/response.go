package app

import (
	"admin/pkg/e"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(HttpCode,ErrCode int,data interface{}){
	g.C.JSON(HttpCode,gin.H{
		"code":ErrCode,
		"msg":e.GetMsg(ErrCode),
		"data":data,
	})
}
