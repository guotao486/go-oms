package middleware

import (
	"fmt"
	"net/http"
	"oms/global"
	"oms/pkg/errcode"
	"oms/pkg/sessions"

	"github.com/gin-gonic/gin"
)

func IsLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessions := sessions.NewSession(c)
		user := sessions.Get("userinfo")
		fmt.Printf("user: %v\n", user)
		if user != nil {
			c.Next()
		}
		global.Logger.Errorf(c, "middleware login err: ", errcode.NotLogin)
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
	}
}
