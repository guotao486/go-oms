package controller

import (
	"encoding/gob"
	"net/http"
	"oms/global"
	"oms/internal/model"
	"oms/internal/request"
	"oms/internal/service"
	"oms/pkg/app"
	"oms/pkg/errcode"
	"oms/pkg/sessions"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
}

func NewAuth() AuthController {
	return AuthController{}
}

// Login 登录页面
func (a *AuthController) Login(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "login", nil)
	} else {
		param := request.LoginRequest{}
		response := app.NewResponse(c)
		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
			response.ToErrorResponse(errcode.NewError(400, errs.Error()))
			return
		}
		svc := service.New(c.Request.Context())
		user, err := svc.Login(&param)
		if err != nil {
			global.Logger.Errorf(c, "svc.Login err: %v", err)
			switch e := err.(type) {
			case *errcode.Error:
				response.ToErrorResponse(e)
			default:
				response.ToErrorResponse(errcode.ErrorLoginFail)
			}
			return
		}

		gob.Register(model.User{})
		sessions := sessions.NewSession(c)
		sessions.Set("userinfo", user)
		sessions.Save()
		response.ToSuccessResponse(nil)
		return
	}
}

func (a *AuthController) Logout(c *gin.Context) {
	session := sessions.NewSession(c)
	session.Delete("userinfo")
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/login")
}
