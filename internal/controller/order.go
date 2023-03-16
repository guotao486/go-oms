package controller

import (
	"fmt"
	"net/http"
	"oms/global"
	"oms/internal/request"
	"oms/internal/service"
	"oms/pkg/app"
	"strings"

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
	} else {
		// p := c.PostFormArray("product")
		// fmt.Printf("p: %v\n", p)
		// p2 := c.PostFormMap("product")
		// fmt.Printf("p2: %v\n", p2)
		dicts := make(map[string]*request.CreateOrderProductRequest)
		key := ""
		key2 := ""
		c.Request.ParseForm()
		for k, v := range c.Request.PostForm {
			if i := strings.IndexByte(k, '['); i >= 1 && k[0:i] == "product" {
				if j := strings.IndexByte(k[i+1:], ']'); j >= 1 {
					key = k[i+1:][:j]
					if dicts[key] == nil {
						dicts[key] = &request.CreateOrderProductRequest{}
					}

					if i2 := strings.IndexByte(k[i+1:][j:], '['); i2 >= 1 {
						if j2 := strings.IndexByte(k[i+1:][j:][i2+1:], ']'); j2 >= 1 {
							key2 = k[i+1:][j:][i2+1:][:j2]
							switch key2 {
							case "name":
								dicts[key].Name = v[0]
							case "sku":
								dicts[key].Sku = v[0]
							case "image":
								dicts[key].Images = v[0]
							case "attribute":
								dicts[key].Attribute = v[0]
							}
						}
					}
				}
			}
		}
		fmt.Printf("dicts: %v\n", dicts)
		param := request.CreateOrderRequest{OrderProducts: dicts}
		valid, errs := app.BindAndValid(c, &param)
		fmt.Printf("param: %v\n", param)
		fmt.Printf("param.OrderProducts: %v\n", param.OrderProducts)
		for _, v := range param.OrderProducts {
			fmt.Printf("v.Name: %v\n", v.Name)
			fmt.Printf("v.Sku: %v\n", v.Sku)
			fmt.Printf("v.Images: %v\n", v.Images)
			fmt.Printf("v.Attribute: %v\n", v.Attribute)
		}
		if !valid {
			global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
			return
		}
	}
}
func (o *OrderController) Update(c *gin.Context) {

}
func (o *OrderController) Delete(c *gin.Context) {

}
