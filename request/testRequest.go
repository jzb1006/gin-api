package request

import "github.com/gin-gonic/gin"

type Product struct {
	Name string `form:"name" json:"name" binding:"required"`
}

func PostRequest(c *gin.Context) (err error){
	err = c.ShouldBind(&Product{})
	return err
}
