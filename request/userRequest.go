package request

type User struct {
	Code string `form:"code" json:"code" binding:"required"`
}
