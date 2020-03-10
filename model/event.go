package model

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/init/mysql"
	page2 "go-gin-api/service/page"
	"go-gin-api/util"
	"go-gin-api/util/encode"
	"time"
)

type Event struct {
	Id            int              `json:"id"`
	Name          string           `json:"name"`
	CoverImgUrl   string           `json:"cover_img_url"`
	DetailImgUrl  string           `json:"detail_img_url"`
	StartDate     encode.LocalTime `json:"start_date"`
	EndDate       encode.LocalTime `json:"end_date"`
	Location      string           `json:"location"`
	Outer         string           `json:"outer"`
	Cost          string           `json:"cost"`
	TypeId        int              `json:"type_id"`
	Type          int              `json:"type"`
	Instruction   string           `json:"instruction"`
	Agenda        string           `json:"agenda"`
	ViewCount     int              `json:"view_count"`
	FavoriteCount int              `json:"favorite_count"`
	CreatedAt     encode.LocalTime `json:"created_at"`
	DeletedAt     time.Time        `json:"-" gorm:"null"` //查询时过滤删除的数据
	HostId        int              `json:"host_id"`
	Host          Host             `json:"host" gorm:"foreignkey:HostId"`
	Types         Type             `json:"types" gorm:"foreignkey:TypeId"`
}

func (e *Event) Index(c *gin.Context, search string, host_id int, type_id int, date string) (event []Event, meta util.Meta, err error) {
	data := mysql.DEFAULTDB
	if search != "" {
		data = data.Where("name LIKE ?", "%"+search+"%")
	}
	if host_id < 0 {
		data = data.Where("host_id = ?", -1)
	} else if host_id > 0 {
		data = data.Where("host_id = ?", host_id)
	}
	if type_id > 0 {
		data = data.Where("type_id = ?", type_id)
	}
	if date == "recent" {
		data = data.Where("end_date <= ?", time.Now().AddDate(0, 1, 0)).Where("end_date >= ?", time.Now())
	} else if date == "past" {
		data = data.Where("end_date <= ?", time.Now())
	}
	data, meta = page2.Model{Model: &event}.Paginator(data, c) //抽象出分页器

	data = data.Order("start_date asc")
	err = data.Preload("Host").Preload("Types").Find(&event).Error //Preload关联host表
	for i,_ := range event{
		event[i].Type = 1
	}
	return
}

func (e Event) Show(id int) (event Event, err error)  {
	data := mysql.DEFAULTDB
	err = data.Where("id = ?",id).Preload("Host").Preload("Types").Find(&event).Error
	event.Type = 1
	return
}



