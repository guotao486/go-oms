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
		"menus":        menus.GetMenusTree(ctx),
		"currentMenus": menus.GetCurrent(ctx),
		"urlPath":      ctx.Request.URL.Path,
		"appConfig":    nil,
		"appTitle":     global.AppSetting.AppName,
	}
	ctx.HTML(code, name, data)
}
