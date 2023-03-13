/*
 * @Author: GG
 * @Date: 2023-03-13 14:52:36
 * @LastEditTime: 2023-03-13 16:55:14
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\model\order.go
 *
 */
package model

import "oms/global"

type Order struct {
	*Model
	OrderNo           string  `gorm:"type:varchar(30);not null;comment:订单编号" json:"order_no"`
	WebsiteId         int32   `gorm:"default:0;comment:网站ID" json:"-"`
	ShippingName      string  `gorm:"type:varchar(50);not null;comment:收件人姓名" json:"shipping_name"`
	ShippingTelephone string  `gorm:"type:varchar(20);not null;comment:收件人电话" json:"shipping_telephone"`
	ShippingCountry   string  `gorm:"type:varchar(30);not null;comment:收件人国家" json:"shipping_country"`
	ShippingProvince  string  `gorm:"type:varchar(30);not null;comment:收件人省/州" json:"shipping_province"`
	ShippingCity      string  `gorm:"type:varchar(30);not null;comment:收件人城市" json:"shipping_city"`
	ShippingAddress   string  `gorm:"type:varchar(100);not null;comment:收件人地址" json:"shipping_address"`
	ShippingZipcode   string  `gorm:"type:varchar(10);not null;comment:收件人邮政编码" json:"shipping_zipcode"`
	BillingName       string  `gorm:"type:varchar(50);not null;comment:付款人姓名" json:"billing_name"`
	OrderEmail        string  `gorm:"type:varchar(30);not null;comment:客户邮箱" json:"order_email"`
	OrderAmount       float32 `gorm:"not null;comment:订单金额" json:"order_amount"`
	OrderCurrency     int32   `gorm:"default:1;not null;comment:订单货币" json:"order_currency"`
	PaymentType       int32   `gorm:"default:1;not null;comment:"支付类型" json:"payment_type"`
	PaymentStatus     int32   `gorm:"default:1;not null;comment:支付状态" json:"payment_status"`
	PaymentAccount    string  `gorm:"type:varchar(40);not null;comment:收款账号" json:"payment_account"`
	OrderShipping     int32   `gorm:"default:1;not null;comment:物流方式" json:"order_shipping"`
	OrderStatus       int32   `gorm:"default:1;not null;comment:订单状态" json:"order_status"`
	OrderItems        []*OrderItem
}

func NewOrder() *Order {
	return &Order{}
}

func (o *Order) TableName() string {
	return global.DatabaseSetting.TablePrefix + "order"
}
