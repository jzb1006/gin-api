package hosts

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-api/model"
	"go-gin-api/util"
	"net/http"
)

func Index(c *gin.Context) {
	var utilGin = util.Gin{Ctx: c}
	err, data := new(model.Host).Index()
	fmt.Println(err, data)
	if err != nil {
		utilGin.Response(http.StatusBadRequest, err.Error())
	}
	utilGin.Response(http.StatusOK, data)
	return
}
