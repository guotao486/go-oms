/*
 * @Author: GG
 * @Date: 2023-02-28 10:57:20
 * @LastEditTime: 2023-02-28 15:32:48
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\model\user_group.go
 *
 */
package model

import (
	"oms/global"
	"oms/pkg/enum"
)

type UserGroup struct {
	*Model
	Title  string `gorm:"type:varchar(100);not null" json:"title"`
	State  uint8  `gorm:"default:1;not null" json:"state"`
	Leader uint8  `gorm:"default:0;not null" json:"leader"`
}

func init() {
	global.ModelAutoMigrate = append(global.ModelAutoMigrate, &UserGroup{})
}

func NewUserGroup() *UserGroup {
	return &UserGroup{
		State:  enum.DEFAULT_STATE,
		Leader: enum.DEFAULT,
	}
}

func (u UserGroup) TableName() string {
	return "oms_user_group"
}
