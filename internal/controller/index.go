package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct{}

func NewIndex() *IndexController {
	return &IndexController{}
}

// 首页
func (i *IndexController) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index", nil)
}
