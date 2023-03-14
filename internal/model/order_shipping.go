package model

import (
	"context"
	"oms/global"
)

type OrderShipping struct {
	*Model
	Name string `gorm:"type:varchar(20);not null;comment:'名称'" json:"name"`
}

func init() {
	global.ModelAutoMigrate = append(global.ModelAutoMigrate, &OrderShipping{})
	global.ModeInitData = append(global.ModeInitData, InitDataOrderShipping)
}

func InitDataOrderShipping() {
	var count int64
	m := NewOrderShipping()
	err := global.DBEngine.Model(&m).Count(&count).Error
	if err != nil {
		global.Logger.Errorf(context.Background(), "model.OrderShipping InitData err: %s", err)
	}
	if count == 0 {
		orderShippings := []OrderShipping{
			{
				Name: "DHL",
			},
		}
		for _, v := range orderShippings {
			global.DBEngine.Create(&v)
		}
	}
}

func NewOrderShipping() *OrderShipping {
	return &OrderShipping{}
}

func (o *OrderShipping) TableName() string {
	return global.DatabaseSetting.TablePrefix + "order_shipping"
}
