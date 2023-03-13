package model

import "oms/global"

type OrderStatus struct {
	*Model
	Name string `gorm:"type:varchar(20);not null;comment:名称" json:"name"`
}

func (o *OrderStatus) TableName() string {
	return global.DatabaseSetting.TablePrefix + "order_status"
}
