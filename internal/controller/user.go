/*
 * @Author: GG
 * @Date: 2023-03-01 09:30:13
 * @LastEditTime: 2023-03-06 09:21:40
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\controller\user.go
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

type UserController struct {
	*Controller
}

func NewUser() *UserController {
	return &UserController{}
}

// 用户列表
func (u *UserController) Index(c *gin.Context) {
	u.RenderHtml(c, http.StatusOK, "user/list", nil)
}

// 用户列表数据
func (u *UserController) List(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		param := request.GetListUserRequest{}
		response := app.NewResponse(c)
		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			response.ToErrorResponse(errRsp)
			return
		}

		svc := service.New(c.Request.Context())
		pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
		totalRows, err := svc.GetUserCountList(&param)
		if err != nil {
			// 统计错误
			global.Logger.Errorf(c, "svc.GetUserCountList err: %v", err)
			response.ToErrorResponse(errcode.ErrorCountUserFail)
			return
		}

		users, err := svc.GetUserListPager(&param, &pager)
		if err != nil {
			// 分页查询错误
			global.Logger.Errorf(c, "svc.GetUserListPager err: %v", err)
			response.ToErrorResponse(errcode.ErrorGetUserListFail)
			return
		}

		response.ToResponseList(users, totalRows)
		return
	}
}

// 创建用户
func (u *UserController) Create(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		u.RenderHtml(c, http.StatusOK, "user/create", nil)
	} else {
		param := request.CreateUserRequest{}
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

		// 检查用户名是否存在
		user, err := svc.GetUserInfoByUsername(param.Username)
		if user != nil {
			global.Logger.Errorf(c, "svc.GetUserInfoByUsername err: %v", errcode.ErrorUserExists.Msg())
			response.ToErrorResponse(errcode.ErrorUserExists)
			return
		}

		// 创建用户
		err = svc.CreateUser(&param)
		if err != nil {
			global.Logger.Errorf(c, "svc.CreateUser err: %v", err)
			response.ToErrorResponse(errcode.ErrorCreateUserFail)
			return
		}
		response.ToResponse(nil)
		return
	}
}

// 修改用户
func (u *UserController) Update(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		param := request.UpdateUserGetRequest{}
		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			// 参数错误
			u.ToErrorBadRequestHtml(c, errRsp)
			return
		}
		svc := service.New(c.Request.Context())
		user, _ := svc.GetUserInfoById(param.ID)
		if user == nil {
			// 没有数据
			u.ToErrorNotFoundHtml(c, errcode.ErrorUserNotFound)
			return
		}
		u.RenderHtml(c, http.StatusOK, "user/update", user)
		return
	} else {
		param := request.UpdateUserPostRequest{}
		response := app.NewResponse(c)
		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			response.ToErrorResponse(errRsp)
			return
		}
		svc := service.New(c.Request.Context())
		err := svc.UpdateUser(&param)
		if err != nil {
			global.Logger.Errorf(c, "svc.UpdateUser err: %v", err)

			switch e := err.(type) {
			case *errcode.Error:
				// 自定义的错误，用户是否存在，新用户名是否存在
				response.ToErrorResponse(e)
			default:
				response.ToErrorResponse(errcode.ErrorUpdateUserFail)
			}
			return
		}
		response.ToSuccessResponse(nil)
		return
	}
}

// 删除用户，默认软删除
func (u *UserController) Delete(c *gin.Context) {
	if c.Request.Method == http.MethodDelete {
		param := request.DeleteUserRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
		response := app.NewResponse(c)
		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			response.ToErrorResponse(errRsp)
			return
		}

		svc := service.New(c.Request.Context())
		err := svc.DeleteUser(&param)
		if err != nil {
			global.Logger.Errorf(c, "svc.DeleteUser err: %v", err)
			switch e := err.(type) {
			case *errcode.Error:
				response.ToErrorResponse(e)
			default:
				response.ToErrorResponse(errcode.ErrorDeleteUserFail)
			}
			return
		}
		response.ToSuccessResponse(nil)
		return
	}
}
