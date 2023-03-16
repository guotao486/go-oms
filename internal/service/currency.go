package service

import (
	"oms/global"
	"oms/internal/model"
	"oms/pkg/convert"
)

func (s Service) GetCurrencyList() ([]*model.Currency, error) {
	cache := global.CacheStore.Engine
	var currencyList []*model.Currency

	cacheList, err := cache.Get(model.CacheCurrencyListKey)

	// 没有缓存则从数据库查询并写入缓存
	if cacheList == nil {
		currencyList, err = s.dao.GetCurrencyList()
		if err != nil {
			return nil, err
		}
		st := &convert.StructTo{V: currencyList}
		buf, err := st.StructToBytes()
		if err != nil {
			return currencyList, err
		}
		err = cache.Set(model.CacheCurrencyListKey, buf)
		return currencyList, err
	}

	// 将缓存中的字节转成结构体
	bs := &convert.ByteTo{V: &currencyList}
	err = bs.ByteToStruct(cacheList)
	if err != nil {
		return nil, err
	}
	return currencyList, nil
}
