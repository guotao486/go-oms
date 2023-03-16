/*
 * @Author: GG
 * @Date: 2023-03-15 16:32:14
 * @LastEditTime: 2023-03-16 09:39:19
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\service\order_status.go
 *
 */
/*
 * @Author: GG
 * @Date: 2023-03-15 16:32:14
 * @LastEditTime: 2023-03-15 16:37:17
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\service\order_status.go
 *
 */
package service

import (
	"oms/global"
	"oms/internal/model"
	"oms/pkg/convert"
)

func (s Service) GetOrderStatusList() ([]*model.OrderStatus, error) {
	cache := global.CacheStore.Engine
	var statusList []*model.OrderStatus

	cacheList, err := cache.Get(model.CacheOrderStatusListKey)

	// 没有缓存则从数据库查询并写入缓存
	if cacheList == nil {
		statusList, err = s.dao.GetOrderStatusList()
		if err != nil {
			return nil, err
		}
		st := &convert.StructTo{V: statusList}
		buf, err := st.StructToBytes()
		if err != nil {
			return statusList, err
		}
		err = cache.Set(model.CacheOrderStatusListKey, buf)
		return statusList, err
	}

	// 将缓存中的字节转成结构体
	bs := &convert.ByteTo{V: &statusList}
	err = bs.ByteToStruct(cacheList)
	if err != nil {
		return nil, err
	}
	return statusList, nil
}
