package model

import (
	"context"
	"oms/global"
)

type OrderStatus struct {
	*Model
	Name string `gorm:"type:varchar(20);not null;comment:'名称'" json:"name"`
}

var CacheOrderStatusListKey = "cache_order_status_list"

func init() {
	global.ModelAutoMigrate = append(global.ModelAutoMigrate, &OrderStatus{})
	global.ModeInitData = append(global.ModeInitData, InitDataOrderStatus)
}
func InitDataOrderStatus() {
	var count int64
	m := NewOrderStatus()
	err := global.DBEngine.Model(&m).Count(&count).Error
	if err != nil {
		global.Logger.Errorf(context.Background(), "model.OrderStatus InitData err: %s", err)
	}
	if count == 0 {
		orderStatuss := []OrderStatus{
			{
				Name: "新创建",
			},
			{
				Name: "备货中",
			},
			{
				Name: "缺货中",
			},
			{
				Name: "已发货",
			},
			{
				Name: "已退款/取消",
			},
			{
				Name: "已完结",
			},
		}
		for _, v := range orderStatuss {
			global.DBEngine.Create(&v)
		}
	}
}

func NewOrderStatus() *OrderStatus {
	return &OrderStatus{}
}
func (o *OrderStatus) TableName() string {
	return global.DatabaseSetting.TablePrefix + "order_status"
}
