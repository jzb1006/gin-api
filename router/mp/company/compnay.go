package company

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/controller/mp/company"
)

func CompanyRouter(r *gin.RouterGroup)  {
	ApiPrefix := r.Group("api")
	{
		ApiPrefix.GET("/company",company.Index)
	}
}
