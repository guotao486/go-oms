package controller

import (
	"net/http"
	"oms/global"
	"oms/internal/request"
	"oms/internal/service"
	"oms/pkg/app"
	"oms/pkg/convert"
	"oms/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	*Controller
}

func NewOrder() *OrderController {
	return &OrderController{}
}

func (o *OrderController) Index(c *gin.Context) {
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
	o.RenderHtml(c, http.StatusOK, "order/list", gin.H{
		"currency":      currencyList,
		"orderShipping": orderShippingList,
		"orderStatus":   orderStatusList,
		"paymentStatus": paymentStatusList,
		"paymentType":   paymentTypeList,
	})
}

func (o *OrderController) List(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		param := request.GetOrderListRequest{}
		response := app.NewResponse(c)
		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			response.ToErrorResponse(errRsp)
			return
		}

		svc := service.New(c.Request.Context())
		pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
		totalRows, err := svc.GetOrderCountList(&param)
		if err != nil {
			// 统计错误
			global.Logger.Errorf(c, "svc.GetOrderCountList err: %v", err)
			response.ToErrorResponse(errcode.ErrorCountOrderFail)
			return
		}

		orders, err := svc.GetOrderListPager(&param, &pager)
		if err != nil {
			// 分页查询错误
			global.Logger.Errorf(c, "svc.GetOrderListPager err: %v", err)
			response.ToErrorResponse(errcode.ErrorGetOrderListFail)
			return
		}

		response.ToResponseList(orders, totalRows)
		return
	}
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

		dicts := make(map[string]*request.CreateOrderProductRequest)
		c.Request.ParseForm()

		p := app.GetPostMapForm(c.Request.PostForm, "product")
		for k, v := range p {
			product := &request.CreateOrderProductRequest{}
			app.GetPostMapFormItem(c.Request.PostForm, v["i"], v["j"], "product", k, product)
			dicts[k] = product
		}

		param := request.CreateOrderRequest{OrderProducts: dicts}
		response := app.NewResponse(c)
		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			response.ToErrorResponse(errRsp)
			return
		}
		svc := service.New(c)
		err := svc.CreateOrder(&param)
		if err != nil {
			global.Logger.Errorf(c, "svc.CreateOrder err :%v", err)
			switch e := err.(type) {
			case *errcode.Error:
				response.ToErrorResponse(e)
			default:
				response.ToErrorResponse(errcode.ErrorCreateOrderFail)
			}
			return
		}
		response.ToSuccessResponse(nil)
		return
	}
}
func (o *OrderController) Update(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		param := request.UpdateOrderGetRequest{}
		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			o.ToErrorBadRequestHtml(c, errRsp)
			return
		}
		svc := service.New(c.Request.Context())
		order, err := svc.GetOrderById(param.ID)
		if err != nil {
			global.Logger.Errorf(c, "svc.GetOrderById err: %v", errs)
			o.ToErrorBadRequestHtml(c, errcode.ErrorOrderNotFoundFail)
			return
		}

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
		o.RenderHtml(c, http.StatusOK, "order/update", gin.H{
			"orderInfo":     order,
			"currency":      currencyList,
			"orderShipping": orderShippingList,
			"orderStatus":   orderStatusList,
			"paymentStatus": paymentStatusList,
			"paymentType":   paymentTypeList,
		})
	} else {

		dicts := make(map[string]*request.CreateOrderProductRequest)
		c.Request.ParseForm()
		p := app.GetPostMapForm(c.Request.PostForm, "product")
		for k, v := range p {
			product := &request.CreateOrderProductRequest{}
			app.GetPostMapFormItem(c.Request.PostForm, v["i"], v["j"], "product", k, product)
			dicts[k] = product
		}
		param := request.UpdateOrderPostRequest{OrderProducts: dicts}
		response := app.NewResponse(c)
		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			response.ToErrorResponse(errRsp)
			return
		}

		svc := service.New(c)
		err := svc.UpdateOrder(&param)
		if err != nil {
			global.Logger.Errorf(c, "svc.UpdateOrder err :%v", err)
			switch e := err.(type) {
			case *errcode.Error:
				response.ToErrorResponse(e)
			default:
				response.ToErrorResponse(errcode.ErrorUpdateOrderFail)
			}
			return
		}
		response.ToSuccessResponse(nil)
		return
	}
}
func (o *OrderController) Delete(c *gin.Context) {
	if c.Request.Method == http.MethodDelete {
		param := request.DeleteOrderRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
		response := app.NewResponse(c)
		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			response.ToErrorResponse(errRsp)
			return
		}

		svc := service.New(c.Request.Context())
		err := svc.DeleteOrder(&param)
		if err != nil {
			global.Logger.Errorf(c, "svc.DeleteUser err: %v", err)
			switch e := err.(type) {
			case *errcode.Error:
				response.ToErrorResponse(e)
			default:
				response.ToErrorResponse(errcode.ErrorDeleteUserFail)
			}
			return
		}
		response.ToSuccessResponse(nil)
		return
	}
}

