/*
 * @Author: GG
 * @Date: 2023-03-13 15:40:06
 * @LastEditTime: 2023-03-14 14:14:30
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\model\payment_type.go
 *
 */
package model

import (
	"context"
	"oms/global"
)

type PaymentType struct {
	*Model
	Name string `gorm:"type:varchar(20);not null;comment:'名称'" json:"name"`
}

func init() {
	global.ModelAutoMigrate = append(global.ModelAutoMigrate, &PaymentType{})
	global.ModeInitData = append(global.ModeInitData, InitDataPaymentType)
}
func InitDataPaymentType() {
	var count int64
	m := NewPaymentType()
	err := global.DBEngine.Model(&m).Count(&count).Error
	if err != nil {
		global.Logger.Errorf(context.Background(), "model.PaymentType InitData err: %s", err)
	}
	if count == 0 {
		paymentType := []PaymentType{
			{
				Name: "未支付",
			},
			{
				Name: "支付宝",
			},
			{
				Name: "微信",
			},
			{
				Name: "PayPal",
			},
			{
				Name: "Zelle",
			},
			{
				Name: "WintoPay",
			},
			{
				Name: "其它",
			},
		}
		for _, v := range paymentType {
			global.DBEngine.Create(&v)
		}
	}
}

func NewPaymentType() *PaymentType {
	return &PaymentType{}
}
func (p *PaymentType) TableName() string {
	return global.DatabaseSetting.TablePrefix + "payment_type"
}
