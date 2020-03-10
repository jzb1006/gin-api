package user

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/controller/mp/event"
	"go-gin-api/controller/mp/report"
	"go-gin-api/controller/mp/user"
	"go-gin-api/router/middleware/jwt"
)

func UserRouter(r *gin.RouterGroup) {
	UserApi := r.Group("api/user")
	{
		UserApi.GET("/login", user.Login)
	}
	UserAuthApi := UserApi.Use(jwt.JWT())
	{
		UserAuthApi.GET("/update", user.Info)
		UserAuthApi.GET("events/favorite", event.FaveHistory)
		UserAuthApi.GET("events/enters", event.EnterHistory)
		UserAuthApi.GET("reports/favorite", report.FaveHistory)
		UserAuthApi.GET("reports/view", report.ViewHistory)
	}
}
