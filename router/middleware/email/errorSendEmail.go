package email

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-api/config"
	"go-gin-api/util"
	"go-gin-api/util/errorNotify"
	"runtime/debug"
	"strings"
	"time"
)

func ErrorSentMail() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "<br>"
				}
				subject := fmt.Sprintf("【重要错误】%s 项目出错了！", config.GinConfig.App.Name)
				body := strings.ReplaceAll(MailTemplate, "{ErrorMsg}", fmt.Sprintf("%s", err))
				body = strings.ReplaceAll(body, "{RequestTime}", time.Now().Format("2006/01/02 - 15:04:05"))
				body = strings.ReplaceAll(body, "{RequestURL}", c.Request.Method+"  "+c.Request.Host+c.Request.RequestURI)
				body = strings.ReplaceAll(body, "{RequestUA}", c.Request.UserAgent())
				body = strings.ReplaceAll(body, "{RequestIP}", c.ClientIP())
				body = strings.ReplaceAll(body, "{DebugStack}", DebugStack)
				//记录到邮件
				_ = errorNotify.SendMail("jiangzhibin1994@gmail.com", subject, body)
				//记录到日志
				 errorNotify.Logger(DebugStack)
				utilGin := util.Gin{Ctx: c}
				utilGin.Response(500, "系统异常，请联系管理员！", nil)
			}
		}()
		c.Next()
	}
}

func ErrorToFile() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "<br>"
				}
				errorNotify.Logger(DebugStack)
			}
		}()
	}
}
