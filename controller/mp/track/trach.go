package track

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/entity/track"
)

func Share(c *gin.Context)  {
	track.Share(c)
}