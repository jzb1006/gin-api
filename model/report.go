package model

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/init/mysql"
	"go-gin-api/service/page"
	"go-gin-api/util"
	"go-gin-api/util/encode"
	"time"
)

type Report struct {
	Id            int              `json:"id"`
	Name          string           `json:"name"`
	ImgUrl        string           `json:"img_url"`
	PostTime      encode.LocalTime `json:"post_time"`
	Author        string           `json:"author"`
	HostId        int              `json:"host_id"`
	HostNae       string           `json:"host_name"`
	Keywords      string           `json:"keywords"`
	Detail        string           `json:"detail"`
	ViewCount     int              `json:"view_count"`
	FavoriteCount int              `json:"favorite_count"`
	ShareCount    int              `json:"share_count"`
	Host          Host             `json:"host" gorm:"foreignkey:HostId"`
	CreatedAt     encode.LocalTime `json:"created_at"`
	UpdatedAt     encode.LocalTime `json:"updated_at"`
	DeletedAt     *time.Time       `sql:"index" json:"-"`
}

func (r *Report) Index(c *gin.Context, search string, hostId int) (report []Report, meta util.Meta, err error) {
	data := mysql.DEFAULTDB
	if search != "" {
		data = data.Where("name LIKE ?", "%"+search+"%")
	}

	if hostId > 0 {
		data = data.Where("host_id = ?", hostId)
	}

	data, meta = page.Model{Model: &report}.Paginator(data, c) //抽象出分页器
	data = data.Order("post_time asc")
	err = data.Preload("Host").Find(&report).Error //Preload关联host表
	return
}

func (r *Report) Show(Id int) (report Report, err error) {
 	err = mysql.DEFAULTDB.Where("id = ?", Id).Preload("Host").First(&report).Error
 	return
}

