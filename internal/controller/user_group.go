package controller

import (
	"net/http"
	"oms/internal/service"

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
			return
		}
		u.RenderHtml(c, http.StatusOK, "group/create", gin.H{
			"userList": userList,
		})
	} else {

	}
}
