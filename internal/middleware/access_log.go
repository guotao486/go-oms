/*
 * @Author: GG
 * @Date: 2023-02-07 16:16:49
 * @LastEditTime: 2023-02-19 12:10:57
 * @LastEditors: GG
 * @Description: 访问日志中间件
 * @FilePath: \oms\internal\middleware\access_log.go
 *
 */
package middleware

import (
	"bytes"
	"oms/global"
	"oms/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

// 响应日志结构体
//
type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

// AccessLog
/**
 * @description: 访问日志中间件
 * @return {*}
 */
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWrite := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWrite

		beginTime := time.Now()
		c.Next()
		endTime := time.Now()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWrite.body.String(),
		}

		global.Logger.WithFields(fields).Infof(c, "access log: method: %s, status_code: %d, begin_time: %s, end_time: %s", c.Request.Method, bodyWrite.Status(), beginTime, endTime)

	}
}
