package model

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/init/mysql"
	"go-gin-api/service/page"
	"go-gin-api/util"
)

type ReportHistories struct {
	*BaseModelField
	Id       int    `json:"id"`
	UserId   uint   `json:"user_id"`
	ReportId int    `json:"report_id"`
	Report   Report `json:"report"`
}

func (r *ReportHistories) Store(UserId uint, Id int) (err error) {
	if err = mysql.DEFAULTDB.Where("user_id = ? and report_id = ?", UserId, Id).First(&ReportHistories{}).Error; err != nil {
		err = mysql.DEFAULTDB.Create(&ReportHistories{UserId: UserId, ReportId: Id}).Error
	}
	return
}

func (r *ReportHistories) Index(c *gin.Context, UserId uint) (h []ReportHistories, meta util.Meta, err error) {
	data := mysql.DEFAULTDB.Where("user_id = ? ", UserId)
	data, meta = page.Model{Model: &h}.Paginator(data, c)
	err = data.Preload("Report").Find(&h).Error
	return
}
