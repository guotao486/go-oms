package model

import (
	"context"
	"oms/global"
)

type PaymentStatus struct {
	*Model
	Name string `gorm:"type:varchar(20);not null;comment:'名称'" json:"name"`
}

func init() {
	global.ModelAutoMigrate = append(global.ModelAutoMigrate, &PaymentStatus{})
	global.ModeInitData = append(global.ModeInitData, InitDataPaymentStatus)
}
func InitDataPaymentStatus() {
	var count int64
	m := NewPaymentStatus()
	err := global.DBEngine.Model(&m).Count(&count).Error
	if err != nil {
		global.Logger.Errorf(context.Background(), "model.PaymentStatus InitData err: %s", err)
	}
	if count == 0 {
		paymentStatuss := []PaymentStatus{
			{
				Name: "未支付",
			},
			{
				Name: "支付成功",
			},
			{
				Name: "支付失败",
			},
			{
				Name: "其它",
			},
		}
		for _, v := range paymentStatuss {
			global.DBEngine.Create(&v)
		}
	}
}

func NewPaymentStatus() *PaymentStatus {
	return &PaymentStatus{}
}
func (p *PaymentStatus) TableName() string {
	return global.DatabaseSetting.TablePrefix + "payment_status"
}
