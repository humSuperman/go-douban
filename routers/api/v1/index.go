package v1

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context){
	c.JSON(200,gin.H{
		"msg":"this is index",
	})
}

func About(c *gin.Context){
	c.JSON(200,gin.H{
		"msg":"this is about",
	})
}

func Job(c *gin.Context){
	c.JSON(200,gin.H{
		"msg":"this is job",
	})
}

func Articles(c *gin.Context){
	c.JSON(200,gin.H{
		"msg":"this is articles",
	})
}


