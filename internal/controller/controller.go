package controller

import (
	"net/http"
	"oms/global"
	"oms/pkg/errcode"
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

// 400
func (c *Controller) ToErrorBadRequestHtml(ctx *gin.Context, err *errcode.Error) {
	c.RenderHtml(ctx, http.StatusBadRequest, "error/400", err.Msg())
}

// 403
func (c *Controller) ToErrorForbiddenHtml(ctx *gin.Context, err *errcode.Error) {
	c.RenderHtml(ctx, http.StatusForbidden, "error/403", err.Msg())
}

// 404
func (c *Controller) ToErrorNotFoundHtml(ctx *gin.Context, err *errcode.Error) {
	c.RenderHtml(ctx, http.StatusNotFound, "error/404", err.Msg())
}

// 401
func (c *Controller) ToErrorUnauthorizedHtml(ctx *gin.Context, err *errcode.Error) {
	c.RenderHtml(ctx, http.StatusUnauthorized, "error/401", err.Msg())
}
