package model

import "go-gin-api/init/mysql"

type Type struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (t *Type) Index() (Type []Type, err error) {
	err = mysql.DEFAULTDB.Find(&Type).Error
	return
}
