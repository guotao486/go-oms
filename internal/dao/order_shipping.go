/*
 * @Author: GG
 * @Date: 2023-03-15 09:57:11
 * @LastEditTime: 2023-03-15 10:02:10
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\dao\order_shipping.go
 *
 */
package dao

import (
	"oms/internal/model"
	"oms/pkg/enum"
)

// 获取物流信息
func (d Dao) GetOrderShippingList() ([]*model.OrderShipping, error) {
	var orderShippingList []*model.OrderShipping
	err := d.engine.Where("is_del = ?", enum.DEFAULT_IS_DEL).Find(&orderShippingList).Error
	if err != nil {
		return nil, err
	}
	return orderShippingList, nil
}
