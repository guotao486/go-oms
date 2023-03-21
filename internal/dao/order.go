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

import "oms/internal/model"

// 创建订单
func (d *Dao) CreateOrder(order *model.Order) error {
	return d.engine.Model(&order).Create(&order).Error
}

// 更新订单
func (d *Dao) UpdateOrder(order *model.Order) error {
	return d.engine.Model(&order).Save(&order).Error
}

// 根据ID 查询订单
func (d *Dao) GetOrderById(id uint32) (*model.Order, error) {
	order := model.NewOrder()
	err := d.engine.Model(&order).Where("id = ?", id).Preload("OrderProducts").First(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

// 根据orderNo 查询订单
func (d *Dao) GetOrderByOrderNo(orderNo string) (*model.Order, error) {
	order := model.NewOrder()
	err := d.engine.Model(&order).Where("order_no = ?", orderNo).Preload("OrderProducts").First(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}
