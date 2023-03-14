/*
 * @Author: GG
 * @Date: 2023-03-13 15:35:39
 * @LastEditTime: 2023-03-14 14:35:50
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\model\currency.go
 *
 */
package model

import (
	"context"
	"oms/global"
)

type Currency struct {
	*Model
	Name string  `gorm:"type:varchar(20);not null;comment:'名称'" json:"name"`
	Rate float32 `gorm:"default:0;not null;comment:'汇率'" json:"rate"`
}

func init() {
	global.ModelAutoMigrate = append(global.ModelAutoMigrate, &Currency{})
	global.ModeInitData = append(global.ModeInitData, InitDataCurreny)
}

func InitDataCurreny() {
	var count int64

	m := NewCurrency()
	err := global.DBEngine.Model(&m).Count(&count).Error
	if err != nil {
		global.Logger.Errorf(context.Background(), "model.currency InitData err: %s", err)
	}
	if count == 0 {
		currencys := []Currency{
			{
				Name: "USD",
			},
			{
				Name: "RMB",
			},
		}
		for _, v := range currencys {
			global.DBEngine.Create(&v)
		}
	}
}

func NewCurrency() *Currency {
	return &Currency{}
}

func (u *Currency) TableName() string {
	return global.DatabaseSetting.TablePrefix + "currency"
}
