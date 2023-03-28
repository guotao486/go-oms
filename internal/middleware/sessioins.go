package middleware

import (
	"oms/global"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Sessions() gin.HandlerFunc {
	store := cookie.NewStore([]byte(global.AppSetting.AppName))
	store.Options(sessions.Options{MaxAge: 86400 * 30})
	return sessions.Sessions("mysessions", store)
}
