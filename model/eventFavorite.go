package model

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/init/mysql"
	"go-gin-api/service/page"
	"go-gin-api/util"
	"go-gin-api/util/encode"
	"time"
)

type EventFavorite struct {
	Id        int              `json:"id"`
	UserId    uint             `json:"user_id"`
	EventId   int              `json:"event_id"`
	CreatedAt encode.LocalTime `json:"created_at"`
	UpdatedAt encode.LocalTime `json:"updated_at"`
	DeletedAt *time.Time       `sql:"index" json:"-"`
}

//添加用户收藏
func (EF *EventFavorite) Fave(UserId uint, EventId int) (err error) {
	err = mysql.DEFAULTDB.Where("user_id = ? and event_id = ?", UserId, EventId).Find(&EventFavorite{}).Error
	if err != nil {
		err = mysql.DEFAULTDB.Create(&EventFavorite{UserId: UserId, EventId: EventId}).Error
	}
	return
}

//取消用户收藏
func (EF *EventFavorite) UnFave(UserId uint, EventId int) (err error) {
	err = mysql.DEFAULTDB.Where("user_id = ? and event_id = ?", UserId, EventId).Delete(&EventFavorite{}).Error
	return
}

//分页获取所有收藏
func (EF *EventFavorite) FaveHistory(c *gin.Context, UserId uint) (History []EventFavorite, meta util.Meta, err error) {
	data := mysql.DEFAULTDB.Where("user_id = ?", UserId)
	data, meta = page.Model{Model: &History}.Paginator(data, c)
	err = data.Find(&History).Error
	return
}