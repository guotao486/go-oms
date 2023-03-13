package model

import "oms/global"

type OrderShipping struct {
	*Model
	Name string `gorm:"type:varchar(20);not null;comment:名称" json:"name"`
}

func (o *OrderShipping) TableName() string {
	return global.DatabaseSetting.TablePrefix + "order_shipping"
}
