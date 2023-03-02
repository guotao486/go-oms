/*
 * @Author: GG
 * @Date: 2023-02-06 16:04:18
 * @LastEditTime: 2023-02-19 12:12:51
 * @LastEditors: GG
 * @Description: auth controller
 * @FilePath: \oms\internal\routers\api\auth.go
 *
 */
package api

import (
	"oms/global"
	"oms/internal/request"
	"oms/internal/service"
	"oms/pkg/app"
	"oms/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// @Tags Auth
// @Produce json
// @Summary 获取auth token
// @Param AuthRequest body request.AuthRequest true "AuthRequest"
// @Success 200 {object} app.Success "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /auth [post]
func GetAuth(c *gin.Context) {
	param := request.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToSuccessResponse(gin.H{
		"token": token,
	})
	return
}
