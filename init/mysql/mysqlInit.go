package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-gin-api/config"
	"log"
)

var DEFAULTDB *gorm.DB

//初始化数据库并产生数据库全局变量
func MysqlInit(admin config.MysqlAdmin) *gorm.DB {
	if db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config); err != nil {
		log.Printf("DEFAULTDB数据库启动异常%S", err)
	} else {
		DEFAULTDB = db
		DEFAULTDB.DB().SetMaxIdleConns(10)
		DEFAULTDB.DB().SetMaxOpenConns(100)
	}
	return DEFAULTDB
}