package model

import "oms/global"

type Currency struct {
	*Model
	Name string  `gorm:"type:varchar(20);not null;comment:名称" json:"name"`
	Rate float32 `gorm:"not null;comment:汇率" json:"rate"`
}

func (u *Currency) TableName() string {
	return global.DatabaseSetting.TablePrefix + "currency"
}
