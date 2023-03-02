/*
 * @Author: GG
 * @Date: 2023-02-08 10:43:39
 * @LastEditTime: 2023-02-08 15:35:33
 * @LastEditors: GG
 * @Description: 异常捕获中间件
 * @FilePath: \oms\internal\middleware\recovery.go
 *
 */
package middleware

import (
	"fmt"
	"oms/global"
	"oms/pkg/app"
	"oms/pkg/email"
	"oms/pkg/errcode"
	"time"

	"github.com/gin-gonic/gin"
)

// Recovery
/**
 * @description: 异常捕获中间件
 * @return {*}
 */
func Recovery() gin.HandlerFunc {
	// 默认邮件程序
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		UserName: global.EmailSetting.Username,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
		IsSSL:    global.EmailSetting.IsSSL,
	})
	return func(c *gin.Context) {
		// 捕获异常并记录日志
		defer func() {
			if err := recover(); err != nil {
				// 记录异常错误+内存栈信息
				global.Logger.WithCallersFrames().Panicf(c, "panic recover err: %v", err)

				// 发送邮件
				err := defailtMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间：%s", time.Now()),
					fmt.Sprintf("错误信息： %v", err),
				)
				if err != nil {
					global.Logger.Panicf(c, "mail.SendMail err: %v", err)
				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
