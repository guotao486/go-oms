package controller

import (
	"net/http"
	"oms/global"
	"oms/internal/request"
	"oms/internal/service"
	"oms/pkg/app"
	"oms/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type UserGroupController struct {
	*Controller
}

func NewUserGroup() *UserGroupController {
	return &UserGroupController{}
}

func (u *UserGroupController) Index(c *gin.Context) {
	u.RenderHtml(c, http.StatusOK, "group/list", nil)
}

func (u *UserGroupController) List(c *gin.Context) {}

func (u *UserGroupController) Create(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		svc := service.New(c)
		response := app.NewResponse(c)
		userList, err := svc.GetUserListAll()
		if err != nil {
			global.Logger.Errorf(c, "svc.GetUserListAll errs: %v", err)
			response.ToErrorBadRequestHtml(errcode.ErrorGetUserListFail)
			return
		}
		u.RenderHtml(c, http.StatusOK, "group/create", gin.H{
			"userList": userList,
		})
	} else {
		param := request.CreateUserGroupRequest{}
		response := app.NewResponse(c)
		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			response.ToErrorResponse(errRsp)
			return
		}
		svc := service.New(c)
		err := svc.CreateUserGroup(&param)
		if err != nil {
			global.Logger.Errorf(c, "svc.CreateUserGroup err: %v", err)
			switch e := err.(type) {
			case *errcode.Error:
				response.ToErrorResponse(e)
			default:
				response.ToErrorResponse(errcode.ErrorCreateUserGroupFail)
			}
			return
		}

		response.ToSuccessResponse(nil)
		return
	}
}

func (u *UserGroupController) Update(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		svc := service.New(c)
		response := app.NewResponse(c)
		param := request.UpdateUserGroupGetRequest{}
		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			response.ToErrorBadRequestHtml(errRsp)
			return
		}
		userGroup, err := svc.GetUserGroupById(param.ID)
		if err != nil {
			global.Logger.Errorf(c, "svc.GetUserGroupById err: %v", err)
			response.ToErrorNotFoundHtml(errcode.ErrorUserGroupNotFoundFail)
			return
		}
		userList, err := svc.GetUserListAll()
		if err != nil {
			global.Logger.Errorf(c, "svc.GetUserListAll errs: %v", err)
			response.ToErrorBadRequestHtml(errcode.ErrorGetUserListFail)
			return
		}
		u.RenderHtml(c, http.StatusOK, "group/update", gin.H{
			"userList": userList,
			"detail":   userGroup,
		})
		return
	} else {
		param := request.UpdateUserGroupPostRequest{}
		response := app.NewResponse(c)
		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			response.ToErrorResponse(errRsp)
			return
		}

		svc := service.New(c)
		err := svc.UpdateUserGroup(&param)
		if err != nil {
			switch e := err.(type) {
			case *errcode.Error:
				response.ToErrorResponse(e)
			default:
				response.ToErrorResponse(errcode.ErrorUpdateUserGroupFail)
			}
			return
		}
		response.ToResponse(nil)
		return
	}
}
