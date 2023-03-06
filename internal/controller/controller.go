package controller

import (
	"net/http"
	"oms/global"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func (c *Controller) RenderHtml(ctx *gin.Context, code int, name string, obj interface{}) {
	data := gin.H{
		"data":      obj,
		"appConfig": nil,
		"appTitle":  global.AppSetting.AppName,
	}
	ctx.HTML(http.StatusOK, "user/list", data)
}
