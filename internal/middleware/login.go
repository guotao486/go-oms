/*
 * @Author: GG
 * @Date: 2023-03-28 16:20:12
 * @LastEditTime: 2023-04-06 15:12:25
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\middleware\login.go
 *
 */
package middleware

import (
	"encoding/gob"
	"net/http"
	"oms/global"
	"oms/internal/model"
	"oms/pkg/errcode"
	"oms/pkg/sessions"

	"github.com/gin-gonic/gin"
)

func IsLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessions := sessions.NewSession(c)
		user := sessions.Get("userinfo")
		if user != nil {
			gob.Register(model.User{})
			sessions.Set("userinfo", user)
			c.Next()
		}
		global.Logger.Errorf(c, "middleware login err: ", errcode.NotLogin)
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
	}
}
