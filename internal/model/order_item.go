package model

import "oms/global"

type OrderItem struct {
	*Model
	OrderID   uint32 `gorm:"not null;comment:'订单ID'" json:"order_id"`
	Name      string `gorm:"type:varchar(50);not null;comment:'商品名称'" json:"name"`
	Sku       string `gorm:"type:varchar(50);not null;comment:'sku'" json:"sku"`
	Attribute string `gorm:"type:varchar(255);comment:'商品属性'" json:"attribute"`
}

func init() {
	global.ModelAutoMigrate = append(global.ModelAutoMigrate, &OrderItem{})
}
func (o *OrderItem) TableName() string {
	return global.DatabaseSetting.TablePrefix + "order_item"
}
