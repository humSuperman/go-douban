package main

import (
	"admin/models"
	"admin/pkg/gredis"
	"admin/pkg/logging"
	"admin/pkg/setting"
	"admin/routers"
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"net/http"
	"time"
)
type Server struct {
	RunMode	string
	HttpPort int
	ReadTimeOut time.Duration
	WriteTimeOut time.Duration
}

var ServerSetting = &Server{}

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeOut,
		WriteTimeout:   setting.ServerSetting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
func Demo(){
	Cfg,err := ini.Load("src/admin/conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini' : %v\n",err)
	}

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSeting err %v\n",err)
	}
	fmt.Println(ServerSetting)
}