package jwt

import "github.com/gin-gonic/gin"

func GetToken(c *gin.Context) (token string) {
	token = c.GetHeader("Authorization")
	return
}
