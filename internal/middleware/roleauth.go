/*
 * @Author: GG
 * @Date: 2023-04-06 15:56:04
 * @LastEditTime: 2023-04-06 16:55:06
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\middleware\roleauth.go
 *
 */
package middleware

import (
	"net/http"
	"oms/global"
	"oms/internal/controller"
	"oms/internal/model"
	"oms/pkg/app"
	"oms/pkg/errcode"
	"oms/pkg/menus"
	"oms/pkg/sessions"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RoleAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessions := sessions.NewSession(c)
		user := sessions.Get("userinfo")
		userinfo, ok := user.(model.User)
		if ok {
			currentMenu := menus.GetCurrent(c)
			if app.InString(currentMenu.Role, strconv.Itoa(int(userinfo.Level))) {
				c.Next()
			} else {
				global.Logger.Errorf(c, "middleware RoleAuth err: ", errcode.NotRoleAuth)
				controller := &controller.Controller{}
				controller.ToErrorUnauthorizedHtml(c, errcode.NotRoleAuth)
				c.Abort()
			}
		}
		global.Logger.Errorf(c, "middleware RoleAuth err: ", errcode.NotLogin)
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
	}
}
