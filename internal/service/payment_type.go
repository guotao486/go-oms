package service

import (
	"fmt"
	"oms/global"
	"oms/internal/model"
	"oms/pkg/convert"
)

func (s Service) GetPaymentTypeList() ([]*model.PaymentType, error) {
	cache := global.CacheStore.Engine
	var typeList []*model.PaymentType

	cacheList, err := cache.Get(model.CachePaymentTypeListKey)
	if err != nil {
		return nil, err
	}

	// 没有缓存则从数据库查询并写入缓存
	if cacheList == nil {
		fmt.Println("cache empty")
		typeList, err = s.dao.GetPaymentTypeList()
		if err != nil {
			return nil, err
		}
		st := &convert.StructTo{V: typeList}
		buf, err := st.StructToBytes()
		if err != nil {
			return typeList, err
		}
		err = cache.Set(model.CachePaymentTypeListKey, buf)
		return typeList, err
	}
	fmt.Println("cache not empty")
	// 将缓存中的字节转成结构体
	bs := &convert.ByteTo{V: &typeList}
	err = bs.ByteToStruct(cacheList)
	if err != nil {
		return nil, err
	}
	return typeList, nil
}
