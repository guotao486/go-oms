/*
 * @Author: GG
 * @Date: 2023-03-17 14:22:12
 * @LastEditTime: 2023-03-22 17:26:58
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\service\order.go
 *
 */
package service

import (
	"fmt"
	"oms/global"
	"oms/internal/model"
	"oms/internal/request"
	"oms/internal/response"
	"oms/pkg/app"
	"oms/pkg/errcode"

	"github.com/jinzhu/gorm"
)

// 根据id获取数据
func (s Service) GetOrderById(id uint32) (*model.Order, error) {
	return s.dao.GetOrderById(id)
}

// 获取用户总数量
func (s *Service) GetOrderCountList(param *request.GetOrderListRequest) (int, error) {
	order := model.NewOrder()
	app.StructAssign(order, param)
	return s.dao.GetOrderListCount(order)
}

// 获取用户分页列表
func (s *Service) GetOrderListPager(param *request.GetOrderListRequest, pager *app.Pager) ([]*response.OrderResponse, error) {
	order := model.NewOrder()
	app.StructAssign(order, param)
	pageOffset := app.GetPageOffset(pager.Page, pager.PageSize)

	var orders []*response.OrderResponse
	orders, err := s.dao.GetOrderListPages(order, pageOffset, pager.PageSize)
	return orders, err
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
	return global.DBEngine.Transaction(func(tx *gorm.DB) error {
		s.SetDao(tx)
		order, err := s.dao.GetOrderById(param.ID)
		if err != nil {
			return errcode.ErrorOrderNotFoundFail
		}
		app.StructAssign(order, param)
		fmt.Printf("param.CouponAmount: %v\n", param.CouponAmount)
		fmt.Printf("1 order.CouponAmount: %v\n", order.CouponAmount)

		err = s.dao.UpdateOrder(order)
		if err != nil {
			return err
		}

		var products []*model.OrderProduct
		for _, p := range param.OrderProducts {
			product := model.NewOrderProduct()
			app.StructAssign(product, p)
			products = append(products, product)
		}
		order.OrderProducts = products
		fmt.Printf("2 order.CouponAmount: %v\n", order.CouponAmount)
		return s.dao.UpdateOrderProducts(order)
	})
}

func (s *Service) DeleteOrder(param *request.DeleteOrderRequest) error {
	// 检查是否有该用户
	order, err := s.dao.GetOrderById(param.ID)
	if err != nil {
		return errcode.ErrorOrderNotFoundFail
	}

	return s.dao.DeleteOrder(order)
}
