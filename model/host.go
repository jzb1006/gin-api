package model

import (
	"go-gin-api/init/mysql"
	"go-gin-api/util/encode"
)

type Host struct {
	Id int `json:"id"`
	Name string `json:"name"`
	CreatedAt encode.LocalTime `json:"created_at" gorm:"default:'nil'"`
}

func (h *Host)Index () (err error, host []Host)  {
	err = mysql.DEFAULTDB.Find(&host).Error
	return
}

func (h *Host)Update (id int,params map[string]interface{})(err error, host *Host)  {
	err = mysql.DEFAULTDB.Model(&Host{}).Where("id = ?",id).Updates(params).Error
	return
}