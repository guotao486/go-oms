package dao

import "oms/internal/model"

// 创建订单
func (d *Dao) CreateOrder(order *model.Order) error {
	return d.engine.Model(&order).Create(&order).Error
}

func (d *Dao) UpdateOrder(order *model.Order) error {
	return d.engine.Model(&order).Save(&order).Error
}
