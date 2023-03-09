/*
 * @Author: GG
 * @Date: 2023-03-06 15:22:05
 * @LastEditTime: 2023-03-09 16:41:40
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\dao\user_group.go
 *
 */
package dao

import (
	"oms/internal/model"
	"oms/internal/response"
	"oms/pkg/enum"
)

// 创建用户组
func (d *Dao) CreateUserGroup(userGroup *model.UserGroup) error {
	return d.engine.Model(&userGroup).Create(&userGroup).Error
}

// 修改用户组
func (d *Dao) UpdateUserGroup(userGroup *model.UserGroup, value interface{}) error {
	return d.engine.Model(&userGroup).Where("id = ? and is_del = ?", userGroup.ID, enum.IS_DEL_UNABLE).Updates(value).Error
}

// 删除用户组
func (d *Dao) DeleteUserGroup(userGroup *model.UserGroup) error {
	return d.engine.Model(&userGroup).Where("id = ? and is_del = ?", userGroup.ID, enum.IS_DEL_UNABLE).Delete(&userGroup).Error
}

// 根据ID 查找用户组
func (d *Dao) GetUserGroupById(id uint32) (*model.UserGroup, error) {
	userGroup := model.NewUserGroup()
	err := d.engine.Table(userGroup.TableName()).Where("id = ? and is_del = ? ", id, enum.IS_DEL_UNABLE).First(&userGroup).Error
	if err != nil {
		return nil, err
	}
	return userGroup, nil
}

// 根据名称查找用户组
func (d *Dao) GetUserGroupByTitle(title string) (*model.UserGroup, error) {
	userGroup := model.NewUserGroup()
	userGroup.Title = title
	err := d.engine.Model(&userGroup).Where("title = ? and is_del = ? ", userGroup.Title, enum.IS_DEL_UNABLE).First(&userGroup).Error
	if err != nil {
		return nil, err
	}
	return userGroup, nil
}

// 总数量
func (d *Dao) GetUserGroupListCount(userGroup *model.UserGroup, user *model.User) (int, error) {
	var count int
	db := d.engine.Table(userGroup.TableName())

	if user.Username != "" {
		if user, err := d.GetFirstByUsername(user.Username); err != nil {
			return 0, err
		} else {
			db = db.Where("id = ?", user.GroupID)
		}
	}
	if err := db.Where("is_del = ?", enum.DEFAULT_IS_DEL).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// 用户组分页列表
func (d *Dao) GetUserGroupListPage(userGroup *model.UserGroup, user *model.User, pageOffset, pageSize int) ([]*response.UserGroupResponse, error) {
	var userGroupList []*response.UserGroupResponse
	db := d.engine.Model(&userGroup)
	db = db.Select("id,title,state,leader,created_on")
	// 分页查询
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if user.Username != "" {
		if user, err := d.GetFirstByUsername(user.Username); err != nil {
			return nil, err
		} else {
			db = db.Where("id = ?", user.GroupID)
		}
	}
	if err := db.Where("is_del = ?", enum.DEFAULT_IS_DEL).Preload("UserList").Find(&userGroupList).Error; err != nil {
		return nil, err
	}
	return userGroupList, nil
}
