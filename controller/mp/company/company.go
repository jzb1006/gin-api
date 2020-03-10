package company

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/model"
	"go-gin-api/util"
	"net/http"
)

func Index(c *gin.Context) {
	res := util.Gin{Ctx: c}
	data, err := new(model.SubCompany).Index(c)
	if err == nil {
		res.Response(http.StatusOK, data)
	}
}
