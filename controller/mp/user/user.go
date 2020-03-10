package user

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/entity/user"
)

func Login(c *gin.Context) {
	user.Login(c)
}

func Info(c *gin.Context)  {
	user.Info(c)
}