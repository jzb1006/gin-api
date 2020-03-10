package track

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/model"
	"go-gin-api/request"
	"go-gin-api/util"
	"go-gin-api/util/jwt"
	"net/http"
)


func Share(c *gin.Context) {
	res := util.Gin{Ctx: c}
	validate := new(request.TrackShare).Validator(c) //验证数据
	if validate  {
		return
	}
	form := request.TrackShare{}
	err := c.ShouldBind(&form)
	if err != nil {
		res.Response(http.StatusUnprocessableEntity, err.Error())
		return
	}
	auth := jwt.GetAuthInfo(c)
	if err := new(model.TrackShare).Share(auth.Model.ID, form.Type, form.Id); err != nil {
		res.Response(http.StatusBadRequest, err.Error())
		return
	}
	res.Response(http.StatusCreated, "")
}
