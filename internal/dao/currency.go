package dao

import (
	"oms/internal/model"
	"oms/pkg/enum"
)

// 获取货币信息
func (d Dao) GetCurrencyList() ([]*model.Currency, error) {
	var currencyList []*model.Currency
	err := d.engine.Where("is_del = ?", enum.DEFAULT_IS_DEL).Find(&currencyList).Error
	if err != nil {
		return nil, err
	}
	return currencyList, nil
}
