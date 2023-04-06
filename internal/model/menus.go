/*
 * @Author: GG
 * @Date: 2023-03-30 10:42:26
 * @LastEditTime: 2023-03-30 14:30:09
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\model\menus.go
 *
 */
package model

import "oms/global"

type Menus struct {
	*Model
	Title     string `gorm:"not null;" json:"title"`
	Router    string `gorm:"not null;" json:"router"`
	Sort      uint8  `gorm:"default:0" json:"sort"`
	Role      string `gorm:"not null;" json:"role"`
	ParentID  uint32 `gorm:"default:0;" json:"parent_id"`
	State     uint8  `gorm:"default:1;not null" json:"state"`
	ChildNode []*Menus
}

var CacheMenusListKey = "cache_menus_list"
var CacheParentMenusListKey = "cache_parent_menus_list"

func init() {
	global.ModelAutoMigrate = append(global.ModelAutoMigrate, &Menus{})
}

func NewMenus() *Menus {
	return &Menus{}
}

func (m Menus) TableName() string {
	return global.DatabaseSetting.TablePrefix + "menus"
}
