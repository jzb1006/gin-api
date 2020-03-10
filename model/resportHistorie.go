package model

import "go-gin-api/init/mysql"

type ReportHistorie struct {
	*BaseModelField
	Id       int `json:"id"`
	UserId   uint  `json:"user_id"`
	ReportId int  `json:"report_id"`
}

func (r *ReportHistorie) Store(UserId uint, Id int) (err error) {
	err = mysql.DEFAULTDB.Create(ReportHistorie{UserId:UserId,ReportId:Id}).Error
	return
}
