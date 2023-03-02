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

type UserController struct{}

func NewUser() *UserController {
	return &UserController{}
}

// 用户列表
func (u *UserController) List(c *gin.Context) {

}

// 创建用户
func (u *UserController) Create(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "user/create", nil)
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
		err := svc.CreateUser(&param)
		if err != nil {
			global.Logger.Errorf(c, "svc.CreateTag err: %v", err)
			response.ToErrorResponse(errcode.ErrorCreateUserFail)
			return
		}
		response.ToResponse(nil)
		return
	}
}

// 修改用户
func (u *UserController) Update(c *gin.Context) {

}

// 删除用户，默认软删除
func (u *UserController) Delete(c *gin.Context) {

}
