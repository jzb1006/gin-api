package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-api/request"
	"go-gin-api/service/wechat"
	"go-gin-api/util"
	"go-gin-api/util/jwt"
	"math"
	"net/http"
	"time"
)

var identityKey = "id"

type Token struct {
	Token  string
	Expire int
}

func Login(c *gin.Context) {
	var token string
	var expireTime time.Time
	res := util.Gin{Ctx: c}
	user := &request.User{}
	if err := c.ShouldBind(&user); err != nil {
		res.Response(http.StatusBadRequest, err.Error())
		return
	}
	data, _ := wechat.Login(user.Code)
	if data.OpenID == "" {
		res.Response(http.StatusBadRequest, data.ErrMSG)
		return
	}
	userInfo, err := wechat.Exist(data.OpenID)
	fmt.Println(userInfo, err)
	if err != nil {
		userInfo, err = wechat.Save(data.OpenID, data.SessionKey, data.UnionID)
		if err != nil {
			res.Response(http.StatusBadRequest, err.Error())
			return
		}
	}
	token, expireTime, err = jwt.GenerateToken(userInfo)

	res.Response(http.StatusOK, Token{Token: token, Expire: int(math.Floor(expireTime.Sub(time.Now()).Minutes()))})
	return
}

func Info(c *gin.Context) {
	res := util.Gin{Ctx: c}
	Auth := jwt.GetAuthInfo(c)
	res.Response(http.StatusOK, Auth)
}
