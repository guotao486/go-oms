/*
 * @Author: GG
 * @Date: 2023-03-15 10:02:54
 * @LastEditTime: 2023-03-15 10:06:33
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\dao\order_status.go
 *
 */
package dao

import (
	"oms/internal/model"
	"oms/pkg/enum"
)

// 订单状态列表
func (d Dao) GetOrderStatusList() ([]*model.OrderStatus, error) {
	var orderStatusList []*model.OrderStatus
	err := d.engine.Where("is_del = ?", enum.DEFAULT_IS_DEL).Find(&orderStatusList).Error
	if err != nil {
		return nil, err
	}
	return orderStatusList, nil
}
