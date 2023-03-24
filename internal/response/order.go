/*
 * @Author: GG
 * @Date: 2023-03-13 16:32:06
 * @LastEditTime: 2023-03-23 09:28:03
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\response\order.go
 *
 */
package response

import (
	"encoding/json"
	"oms/internal/model"
	"oms/pkg/util"
)

type OrderResponse struct {
	ID                uint32                `json:"id"`
	OrderNo           string                `json:"order_no"`
	WebsiteId         int32                 `json:"-"`
	ShippingName      string                `json:"shipping_name"`
	ShippingTelephone string                `json:"shipping_telephone"`
	ShippingCountry   string                `json:"shipping_country"`
	ShippingProvince  string                `json:"shipping_province"`
	ShippingCity      string                `json:"shipping_city"`
	ShippingAddress   string                `json:"shipping_address"`
	ShippingZipcode   string                `json:"shipping_zipcode"`
	BillingName       string                `json:"billing_name"`
	OrderEmail        string                `json:"order_email"`
	OrderAmount       float32               `json:"order_amount"`
	DiscountAmount    float32               `json:"discount_amount"`
	ShippingAmount    float32               `json:"shipping_amount"`
	CouponAmount      float32               `json:"coupon_amount"`
	OrderCurrency     int32                 `json:"-"`
	CurrencyInfo      *model.Currency       `gorm:"foreignKey:OrderCurrency" json:"currency_info"`
	PaymentType       int32                 `json:"-"`
	PaymentTypeInfo   *model.PaymentType    `gorm:"foreignKey:PaymentType" json:"payment_type_info"`
	PaymentStatus     int32                 `json:"-"`
	PaymentStatusInfo *model.PaymentStatus  `gorm:"foreignKey:PaymentStatus" json:"payment_status_info"`
	PaymentAccount    string                `json:"payment_account"`
	OrderShipping     int32                 `json:"-"`
	OrderShippingInfo *model.OrderShipping  `gorm:"foreignKey:OrderShipping" json:"order_shipping_info"`
	OrderStatus       int32                 `json:"-"`
	OrderStatusInfo   *model.OrderStatus    `gorm:"foreignKey:OrderStatus" json:"order_status_info"`
	OrderProducts     []*model.OrderProduct `gorm:"foreignKey:OrderID" json:"order_products"`
	CreatedOn         uint32                `json:"created_on"`
}

func (o *OrderResponse) TableName() string {
	return model.NewOrder().TableName()
}

func (o *OrderResponse) MarshalJSON() ([]byte, error) {
	type Alias OrderResponse
	return json.Marshal(&struct {
		CreatedOn string `json:"created_on"`
		*Alias
	}{
		CreatedOn: util.UnixToString(int64(o.CreatedOn)),
		Alias:     (*Alias)(o),
	})
}
