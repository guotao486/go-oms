package response

import "oms/internal/model"

type OrderResponse struct {
	ID                uint32               `json:"id"`
	OrderNo           string               `json:"order_no"`
	WebsiteId         int32                `json:"-"`
	ShippingName      string               `json:"shipping_name"`
	ShippingTelephone string               `json:"shipping_telephone"`
	ShippingCountry   string               `json:"shipping_country"`
	ShippingProvince  string               `json:"shipping_province"`
	ShippingCity      string               `json:"shipping_city"`
	ShippingAddress   string               `json:"shipping_address"`
	ShippingZipcode   string               `json:"shipping_zipcode"`
	BillingName       string               `json:"billing_name"`
	OrderEmail        string               `json:"order_email"`
	OrderAmount       float32              `json:"order_amount"`
	OrderCurrency     int32                `json:"order_currency"`
	CurrencyInfo      *model.Currency      `json:"currencyInfo"`
	PaymentType       int32                `json:"-"`
	PaymentTypeInfo   *model.PaymentType   `gorm:"foreignKey:PaymentType" json:"payment_type_info"`
	PaymentStatus     int32                `json:"-"`
	PaymentStatusInfo *model.PaymentStatus `gorm:"foreignKey:PaymentStatus" json:"payment_status_info"`
	PaymentAccount    string               `json:"payment_account"`
	OrderShipping     int32                `json:"-"`
	OrderShippingInfo *model.OrderShipping `gorm:"foreignKey:OrderShipping" json:"order_shipping_info"`
	OrderStatus       int32                `json:"order_status"`
	OrderStatusInfo   *model.OrderStatus   `gorm:"foreignKey:OrderStatus" json:"order_status_info"`
}

func (o *OrderResponse) TableName() string {
	return model.NewOrder().TableName()
}
