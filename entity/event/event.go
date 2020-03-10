package event

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/model"
	"go-gin-api/request"
	"go-gin-api/util"
	"go-gin-api/util/jwt"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	res := util.Gin{Ctx: c}
	Search := request.EventRequest{}
	c.ShouldBind(&Search)
	data, meta, err := new(model.Event).Index(c, Search.Search, Search.HostId, Search.TypeId, Search.Time)
	if err != nil {
		res.Response(http.StatusBadRequest, err.Error())
		return
	}
	res.ResponsePaging(http.StatusOK, data, meta)
}

func Show(c *gin.Context) {
	res := util.Gin{Ctx: c}
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := new(model.Event).Show(id)
	if err != nil {
		res.Response(http.StatusBadRequest, err.Error())
		return
	}
	res.Response(http.StatusOK, data)
}

func Fave(c *gin.Context) {
	res := util.Gin{Ctx: c}
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := new(model.Event).Show(id)
	if err != nil {
		res.Response(http.StatusBadRequest, err.Error())
		return
	}
	Auth := jwt.GetAuthInfo(c)
	err = new(model.EventFavorite).Fave(Auth.Model.ID, data.Id)
	res.Response(http.StatusCreated, "")
}

func UnFave(c *gin.Context) {
	res := util.Gin{Ctx: c}
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := new(model.Event).Show(id)
	if err != nil {
		res.Response(http.StatusBadRequest, err.Error())
		return
	}
	Auth := jwt.GetAuthInfo(c)
	err = new(model.EventFavorite).UnFave(Auth.Model.ID, data.Id)
	res.Response(http.StatusCreated, "")
}

func FaveHistory(c *gin.Context) {
	res := util.Gin{Ctx: c}
	Auth := jwt.GetAuthInfo(c)
	data, mate, err := new(model.EventFavorite).FaveHistory(c, Auth.Model.ID)
	if err != nil {
		res.Response(http.StatusBadRequest, err.Error())
		return
	}
	res.ResponsePaging(http.StatusOK, data, mate)
}

func Types(c *gin.Context) {
	res := util.Gin{Ctx: c}
	data, err := new(model.Type).Index()
	if err != nil {
		res.Response(http.StatusBadRequest, err.Error())
		return
	}
	res.Response(http.StatusOK, data)
}

func Enter(c *gin.Context) {
	res := util.Gin{Ctx: c}
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := new(model.Event).Show(id)
	if err != nil {
		res.Response(http.StatusBadRequest, err.Error())
		return
	}
	Auth := jwt.GetAuthInfo(c)
	err = new(model.EventEnter).Enter(Auth.Model.ID, data.Id)
	res.Response(http.StatusCreated, "")
}

func UnEnter(c *gin.Context) {
	res := util.Gin{Ctx: c}
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := new(model.Event).Show(id)
	if err != nil {
		res.Response(http.StatusBadRequest, err.Error())
		return
	}
	Auth := jwt.GetAuthInfo(c)
	err = new(model.EventEnter).UnEnter(Auth.Model.ID, data.Id)
	res.Response(http.StatusCreated, "")
}

func EnterHistory(c *gin.Context) {
	res := util.Gin{Ctx: c}
	Auth := jwt.GetAuthInfo(c)
	data, mate, err := new(model.EventEnter).EnterHistory(c, Auth.Model.ID)
	if err != nil {
		res.Response(http.StatusBadRequest, err.Error())
		return
	}
	res.ResponsePaging(http.StatusOK, data, mate)
}
