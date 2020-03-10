package util

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	Ctx *gin.Context
}

type Meta struct {
	Total int `json:"total"`
	CurrentPage int `json:"current_page"`
	TotalPage int `json:"total_page"`
}
type responsePage struct {
	Data    interface{} `json:"data"`
	Meta Meta `json:"meta"`
}

type response struct {
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(httpCode int, data interface{}) {
	g.Ctx.JSON(httpCode, response{ Data: data})
	return
}

func (g *Gin) ResponsePaging(httpCode int, data interface{},meta Meta) {
	g.Ctx.JSON(httpCode, responsePage{ Data: data,Meta:meta})
	return
}


