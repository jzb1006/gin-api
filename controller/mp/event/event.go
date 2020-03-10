package event

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/entity/event"
)

func Index(c *gin.Context) {
	event.Index(c)
}

func Show(c *gin.Context) {
	event.Show(c)
}

func Fave(c *gin.Context) {
	event.Fave(c)
}

func UnFave(c *gin.Context) {
	event.UnFave(c)
}

func FaveHistory(c *gin.Context) {
	event.FaveHistory(c)
}

func Type(c *gin.Context) {
	event.Types(c)
}

func Enter(c *gin.Context) {
	event.Enter(c)
}

func UnEnter(c *gin.Context) {
	event.UnEnter(c)
}

func EnterHistory(c *gin.Context) {
	event.EnterHistory(c)
}
