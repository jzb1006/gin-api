package errorNotify

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-gin-api/config"
	log2 "go-gin-api/init/log"
	"gopkg.in/gomail.v2"
	"strings"
	"time"
)

func SendMail(mailTo string, subject string, body string) error {
	if config.GinConfig.ErrorMailSend.Active != true {
		return nil
	}
	m := gomail.NewMessage()
	m.SetHeader("From", config.GinConfig.ErrorMailSend.User)
	mailArrTo := strings.Split(mailTo, ",")
	m.SetHeader("To", mailArrTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(config.GinConfig.ErrorMailSend.Host, config.GinConfig.ErrorMailSend.Port, config.GinConfig.ErrorMailSend.User, config.GinConfig.ErrorMailSend.Pass)
	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func Logger(str string) {
	log2.InitWarnLog()
	errorLogMap := make(map[string]interface{})
	errorLogMap["info"] = str
	errorLogMap["time"] = time.Now().Format("2006/01/02 - 15:04:05")
	//errorLogJson, _ := json.Marshal(errorLogMap)
	//实例化
	logger := log2.Logger

	// 日志格式
	logger.WithFields(
		logrus.Fields{
			"time": errorLogMap["time"],
			"info": errorLogMap["info"],
		}).Warn()

}
