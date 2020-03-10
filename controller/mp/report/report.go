package report

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/entity/report"
)

func Index(c *gin.Context) {
	report.Index(c)
}

func Show(c *gin.Context) {
	report.Show(c)
}

func Fave(c *gin.Context) {
	report.Fave(c)
}

func UnFave(c *gin.Context) {
	report.UnFave(c)
}

func FaveHistory(c *gin.Context) {
	report.FaveHistory(c)
}

func IsFave(c *gin.Context) {
	report.IsFave(c)
}

func ViewHistoryStore(c *gin.Context) {
	report.ViewHistory(c)
}

func ViewHistory(c *gin.Context) {
	report.ViewHistory(c)
}
