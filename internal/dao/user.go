/*
 * @Author: GG
 * @Date: 2023-02-28 11:41:40
 * @LastEditTime: 2023-03-01 10:51:54
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\dao\user.go
 *
 */
package dao

import (
	"oms/internal/model"
)

func (d *Dao) CreateUser(user *model.User) error {
	return d.engine.Model(&user).Create(&user).Error
}
