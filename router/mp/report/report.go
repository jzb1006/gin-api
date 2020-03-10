package report

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/controller/mp/report"
	"go-gin-api/router/middleware/jwt"
)

func ReportRouter(c *gin.RouterGroup) {
	EventApi := c.Group("api/report/")
	{
		EventApi.GET("", report.Index)
		EventApi.GET(":id", report.Show)
		UserAuthApi := EventApi.Use(jwt.JWT())
		{
			UserAuthApi.GET(":id/favorite", report.Fave)
			UserAuthApi.DELETE(":id/favorite", report.UnFave)
			UserAuthApi.GET(":id/is-favorite", report.IsFave)
			UserAuthApi.POST(":id/view", report.ViewHistoryStore)
		}
	}
}
