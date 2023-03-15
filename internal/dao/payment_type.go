/*
 * @Author: GG
 * @Date: 2023-03-15 10:08:22
 * @LastEditTime: 2023-03-15 10:17:41
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\dao\payment_type.go
 *
 */
package dao

import (
	"oms/internal/model"
	"oms/pkg/enum"
)

// 支付类型
func (d Dao) GetPaymentTypeList() ([]*model.PaymentType, error) {
	var paymentTypeList []*model.PaymentType
	err := d.engine.Where("is_del = ?", enum.DEFAULT_IS_DEL).Find(&paymentTypeList).Error
	if err != nil {
		return nil, err
	}
	return paymentTypeList, nil
}
