/*
 * @Author: GG
 * @Date: 2023-03-06 15:22:05
 * @LastEditTime: 2023-03-07 12:05:56
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\dao\user_group.go
 *
 */
package dao

import (
	"oms/internal/model"
	"oms/pkg/enum"
)

// 创建用户组
func (d *Dao) CreateUserGroup(userGroup *model.UserGroup) error {
	return d.engine.Model(&userGroup).Create(&userGroup).Error
}

func (d *Dao) UpdateUserGroup(userGroup *model.UserGroup, value interface{}) error {
	return d.engine.Model(&userGroup).Where("id = ? and is_del = ?", userGroup.ID, enum.IS_DEL_UNABLE).Updates(value).Error
}

// 根据ID 查找用户组
func (d *Dao) GetUserGroupById(id uint32) (*model.UserGroup, error) {
	userGroup := model.NewUserGroup()
	err := d.engine.Table(userGroup.TableName()).Where("id = ? and is_del = ? ", id, enum.IS_DEL_UNABLE).First(&userGroup).Error
	return userGroup, err
}

// 根据名称查找用户组
func (d *Dao) GetUserGroupByTitle(title string) (*model.UserGroup, error) {
	userGroup := model.NewUserGroup()
	userGroup.Title = title
	err := d.engine.Model(&userGroup).Where("title = ? and is_del = ? ", userGroup.Title, enum.IS_DEL_UNABLE).First(&userGroup).Error
	return userGroup, err
}
