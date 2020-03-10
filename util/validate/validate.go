package validate

import (
	"github.com/gin-gonic/gin"
	zhongwen "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"go-gin-api/util"
	"net/http"
)
type BaseField struct {
	Field interface{}
}

func (Base *BaseField) Validator(c *gin.Context) (b bool) {
	b = false
	res := util.Gin{Ctx: c}
	validate := validator.New()
	zh := zhongwen.New()
	uni := ut.New(zh, zh)
	trans, _ := uni.GetTranslator("zh")
	zh_translations.RegisterDefaultTranslations(validate, trans)
	c.ShouldBind(&Base.Field)
	if err := validate.Struct(Base.Field); err != nil {
		res.Response(http.StatusUnprocessableEntity, err.(validator.ValidationErrors).Translate(trans))
		b = true
	}
	return
}
