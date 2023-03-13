package model

import "oms/global"

type PaymentStatus struct {
	*Model
	Name string `gorm:"type:varchar(20);not null;comment:名称" json:"name"`
}

func (p *PaymentStatus) TableName() string {
	return global.DatabaseSetting.TablePrefix + "payment_status"
}
