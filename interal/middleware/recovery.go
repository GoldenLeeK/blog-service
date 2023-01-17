package middleware

import (
	"fmt"
	"github.com/GoldenLeeK/blog-service/global"
	"github.com/GoldenLeeK/blog-service/pkg/app"
	"github.com/GoldenLeeK/blog-service/pkg/email"
	"github.com/GoldenLeeK/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"time"
)

func Recovery() gin.HandlerFunc {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf("panic recover :%v", err)

				err := defailtMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出,发生时间:%d", time.Now().Unix()),
					fmt.Sprintf("错误信息:%v", err))

				if err != nil {
					global.Logger.Panicf("mail.SendMail err :%v", err)
				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
