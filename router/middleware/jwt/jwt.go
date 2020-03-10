package jwt

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/util"
	"go-gin-api/util/jwt"
	"net/http"
	"time"
)

var Auth jwt.Claims

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = http.StatusOK
		token := jwt.GetToken(c)
		if token == "" {
			code = http.StatusUnauthorized
		} else {
			Auth, err := jwt.ParseToken(token)
			if err != nil {
				code = http.StatusUnauthorized
			} else if time.Now().Unix() > Auth.ExpiresAt {
				code = http.StatusUnauthorized
			}
		}

		if code != http.StatusOK {
			res := util.Gin{Ctx: c}
			res.Response(code, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
