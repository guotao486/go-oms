/*
 * @Author: GG
 * @Date: 2023-03-15 15:37:05
 * @LastEditTime: 2023-03-15 15:41:26
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\service\order_shipping.go
 *
 */
package service

import (
	"oms/global"
	"oms/internal/model"
	"oms/pkg/convert"
)

func (s Service) GetOrderShippingList() ([]*model.OrderShipping, error) {
	cache := global.CacheStore.Engine
	var shippingList []*model.OrderShipping

	cacheList, err := cache.Get(model.CacheOrderShippingListKey)
	if err != nil {
		return nil, err
	}

	// 没有缓存则从数据库查询并写入缓存
	if cacheList == nil {
		shippingList, err = s.dao.GetOrderShippingList()
		if err != nil {
			return nil, err
		}
		st := &convert.StructTo{V: shippingList}
		buf, err := st.StructToBytes()
		if err != nil {
			return shippingList, err
		}
		err = cache.Set(model.CacheOrderShippingListKey, buf)
		return shippingList, err
	}

	// 将缓存中的字节转成结构体
	bs := &convert.ByteTo{V: &shippingList}
	err = bs.ByteToStruct(cacheList)
	if err != nil {
		return nil, err
	}
	return shippingList, nil
}
