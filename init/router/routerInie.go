package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/router/middleware"
	"go-gin-api/router/mp/company"
	"go-gin-api/router/mp/event"
	"go-gin-api/router/mp/hosts"
	"go-gin-api/router/mp/report"
	"go-gin-api/router/mp/track"
	"go-gin-api/router/mp/user"
)

func RouterInit() *gin.Engine {
	var r = gin.Default()
	r.Use(middleware.LoggerToFile())
	Api := r.Group("")
	hosts.HostRouter(Api)
	company.CompanyRouter(Api)
	event.EventRouter(Api)
	report.ReportRouter(Api)
	user.UserRouter(Api)
	track.TrackRouter(Api)
	return r
}
