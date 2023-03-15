package controller

import (
	"net/http"
	"oms/global"
	"oms/internal/service"

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
		svc := service.New(c)
		currencyList, err := svc.GetCurrencyList()
		if err != nil {
			global.Logger.Errorf(c, "svc.GetCurrencyList err: %v", err)
		}
		orderShippingList, err := svc.GetOrderShippingList()
		if err != nil {
			global.Logger.Errorf(c, "svc.GetCurrencyList err: %v", err)
		}
		orderStatusList, err := svc.GetOrderStatusList()
		if err != nil {
			global.Logger.Errorf(c, "svc.GetCurrencyList err: %v", err)
		}
		paymentStatusList, err := svc.GetPaymentStatusList()
		if err != nil {
			global.Logger.Errorf(c, "svc.GetCurrencyList err: %v", err)
		}
		paymentTypeList, err := svc.GetPaymentTypeList()
		if err != nil {
			global.Logger.Errorf(c, "svc.GetCurrencyList err: %v", err)
		}
		o.RenderHtml(c, http.StatusOK, "order/create", gin.H{
			"currency":      currencyList,
			"orderShipping": orderShippingList,
			"orderStatus":   orderStatusList,
			"paymentStatus": paymentStatusList,
			"paymentType":   paymentTypeList,
		})
	}
}
func (o *OrderController) Update(c *gin.Context) {

}
func (o *OrderController) Delete(c *gin.Context) {

}
