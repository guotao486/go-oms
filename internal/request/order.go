/*
 * @Author: GG
 * @Date: 2023-03-16 10:48:14
 * @LastEditTime: 2023-03-22 11:22:02
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\request\order.go
 *
 */
package request

type GetOrderListRequest struct {
	OrderNo           string `form:"order_no" `
	WebsiteId         int32  `form:"website_id" `
	ShippingName      string `form:"shipping_name" `
	ShippingTelephone string `form:"shipping_telephone" `
	ShippingCountry   string `form:"shipping_country" `
	ShippingProvince  string `form:"shipping_province"`
	ShippingCity      string `form:"shipping_city"`
	ShippingAddress   string `form:"shipping_address"`
	ShippingZipcode   string `form:"shipping_zipcode"`
	BillingName       string `form:"billing_name"`
	OrderEmail        string `form:"order_email"`
	OrderCurrency     int32  `form:"order_currency"`
	PaymentType       int32  `form:"payment_type"`
	PaymentStatus     int32  `form:"payment_status"`
	PaymentAccount    string `form:"payment_account"`
	OrderShipping     int32  `form:"order_shipping"`
	OrderStatus       int32  `form:"order_status"`
}
type CreateOrderRequest struct {
	OrderNo           string                                `form:"order_no" binding:"required,min=2" label:"订单号"`
	WebsiteId         int32                                 `form:"website_id" bingding:"required" label:"网站"`
	ShippingName      string                                `form:"shipping_name" binding:"required,min=2" label:"收件人姓名"`
	ShippingTelephone string                                `form:"shipping_telephone" binding:"required,min=2" label:"收件人电话"`
	ShippingCountry   string                                `form:"shipping_country" binding:"required" label:"收件人国家"`
	ShippingProvince  string                                `form:"shipping_province" binding:"required" label:"收件人省/州"`
	ShippingCity      string                                `form:"shipping_city" binding:"required" label:"收件人城市"`
	ShippingAddress   string                                `form:"shipping_address" binding:"required" label:"收件人地址"`
	ShippingZipcode   string                                `form:"shipping_zipcode" binding:"required" label:"邮政编码"`
	BillingName       string                                `form:"billing_name" binding:"required" label:"付款人姓名"`
	OrderEmail        string                                `form:"order_email" binding:"required,email" label:"客户邮箱"`
	OrderAmount       float32                               `form:"order_amount" binding:"required,numeric" label:"订单金额"`
	DiscountAmount    float32                               `form:"discount_amount" binding:"numeric"`
	ShippingAmount    float32                               `form:"shipping_amount" binding:"numeric"`
	CouponAmount      float32                               `form:"coupon_amount" binding:"numeric"`
	OrderCurrency     int32                                 `form:"order_currency" binding:"required" label:"订单货币"`
	PaymentType       int32                                 `form:"payment_type" binding:"required" label:"支付类型"`
	PaymentStatus     int32                                 `form:"payment_status" binding:"required" label:"支付状态"`
	PaymentAccount    string                                `form:"payment_account" binding:"required" label:"收款账号"`
	OrderShipping     int32                                 `form:"order_shipping" binding:"required" label:"物流方式"`
	OrderStatus       int32                                 `form:"order_status" binding:"required" label:"订单状态"`
	OrderProducts     map[string]*CreateOrderProductRequest `binding:"required" label:"订单商品"`
	Remarks           string                                `form:"remarks" label:"订单备注"`
}

type UpdateOrderGetRequest struct {
	ID uint32 `form:"id" binding:"required" label:"订单编号"` // id
}
type UpdateOrderPostRequest struct {
	ID                uint32                                `form:"id" binding:"required" label:"订单编号"` // id
	OrderNo           string                                `form:"order_no" binding:"required,min=2" label:"订单号"`
	WebsiteId         int32                                 `form:"website_id" bingding:"required" label:"网站"`
	ShippingName      string                                `form:"shipping_name" binding:"required,min=2" label:"收件人姓名"`
	ShippingTelephone string                                `form:"shipping_telephone" binding:"required,min=2" label:"收件人电话"`
	ShippingCountry   string                                `form:"shipping_country" binding:"required" label:"收件人国家"`
	ShippingProvince  string                                `form:"shipping_province" binding:"required" label:"收件人省/州"`
	ShippingCity      string                                `form:"shipping_city" binding:"required" label:"收件人城市"`
	ShippingAddress   string                                `form:"shipping_address" binding:"required" label:"收件人地址"`
	ShippingZipcode   string                                `form:"shipping_zipcode" binding:"required" label:"邮政编码"`
	BillingName       string                                `form:"billing_name" binding:"required" label:"付款人姓名"`
	OrderEmail        string                                `form:"order_email" binding:"required,email" label:"客户邮箱"`
	OrderAmount       float32                               `form:"order_amount" binding:"required,numeric" label:"订单金额"`
	DiscountAmount    float32                               `form:"discount_amount" binding:"numeric"`
	ShippingAmount    float32                               `form:"shipping_amount" binding:"numeric"`
	CouponAmount      float32                               `form:"coupon_amount" binding:"numeric"`
	OrderCurrency     int32                                 `form:"order_currency" binding:"required" label:"订单货币"`
	PaymentType       int32                                 `form:"payment_type" binding:"required" label:"支付类型"`
	PaymentStatus     int32                                 `form:"payment_status" binding:"required" label:"支付状态"`
	PaymentAccount    string                                `form:"payment_account" binding:"required" label:"收款账号"`
	OrderShipping     int32                                 `form:"order_shipping" binding:"required" label:"物流方式"`
	OrderStatus       int32                                 `form:"order_status" binding:"required" label:"订单状态"`
	OrderProducts     map[string]*CreateOrderProductRequest `binding:"required" label:"订单商品"`
	Remarks           string                                `form:"remarks" label:"订单备注"`
}
type CreateOrderProductRequest struct {
	Name      string `form:"name" binding:"required"`
	Sku       string `form:"sku" binding:"required"`
	Attribute string `form:"attribute" binding:"required"`
	Images    string `form:"image" binding:"required"`
}

// 删除
type DeleteOrderRequest struct {
	ID uint32 `uri:"id" form:"id" json:"id" binding:"required,gte=1" label:"订单ID"`
}
