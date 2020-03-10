package model

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/init/mysql"
	"go-gin-api/util"
	"net/http"
)

type SubCompany struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

func (c SubCompany) Index(ctx *gin.Context)(company []SubCompany, err error) {
	res := util.Gin{Ctx: ctx}
	err = mysql.DEFAULTDB.Find(&company).Error
	if err != nil {
		res.Response(http.StatusBadRequest,err.Error())
		return
	}
	return
}

