package track

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/controller/mp/track"
	"go-gin-api/router/middleware/jwt"
)

func TrackRouter(r *gin.RouterGroup) {
	TApi := r.Group("api/track")
	TApiJwt := TApi.Use(jwt.JWT())
	{
		TApiJwt.POST("share", track.Share)
	}
}
