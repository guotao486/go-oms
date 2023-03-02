/*
 * @Author: GG
 * @Date: 2023-02-08 15:43:41
 * @LastEditTime: 2023-02-08 16:00:57
 * @LastEditors: GG
 * @Description: 服务信息存储中间件
 * @FilePath: \oms\internal\middleware\app_info.go
 *
 */
package middleware

import "github.com/gin-gonic/gin"

// AppInfo
/**
 * @description: 服务信息存储中间件
 * @return {*}
 */
func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "oms")
		c.Set("app_version", "1.0.0")
		c.Next()
	}
}
