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

func (u *UserGroupController) Index(c *gin.Context) {}

func (u *UserGroupController) List(c *gin.Context) {}

func (u *UserGroupController) Create(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		svc := service.New(c)
		userList, err := svc.GetUserCountListAll()
		if err != nil {
			global.Logger.Errorf(c, "svc.GetUserCountListAll err: %v", err)
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
				response.ToErrorResponse(errcode.ErrorUpdateUserFail)
			}
			return
		}

		response.ToSuccessResponse(nil)
		return
	}
}
