package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"os"
	"time"
)

//ini文件中的字段顺序需要与结构体中的顺序一致
type App struct {
	JwtSecret string
	PageSize int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath string
	ImageMaxSize int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt string
	TimeFormat string
}

type Server struct {
	RunMode string
	HttpPort int
	ReadTimeOut time.Duration
	WriteTimeOut time.Duration
}

type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
	TablePrefix string
}

type Redis struct {
	Host string
	Password string
	MaxIdle int
	MaxActive int
	IdleTimeout time.Duration
}

var (
	AppSetting = &App{}
	ServerSetting = &Server{}
	DatabaseSetting = &Database{}
	RedisSetting = &Redis{}
	AppPath = "admin"
	cfg *ini.File
)

func Setup() {
	var err error
	cfg, err = ini.Load(fmt.Sprintf("%s/src/%s/conf/app.ini",os.Getenv("GOPATH"),AppPath))
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini' %v\n", err)
	}
	mapTo("app",AppSetting)
	mapTo("server",ServerSetting)
	mapTo("database",DatabaseSetting)
	mapTo("redis",RedisSetting)

	AppSetting.RuntimeRootPath = AppSetting.RuntimeRootPath
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	//移除这里 会有一定概率 造成访问超时未响应
	//TimeOut 的单位为ns
	ServerSetting.WriteTimeOut = ServerSetting.WriteTimeOut * time.Second
	ServerSetting.ReadTimeOut = ServerSetting.ReadTimeOut * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

func mapTo(section string,v interface{}){
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("cfg.MapTO %s error:%v\n",section, err)

	}
}