/*
 * @Author: GG
 * @Date: 2023-01-28 11:04:28
 * @LastEditTime: 2023-02-28 15:21:24
 * @LastEditors: GG
 * @Description: tag controller
 * @FilePath: \oms\internal\routers\api\v1\tag.go
 *
 */
package v1

import (
	"fmt"
	"oms/global"
	"oms/internal/request"
	"oms/internal/service"
	"oms/pkg/app"

	errcode "oms/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

// @Summary 获取单个标签
// @Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Success 200 {object} demo.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tag [get]
func (t Tag) Get(c *gin.Context) {
	param := struct {
		Name  string `form:"name" binding:"bookabledate,max=100"`
		State uint8  `form:"state,default=1" binding:"oneof=0 1"`
	}{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	// 校验失败
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		// 入参错误对象，将参数校验错误信息，存入对象中
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		// 将错误对象，传给错误响应对象
		response.ToErrorResponse(errRsp)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 获取多个标签
// @Produce  json
// @param token header string false "token"
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} demo.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	// 初始化请求对象
	param := request.TagListRequest{}
	// 初始化响应对象
	response := app.NewResponse(c)
	// 参数校验
	valid, errs := app.BindAndValid(c, &param)
	// 校验失败
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		// 入参错误对象，将参数校验错误信息，存入对象中
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		// 将错误对象，传给错误响应对象
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTag(&request.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		// 统计错误
		global.Logger.Errorf(c, "svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	tags, err := svc.GetListTag(&param, &pager)
	if err != nil {
		// 分页查询错误
		global.Logger.Errorf(c, "svc.GetListTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags, totalRows)
	return
}

// @Summary 新增标签
// @Produce  json
// @Param name formData string true "标签名称" minlength(3) maxlength(100)
// @Param state formData int false "状态" Enums(0, 1) default(1)
// @Param created_by formData string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} app.Success "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := request.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	// 校验失败
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		// 入参错误对象，将参数校验错误信息，存入对象中
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		// 将错误对象，传给错误响应对象
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(nil)
	return
}

// @Summary 更新标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Param UpdateTagRequest body request.UpdateTagRequest true "UpdateTagRequest"
// @Success 200 {array} app.Success "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := request.UpdateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	fmt.Printf("param: %v\n", param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	fmt.Printf("param: %v\n", param)
	response.ToResponse(nil)
	return
}

// @Summary 删除标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Success 200 {string} app.Success "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	param := request.DeleteTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	fmt.Printf("param: %v\n", param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}

	response.ToResponse(nil)
	return
}
