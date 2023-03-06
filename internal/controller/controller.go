package controller

import (
	"oms/global"
	"oms/pkg/menus"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func (c *Controller) RenderHtml(ctx *gin.Context, code int, name string, obj interface{}) {
	data := gin.H{
		"data":         obj,
		"menus":        menus.GetMenus(),
		"currentMenus": menus.GetCurrent(ctx.Request.URL.Path),
		"urlPath":      ctx.Request.URL.Path,
		"appConfig":    nil,
		"appTitle":     global.AppSetting.AppName,
	}
	ctx.HTML(code, name, data)
}
