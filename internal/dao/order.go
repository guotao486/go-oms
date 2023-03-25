/*
 * @Author: GG
 * @Date: 2023-03-14 09:22:26
 * @LastEditTime: 2023-03-22 16:40:30
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\dao\order.go
 *
 */
/*
 * @Author: GG
 * @Date: 2023-03-14 09:22:26
 * @LastEditTime: 2023-03-21 10:38:30
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\dao\order.go
 *
 */
package dao

import (
	"oms/internal/model"
	"oms/internal/response"
)

// 创建订单
func (d *Dao) CreateOrder(order *model.Order) error {
	return d.engine.Model(&order).Create(&order).Error
}

// 更新订单
func (d *Dao) UpdateOrder(order *model.Order) error {
	return d.engine.Model(&order).Save(&order).Error
}

func (d *Dao) UpdateOrderProducts(order *model.Order) error {
	return d.engine.Model(&order).Association("OrderProducts").Replace(order.OrderProducts).Error
}

// 根据ID 查询订单
func (d *Dao) GetOrderById(id uint32) (*model.Order, error) {
	order := model.NewOrder()
	err := d.engine.Model(&order).Where("id = ?", id).Scopes(IsDelToUnable).Preload("OrderProducts").First(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

// 根据orderNo 查询订单
func (d *Dao) GetOrderByOrderNo(orderNo string) (*model.Order, error) {
	order := model.NewOrder()
	err := d.engine.Model(&order).Where("order_no = ?", orderNo).Scopes(IsDelToUnable).Preload("OrderProducts").First(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

// 获取用户数量
func (d *Dao) GetOrderListCount(order *model.Order) (int, error) {
	var count int
	db := d.engine.Model(&order)

	// 订单号
	if order.OrderNo != "" {
		db = db.Where("order_no = ?", order.OrderNo)
	}

	// 手机号
	if order.ShippingTelephone != "" {
		db = db.Where("shipping_telephone = ?", order.ShippingTelephone)
	}
	// 邮箱
	if order.OrderEmail != "" {
		db = db.Where("order_email = ?", order.OrderEmail)
	}
	// 货币类型
	if order.OrderCurrency != 0 {
		db = db.Where("order_currency = ?", order.OrderCurrency)
	}
	// 支付类型
	if order.PaymentType != 0 {
		db = db.Where("payment_type = ?", order.PaymentType)
	}
	// 支付状态
	if order.PaymentStatus != 0 {
		db = db.Where("payment_status = ?", order.PaymentStatus)
	}
	// 订单状态
	if order.OrderStatus != 0 {
		db = db.Where("order_status = ?", order.OrderStatus)
	}
	// 物流类型
	if order.OrderShipping != 0 {
		db = db.Where("order_shipping = ?", order.OrderShipping)
	}
	// 收款账号
	if order.PaymentAccount != "" {
		db = db.Where("payment_account like ?", d.WhereLikeString(order.PaymentAccount))
	}
	// 客户信息
	if order.ShippingName != "" {
		db = db.Where("shipping_name like ?", d.WhereLikeString(order.ShippingName))
	}
	if order.ShippingCountry != "" {
		db = db.Where("shipping_country like ?", d.WhereLikeString(order.ShippingCountry))
	}
	if order.ShippingProvince != "" {
		db = db.Where("shipping_province like ?", d.WhereLikeString(order.ShippingProvince))
	}
	if order.ShippingCity != "" {
		db = db.Where("shipping_city like ?", d.WhereLikeString(order.ShippingCity))
	}
	if order.ShippingAddress != "" {
		db = db.Where("shipping_address like ?", d.WhereLikeString(order.ShippingAddress))
	}
	if order.BillingName != "" {
		db = db.Where("billing_name like ?", d.WhereLikeString(order.BillingName))
	}
	if err := db.Scopes(IsDelToUnable).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// 分页列表查询
func (d *Dao) GetOrderListPages(order *model.Order, pageOffest, pageSize int) ([]*response.OrderResponse, error) {
	var orders []*response.OrderResponse
	db := d.engine.Model(&order)

	// 订单号
	if order.OrderNo != "" {
		db = db.Where("order_no = ?", order.OrderNo)
	}

	// 手机号
	if order.ShippingTelephone != "" {
		db = db.Where("shipping_telephone = ?", order.ShippingTelephone)
	}
	// 邮箱
	if order.OrderEmail != "" {
		db = db.Where("order_email = ?", order.OrderEmail)
	}
	// 货币类型
	if order.OrderCurrency != 0 {
		db = db.Where("order_currency = ?", order.OrderCurrency)
	}
	// 支付类型
	if order.PaymentType != 0 {
		db = db.Where("payment_type = ?", order.PaymentType)
	}
	// 支付状态
	if order.PaymentStatus != 0 {
		db = db.Where("payment_status = ?", order.PaymentStatus)
	}
	// 订单状态
	if order.OrderStatus != 0 {
		db = db.Where("order_status = ?", order.OrderStatus)
	}
	// 物流类型
	if order.OrderShipping != 0 {
		db = db.Where("order_shipping = ?", order.OrderShipping)
	}
	// 收款账号
	if order.PaymentAccount != "" {
		db = db.Where("payment_account like ?", d.WhereLikeString(order.PaymentAccount))
	}
	// 客户信息
	if order.ShippingName != "" {
		db = db.Where("shipping_name like ?", d.WhereLikeString(order.ShippingName))
	}
	if order.ShippingCountry != "" {
		db = db.Where("shipping_country like ?", d.WhereLikeString(order.ShippingCountry))
	}
	if order.ShippingProvince != "" {
		db = db.Where("shipping_province like ?", d.WhereLikeString(order.ShippingProvince))
	}
	if order.ShippingCity != "" {
		db = db.Where("shipping_city like ?", d.WhereLikeString(order.ShippingCity))
	}
	if order.ShippingAddress != "" {
		db = db.Where("shipping_address like ?", d.WhereLikeString(order.ShippingAddress))
	}
	if order.BillingName != "" {
		db = db.Where("billing_name like ?", d.WhereLikeString(order.BillingName))
	}

	// 分页查询
	if pageOffest >= 0 && pageSize > 0 {
		db = db.Offset(pageOffest).Limit(pageSize)
	}
	if err := db.Scopes(PreloadAll, IsDelToUnable).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

// 删除用户，有delete_on 和 is_del 字段则是软删除
func (d *Dao) DeleteOrder(order *model.Order) error {
	return d.engine.Model(&order).Scopes(IsDelToUnable).Delete(&order).Error
}
