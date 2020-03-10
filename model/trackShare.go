package model

import (
	"github.com/jinzhu/gorm"
	"go-gin-api/init/mysql"
	"go-gin-api/util/encode"
)

type TrackShare struct {
	Id            int              `json:"id"`
	UserId        uint             `json:"user_id"`
	EventReportId int              `json:"event_report_id"`
	Type          int              `json:"type"`
	CreatedAt     encode.LocalTime `json:"created_at"`
	UpdatedAt     encode.LocalTime `json:"updated_at"`
}

func (t *TrackShare) Share(UserId uint, Type, EventReportId int) (err error) {
	if err = mysql.DEFAULTDB.Create(&TrackShare{EventReportId: EventReportId, UserId: UserId, Type: Type}).Error; err == nil {
		if Type == 2 {
			mysql.DEFAULTDB.Model(&Report{}).Where("id = ? ", EventReportId).UpdateColumn("share_count", gorm.Expr("share_count + ?", 1))
		}
	}
	return
}
