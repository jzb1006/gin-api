package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)



type Config struct {
	MysqlAdmin   MysqlAdmin
	Log Log
	ErrorMailSend
	App
	WechatMP
	Token
}

//数据库配置
type MysqlAdmin struct { // mysql admin 数据库配置
	Username string
	Password string
	Path     string
	Dbname   string
	Config   string
}

//错误日志配置
type Log struct {
	Path string
	Name string
	ErrorPath string
	ErrorName string
}

//软件出现Bug告警邮箱
type ErrorMailSend struct {
	Active bool
	User string
	Pass string
	Host string
	Port int
}
//软件信息配置
type App struct {
	Name string
	Version string
}

type WechatMP struct {
	AppId string
	AppsSecret string
}

type Token struct {
	ExpireTime time.Duration
}

var GinConfig Config

func init() {
	v := viper.New()
	v.SetConfigName("config")           //  设置配置文件名 (不带后缀)
	v.AddConfigPath("./static/config/") // 第一个搜索路径
	v.SetConfigType("json")
	err := v.ReadInConfig() // 搜索路径，并读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	if err := v.Unmarshal(&GinConfig); err != nil {
		fmt.Println(err)
	}
}