func (o *OrderController) AjaxUpdatePayment(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		param := request.AjaxUpdateOrderGetRequest{}
		response := app.NewResponse(c)
		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			response.ToErrorResponse(errRsp)
			return
		}
		svc := service.New(c.Request.Context())
		order, err := svc.GetOrderById(param.ID)
		if err != nil {
			global.Logger.Errorf(c, "svc.GetOrderById err: %v", errs)
			response.ToErrorResponse(errcode.ErrorOrderNotFoundFail)
			return
		}

		var data gin.H
		currencyList, err := svc.GetCurrencyList()
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
		data = gin.H{
			"orderInfo":     order,
			"currency":      currencyList,
			"paymentStatus": paymentStatusList,
			"paymentType":   paymentTypeList,
		}

		c.HTML(http.StatusOK, "ajaxForm/order_ajax_update_payment", data)
		return
	} else {
		param := request.AjaxUpdateOrderPaymentRequest{}
		response := app.NewResponse(c)
		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			response.ToErrorResponse(errRsp)
			return
		}
		svc := service.New(c.Request.Context())
		err := svc.AjaxUpdateOrderPayment(&param)
		if err != nil {
			global.Logger.Errorf(c, "svc.GetOrderById err: %v", errs)
			response.ToErrorResponse(errcode.ErrorOrderNotFoundFail)
			return
		}

		response.ToSuccessResponse(nil)
		return
	}
}

func (o *OrderController) AjaxUpdateStatus(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		param := request.AjaxUpdateOrderGetRequest{}
		response := app.NewResponse(c)
		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			response.ToErrorResponse(errRsp)
			return
		}
		svc := service.New(c.Request.Context())
		order, err := svc.GetOrderById(param.ID)
		if err != nil {
			global.Logger.Errorf(c, "svc.GetOrderById err: %v", errs)
			response.ToErrorResponse(errcode.ErrorOrderNotFoundFail)
			return
		}

		var data gin.H
		orderStatus, err := svc.GetOrderStatusList()
		if err != nil {
			global.Logger.Errorf(c, "svc.GetOrderStatusList err: %v", err)
		}
		data = gin.H{
			"orderInfo":   order,
			"orderStatus": orderStatus,
		}

		c.HTML(http.StatusOK, "ajaxForm/order_ajax_update_status", data)
		return
	} else {
		param := request.AjaxUpdateOrderStatusRequest{}
		response := app.NewResponse(c)
		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			response.ToErrorResponse(errRsp)
			return
		}
		svc := service.New(c.Request.Context())
		err := svc.AjaxUpdateOrderStatus(&param)
		if err != nil {
			global.Logger.Errorf(c, "svc.GetOrderById err: %v", errs)
			response.ToErrorResponse(errcode.ErrorOrderNotFoundFail)
			return
		}

		response.ToSuccessResponse(nil)
		return
	}
}

func (o *OrderController) AjaxUpdateShipping(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		param := request.AjaxUpdateOrderGetRequest{}
		response := app.NewResponse(c)
		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			response.ToErrorResponse(errRsp)
			return
		}
		svc := service.New(c.Request.Context())
		order, err := svc.GetOrderById(param.ID)
		if err != nil {
			global.Logger.Errorf(c, "svc.GetOrderById err: %v", errs)
			response.ToErrorResponse(errcode.ErrorOrderNotFoundFail)
			return
		}

		var data gin.H
		orderShipping, err := svc.GetOrderShippingList()
		if err != nil {
			global.Logger.Errorf(c, "svc.GetOrderShippingList err: %v", err)
		}
		data = gin.H{
			"orderInfo":     order,
			"orderShipping": orderShipping,
		}

		c.HTML(http.StatusOK, "ajaxForm/order_ajax_update_shipping", data)
		return
	} else {
		param := request.AjaxUpdateOrderShippingRequest{}
		response := app.NewResponse(c)
		valid, errs := app.BindAndValid(c, &param)
		if !valid {
			global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
			errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
			response.ToErrorResponse(errRsp)
			return
		}
		svc := service.New(c.Request.Context())
		err := svc.AjaxUpdateOrderShipping(&param)
		if err != nil {
			global.Logger.Errorf(c, "svc.GetOrderById err: %v", errs)
			response.ToErrorResponse(errcode.ErrorOrderNotFoundFail)
			return
		}

		response.ToSuccessResponse(nil)
		return
	}
}
