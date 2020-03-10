package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-api/init/mysql"
	"go-gin-api/service/page"
	"go-gin-api/util"
	"go-gin-api/util/encode"
	"time"
)

type ReportFavorite struct {
	Id        int              `json:"id"`
	UserId    uint             `json:"user_id"`
	ReportId  int              `json:"report_id"`
	CreatedAt encode.LocalTime `json:"created_at"`
	UpdatedAt encode.LocalTime `json:"updated_at"`
	DeletedAt *time.Time       `sql:"index" json:"-"`
}

//添加用户收藏
func (RF *ReportFavorite) Fave(UserId uint, ReportId int) (err error) {
	err = mysql.DEFAULTDB.Where("user_id = ? and report_id = ?", UserId, ReportId).Find(&ReportFavorite{}).Error
	if err != nil {
		err = mysql.DEFAULTDB.Create(&ReportFavorite{UserId: UserId, ReportId: ReportId}).Error
	}
	return
}

//取消用户收藏
func (RF *ReportFavorite) UnFave(UserId uint, ReportId int) (err error) {
	err = mysql.DEFAULTDB.Where("user_id = ? and report_id = ?", UserId, ReportId).Delete(&ReportFavorite{}).Error
	return
}

//分页获取所有收藏
func (RF *ReportFavorite) FaveHistory(c *gin.Context, UserId uint) (History []ReportFavorite, meta util.Meta, err error) {
	data := mysql.DEFAULTDB.Where("user_id = ?", UserId)
	data, meta = page.Model{Model: &History}.Paginator(data, c)
	err = data.Find(&History).Error
	return
}

func (RF *ReportFavorite) IsFave(UserId uint, ReportId int) (b bool, err error) {
	b = true
	err = mysql.DEFAULTDB.Where("user_id = ? and report_id = ?", UserId, ReportId).First(&ReportFavorite{}).Error
	fmt.Println(err)
	if err != nil {
		b = false
	}
	return
}
