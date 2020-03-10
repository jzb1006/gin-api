package request

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/model"
	"go-gin-api/util"
	"go-gin-api/util/validate"
	"net/http"
)

type TrackShare struct {
	Type int `json:"type" form:"type" validate:"required"`
	Id   int `json:"id" form:"id" validate:"required"`
}

func exist(t *TrackShare, c *gin.Context) (b bool) {
	b = false
	res := util.Gin{Ctx: c}
	if t.Type == 1 {
		if _, err := new(model.Event).Show(t.Id); err != nil {
			res.Response(http.StatusUnprocessableEntity, "event不存在")
			b = true
		}
	} else {
		if _, err := new(model.Report).Show(t.Id); err != nil {
			res.Response(http.StatusUnprocessableEntity, "report不存在")
			b = true
		}
	}
	return
}

func (t *TrackShare) Validator(c *gin.Context) (b bool) {
	c.ShouldBind(&t)
	field := validate.BaseField{Field: t}
	b = field.Validator(c)
	if !b {
		b = exist(t, c)  //自定义验证贵族
	}
	return
}
