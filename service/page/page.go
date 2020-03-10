package page

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-gin-api/util"
)

type Model struct {
	Model interface{}
}

type PageRequest struct {
	Page int `form:"page" json:"page"`
	Size int `form:"size" json:"size"`
}

func (m Model) Paginator(db *gorm.DB, c *gin.Context) (data *gorm.DB, meta util.Meta) {
	pageInfo := PageRequest{}
	c.ShouldBind(&pageInfo)
	size := pageInfo.Size
	page := pageInfo.Page
	if size == 0 {
		size = 5
	}
	if page == 0 {
		page = 1
	}
	db.Find(m.Model).Count(&meta.Total)
	meta.TotalPage = (meta.Total + size - 1) / size
	meta.CurrentPage = page
	data = db.Limit(size).Offset((page - 1) * size)
	return
}
