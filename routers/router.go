package routers

import (
	"admin/http/controller/admin"
	"admin/http/controller/home"
	"admin/middleware/jwt"
	"admin/pkg/setting"
	"admin/pkg/upload"
	"admin/routers/api"
	"admin/routers/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine{
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	r.GET("/",v1.Index)
	r.GET("/auth", api.GetAuth)
	api_v1 := r.Group("/api")
	{
		api_v1.GET("/",v1.Index)
		api_v1.GET("/index",v1.Index)
		api_v1.GET("/job",v1.Job)
		api_v1.GET("/articles",v1.Articles)
		//api_v1.GET("/about",v1.About)
	}
	api_v1.Use(jwt.JWT())
	{
		api_v1.GET("/about",v1.About)
	}
	admin_route := r.Group("/admin")
	{
		admin_route.GET("/tags/add",admin.AddTag)
	}
	home_route := r.Group("/home")
	{
		home_route.GET("/upload",home.UploadImage)
		home_route.GET("/movie", home.MovieRecommend)
	}

	return r
}
