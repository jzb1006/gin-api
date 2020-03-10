package event

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/controller/mp/event"
	"go-gin-api/router/middleware/jwt"
)

func EventRouter(c *gin.RouterGroup) {
	TypesApi := c.Group("api/")
	{
		TypesApi.GET("types", event.Type)
	}
	EventApi := c.Group("api/events/")
	{
		EventApi.GET("", event.Index)
		EventApi.GET(":id", event.Show)
		UserAuthApi := EventApi.Use(jwt.JWT())
		{
			UserAuthApi.GET(":id/favorite", event.Fave)
			UserAuthApi.DELETE(":id/favorite", event.UnFave)
			UserAuthApi.GET(":id/enter", event.Enter)
			UserAuthApi.DELETE(":id/enter", event.UnEnter)
		}
	}
}
