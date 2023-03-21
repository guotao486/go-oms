/*
 * @Author: GG
 * @Date: 2023-03-17 14:22:12
 * @LastEditTime: 2023-03-21 14:54:27
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\service\order.go
 *
 */
package service

import (
	"oms/internal/model"
	"oms/internal/request"
	"oms/pkg/app"
	"oms/pkg/errcode"
)

// 根据id获取数据
func (s Service) GetOrderById(id uint32) (*model.Order, error) {
	return s.dao.GetOrderById(id)
}

// 创建订单
func (s Service) CreateOrder(param *request.CreateOrderRequest) error {
	order := model.NewOrder()
	app.StructAssign(order, param)

	var products []*model.OrderProduct
	if len(param.OrderProducts) == 0 {
		return errcode.ErrorOrderProductNotEmptyFail
	}
	for _, p := range param.OrderProducts {
		product := model.NewOrderProduct()
		app.StructAssign(product, p)
		if product == nil {
			return errcode.ErrorAddOrderProductFail
		}
		products = append(products, product)
	}

	order.OrderProducts = products
	return s.dao.CreateOrder(order)
}

// 更新订单
func (s Service) UpdateOrder(param *request.UpdateOrderPostRequest) error {
	order, err := s.dao.GetOrderById(param.ID)
	if err != nil {
		return errcode.ErrorOrderNotFoundFail
	}
	app.StructAssign(order, param)

	var products []*model.OrderProduct
	for _, p := range param.OrderProducts {
		product := model.NewOrderProduct()
		app.StructAssign(product, p)
		products = append(products, product)
	}
	order.OrderProducts = products
	return s.dao.UpdateOrder(order)
}
