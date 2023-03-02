package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
}

func NewAuth() AuthController {
	return AuthController{}
}

// Login 登录页面
func (a *AuthController) Login(c *gin.Context) {
	if c.Request.Method == http.MethodPost {

		c.Redirect(301, "/")
	}
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "login", nil)
	}
}
