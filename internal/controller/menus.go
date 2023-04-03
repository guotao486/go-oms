/*
 * @Author: GG
 * @Date: 2023-03-31 09:33:39
 * @LastEditTime: 2023-04-03 16:03:39
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\controller\menus.go
 *
 */
package controller

import (
	"net/http"
	"oms/global"
	"oms/internal/request"
	"oms/internal/service"
	"oms/pkg/app"
	"oms/pkg/convert"
	"oms/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type MenusController struct {
	*Controller
}

func NewMenus() *MenusController {
	return &MenusController{}
}

func (m *MenusController) Index(c *gin.Context) {
	m.RenderHtml(c, http.StatusOK, "menus/list", nil)
}

func (m *MenusController) List(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	list, err := svc.GetMenusListAll()
	if err != nil {
		global.Logger.Errorf(c, "svc.GetMenusListAll err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetMenusListFail)
		return
	}
	s := app.NewSuccess()
	s.Code = 0
	s.Data = list
	response.ToResponse(s)
	return
}

func (m *MenusController) Create(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		response := app.NewResponse(c)
		svc := service.New(c.Request.Context())
		list, err := svc.GetParentMenusList()
		if err != nil {
			global.Logger.Errorf(c, "svc.GetParentMenusList err: %v", err)
			response.ToErrorBadRequestHtml(errcode.ErrorGetMenusListFail)
			return
		}
		m.RenderHtml(c, http.StatusOK, "menus/create", gin.H{
			"parentMenus": list,
		})
	} else {
		param := request.CreateMenusRequest{}
		response := app.NewResponse(c)

		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			response.ToErrorResponse(errRsp)
			return
		}

		svc := service.New(c.Request.Context())
		err := svc.CreateMenus(&param)
		if err != nil {
			global.Logger.Errorf(c, "svc.CreateMenus err: %v", err)
			switch e := err.(type) {
			case *errcode.Error:
				response.ToErrorResponse(e)
			default:
				response.ToErrorResponse(errcode.ErrorCreateMenusFail)
			}
			return
		}
		response.ToSuccessResponse(nil)
		return
	}
}

func (m *MenusController) Update(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		param := request.UpdateMenusGetRequest{}
		response := app.NewResponse(c)

		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Error(c, "app.BindAndValid errs: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			response.ToErrorBadRequestHtml(errRsp)
			return
		}

		svc := service.New(c.Request.Context())
		menus, err := svc.GetMenusById(param.ID)
		if err != nil {
			global.Logger.Errorf(c, "svc.GetMenusById err: %v", err)
			response.ToErrorBadRequestHtml(errcode.ErrorMenusNotFoundFail)
			return
		}
		list, err := svc.GetParentMenusList()
		if err != nil {
			global.Logger.Errorf(c, "svc.GetParentMenusList err: %v", err)
			response.ToErrorBadRequestHtml(errcode.ErrorGetMenusListFail)
			return
		}
		m.RenderHtml(c, http.StatusOK, "menus/update", gin.H{
			"parentMenus": list,
			"menusInfo":   menus,
		})
		return
	} else {
		param := request.UpdateMenusPostRequest{}
		response := app.NewResponse(c)

		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Error(c, "app.BindAndValid errs: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			response.ToErrorResponse(errRsp)
			return
		}

		svc := service.New(c.Request.Context())
		err := svc.UpdateMenus(&param)
		if err != nil {
			switch e := err.(type) {
			case *errcode.Error:
				response.ToErrorResponse(e)
			default:
				response.ToErrorResponse(errcode.ErrorUpdateMenusFail)
			}
			return
		}

		response.ToSuccessResponse(nil)
		return
	}
}

func (m *MenusController) Delete(c *gin.Context) {
	if c.Request.Method == http.MethodDelete {
		param := request.DeleteMenusRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
		response := app.NewResponse(c)

		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Error(c, "app.BindAndValid errs: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			response.ToErrorResponse(errRsp)
			return
		}

		svc := service.New(c.Request.Context())
		err := svc.DeleteMenus(&param)
		if err != nil {
			switch e := err.(type) {
			case *errcode.Error:
				response.ToErrorResponse(e)
			default:
				response.ToErrorResponse(errcode.ErrorUpdateMenusFail)
			}
			return
		}

		response.ToSuccessResponse(nil)
		return

	}
}

func (m *MenusController) UpdateSort(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		param := request.UpdateMenusSortRequest{}
		response := app.NewResponse(c)

		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Error(c, "app.BindAndValid errs: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			response.ToErrorResponse(errRsp)
			return
		}

		svc := service.New(c.Request.Context())
		err := svc.UpdateMenusSort(&param)
		if err != nil {
			switch e := err.(type) {
			case *errcode.Error:
				response.ToErrorResponse(e)
			default:
				response.ToErrorResponse(errcode.ErrorUpdateMenusFail)
			}
			return
		}

		response.ToSuccessResponse(nil)
		return
	}
}
