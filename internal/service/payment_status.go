package service

import (
	"oms/global"
	"oms/internal/model"
	"oms/pkg/convert"
)

func (s Service) GetPaymentStatusList() ([]*model.PaymentStatus, error) {
	cache := global.CacheStore.Engine
	var statusList []*model.PaymentStatus

	cacheList, err := cache.Get(model.CachePaymentStatusListKey)

	// 没有缓存则从数据库查询并写入缓存
	if cacheList == nil {
		statusList, err = s.dao.GetPaymentStatusList()
		if err != nil {
			return nil, err
		}
		st := &convert.StructTo{V: statusList}
		buf, err := st.StructToBytes()
		if err != nil {
			return statusList, err
		}
		err = cache.Set(model.CachePaymentStatusListKey, buf)
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
