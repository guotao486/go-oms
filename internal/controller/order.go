package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	*Controller
}

func NewOrder() *OrderController {
	return &OrderController{}
}

func (o *OrderController) Index(c *gin.Context) {
	o.RenderHtml(c, http.StatusOK, "order/list", nil)
}

func (o *OrderController) List(c *gin.Context) {

}

func (o *OrderController) Create(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		o.RenderHtml(c, http.StatusOK, "order/create", nil)
	}
}
func (o *OrderController) Update(c *gin.Context) {

}
func (o *OrderController) Delete(c *gin.Context) {

}
