/*
 * @Author: GG
 * @Date: 2023-02-08 17:03:08
 * @LastEditTime: 2023-02-08 17:04:14
 * @LastEditors: GG
 * @Description: 限流器中间件
 * @FilePath: \oms\internal\middleware\limiter.go
 *
 */
package middleware

import (
	"oms/pkg/app"
	"oms/pkg/errcode"
	"oms/pkg/limiter"

	"github.com/gin-gonic/gin"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			// 如果有，它返回移除的令牌数；没有可用的令牌，则返回零
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
