package dao

import (
	"oms/internal/model"
	"oms/pkg/enum"
)

// 支付状态列表
func (d Dao) GetPaymentStatusList() ([]*model.PaymentStatus, error) {
	var paymentStatusList []*model.PaymentStatus
	err := d.engine.Where("is_del = ?", enum.DEFAULT_IS_DEL).Find(&paymentStatusList).Error
	if err != nil {
		return nil, err
	}
	return paymentStatusList, nil
}
