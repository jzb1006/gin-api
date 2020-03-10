package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-gin-api/config"
	"go-gin-api/model"
	"go-gin-api/util"
	"net/http"
	"strings"
	"time"
)

var jwtSecret = []byte("dsfgh8dg798dug89dsf89gd7g98sd79f")

type Claims struct {
	model.User
	jwt.StandardClaims
}

func GenerateToken(user model.User) (string,time.Time, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(config.GinConfig.Token.ExpireTime * time.Hour)
	claims := Claims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginApp",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, expireTime, err
}

func ParseToken(AuthHeader string) (*Claims, error) {
	tmp := strings.Split(AuthHeader, " ")
	token := tmp[1]
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}


func GetAuthInfo(c *gin.Context) (user model.User) {
	token := GetToken(c)
	Auth, err := ParseToken(token)
	res := util.Gin{Ctx: c}
	if err != nil {
		res.Response(http.StatusUnauthorized, "Unauthorized")
	}
	user = Auth.User
	return
}
