/*
 * @Author: GG
 * @Date: 2023-03-13 16:41:02
 * @LastEditTime: 2023-03-16 10:57:24
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\model\order_product.go
 *
 */
package model

import "oms/global"

type OrderProduct struct {
	*Model
	OrderID   uint32 `gorm:"not null;comment:'订单ID'" json:"order_id"`
	Name      string `gorm:"type:varchar(50);not null;comment:'商品名称'" json:"name"`
	Sku       string `gorm:"type:varchar(50);not null;comment:'sku'" json:"sku"`
	Attribute string `gorm:"type:varchar(255);comment:'商品属性'" json:"attribute"`
	Images    string `gorm:"type:varchar(255);comment:'商品图片'" json:"images"`
}

func init() {
	global.ModelAutoMigrate = append(global.ModelAutoMigrate, &OrderProduct{})
}
func (o *OrderProduct) TableName() string {
	return global.DatabaseSetting.TablePrefix + "order_product"
}
