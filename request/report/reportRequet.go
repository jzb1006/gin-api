package report

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/model"
	"go-gin-api/util"
	"go-gin-api/util/validate"
	"net/http"
	"strconv"
)

type ReportRequest struct {
	Search string `form:"search" json:"search"`
	HostId int    `form:"host_id" json:"host_id"`
	Id     int    `json:"id" form:"id" validate:"required"`
}

func exist(c *gin.Context, Id int) (b bool) {
	b = false
	res := util.Gin{Ctx: c}
	if _, err := new(model.Report).Show(Id); err != nil {
		res.Response(http.StatusUnprocessableEntity, "report不存在")
		b = true
	}
	return
}

func (t *ReportRequest) Validator(c *gin.Context) (formDate *ReportRequest, b bool) {
	c.ShouldBind(&t)
	formDate = t
	t.Id, _ = strconv.Atoi(c.Param("id"))
	field := validate.BaseField{Field: t}
	b = field.Validator(c)
	if !b {
		b = exist(c, t.Id)
	}
	return
}
