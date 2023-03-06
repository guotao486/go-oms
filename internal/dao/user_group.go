/*
 * @Author: GG
 * @Date: 2023-03-06 15:22:05
 * @LastEditTime: 2023-03-06 15:26:09
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\dao\user_group.go
 *
 */
package dao

import "oms/internal/model"

// 创建用户组
func (d *Dao) CreateUserGroup(userGroup *model.UserGroup) error {
	return d.engine.Model(&userGroup).Create(&userGroup).Error
}
