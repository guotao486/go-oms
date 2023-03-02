package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

// ContextTimeout
/**
 * @description: 超时时间控制中间件
 * @param {time.Duration} t
 * @return {*}
 */
func ContextTimeout(t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置当前context 超时时间
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
