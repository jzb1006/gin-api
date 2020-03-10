package report

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-api/model"
	"go-gin-api/request/report"
	"go-gin-api/util"
	"go-gin-api/util/jwt"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	res := util.Gin{Ctx: c}
	Search := report.ReportRequest{}
	c.ShouldBind(&Search)
	fmt.Println(Search)
	data, meta, err := new(model.Report).Index(c, Search.Search, Search.HostId)
	if err != nil {
		res.Response(http.StatusBadRequest, err.Error())
		return
	}
	res.ResponsePaging(http.StatusOK, data, meta)
}

func Show(c *gin.Context) {
	res := util.Gin{Ctx: c}
	id, err := strconv.Atoi(c.Param("id"))
	data, err := new(model.Report).Show(id)
	if err != nil {
		res.Response(http.StatusBadRequest, err.Error())
		return
	}
	res.Response(http.StatusOK, data)
}

func Fave(c *gin.Context) {
	res := util.Gin{Ctx: c}
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := new(model.Report).Show(id)
	if err != nil {
		res.Response(http.StatusBadRequest, err.Error())
		return
	}
	Auth := jwt.GetAuthInfo(c)
	err = new(model.ReportFavorite).Fave(Auth.Model.ID, data.Id)
	res.Response(http.StatusCreated, "")
}

func UnFave(c *gin.Context) {
	res := util.Gin{Ctx: c}
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := new(model.Report).Show(id)
	if err != nil {
		res.Response(http.StatusBadRequest, err.Error())
		return
	}
	Auth := jwt.GetAuthInfo(c)
	err = new(model.ReportFavorite).UnFave(Auth.Model.ID, data.Id)
	res.Response(http.StatusCreated, "")
}

func FaveHistory(c *gin.Context) {
	res := util.Gin{Ctx: c}
	Auth := jwt.GetAuthInfo(c)
	data, mate, err := new(model.ReportFavorite).FaveHistory(c, Auth.Model.ID)
	if err != nil {
		res.Response(http.StatusBadRequest, err.Error())
		return
	}
	res.ResponsePaging(http.StatusOK, data, mate)
}

func IsFave(c *gin.Context) {
	res := util.Gin{Ctx: c}
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := new(model.Report).Show(id)
	if err != nil {
		res.Response(http.StatusBadRequest, err.Error())
		return
	}
	Auth := jwt.GetAuthInfo(c)
	IsFave, err := new(model.ReportFavorite).IsFave(Auth.Model.ID, data.Id)
	res.Response(http.StatusCreated, IsFave)
}

func HistoryStore(c *gin.Context)  {
	res := util.Gin{Ctx: c}
	formDate, b := new(report.ReportRequest).Validator(c);if b {
		return
	}
	fmt.Println(formDate)
	//auth := jwt.GetAuthInfo(c)
	//err := new(model.ReportHistorie).Store(auth.Model.ID,1)
	res.Response(http.StatusCreated, "")
}
