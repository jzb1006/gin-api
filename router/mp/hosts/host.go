package hosts

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/controller/mp/hosts"
)

func HostRouter(r *gin.RouterGroup)  {
	ApiPrefix := r.Group("api")
	{
		ApiPrefix.GET("/hosts",hosts.Index)
	}
}
