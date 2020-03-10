package main

import (
	"go-gin-api/config"
	"go-gin-api/init/log"
	"go-gin-api/init/mysql"
	"go-gin-api/init/router"
	"net/http"
	"time"
)

func main() {
	log.InitLog()
	mysql.MysqlInit(config.GinConfig.MysqlAdmin) // 链接初始化数据库
	Router := router.RouterInit() //注册路由
	s := &http.Server{
		Addr:           ":8080",
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}