/*
 * @Author: GG
 * @Date: 2023-02-28 11:41:40
 * @LastEditTime: 2023-03-08 12:58:16
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\dao\user.go
 *
 */
package dao

import (
	"oms/internal/model"
	"oms/internal/response"
	"oms/pkg/enum"
)

// 获取用户数量
func (d *Dao) GetUserListCount(user *model.User) (int, error) {
	var count int
	db := d.engine.Model(&user)
	if user.Username != "" {
		db = db.Where("username = ?", user.Username)
	}

	db = db.Where("state = ?", user.State)
	if err := db.Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// 获取用户分页列表
func (d *Dao) GetUserListPages(user *model.User, pageOffest, pageSize int) ([]*response.UserResponse, error) {
	var users []*response.UserResponse
	db := d.engine.Table(user.TableName())
	db = db.Select("id,username,level,state,group_id,group_leader,created_on")
	// 分页查询
	if pageOffest >= 0 && pageSize > 0 {
		db = db.Offset(pageOffest).Limit(pageSize)
	}

	db = db.Where("state = ?", user.State)

	if user.Username != "" {
		db = db.Where("username like ?", "%"+user.Username+"%")
	}
	if err := db.Where("is_del = ?", enum.IS_DEL_UNABLE).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// 获取所有用户列表
func (d *Dao) GetUserListAll(user *model.User) ([]*response.UserResponse, error) {
	var users []*response.UserResponse
	db := d.engine.Table(user.TableName())
	db = db.Select("id,username,level,state,group_id,group_leader,created_on")
	db = db.Where("state = ?", enum.DEFAULT_STATE)
	if err := db.Where("is_del = ?", enum.IS_DEL_UNABLE).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// 根据ID查找用户
func (d *Dao) GetFirstById(id uint32) (*model.User, error) {
	user := model.NewUser()
	err := d.engine.Table(user.TableName()).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

// 根据username查找用户
func (d *Dao) GetFirstByUsername(Username string) (*model.User, error) {
	user := model.NewUser()
	err := d.engine.Table(user.TableName()).Where("username = ?", Username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

// 创建用户
func (d *Dao) CreateUser(user *model.User) error {
	return d.engine.Model(&user).Create(&user).Error
}

// 修改用户
func (d *Dao) UpdateUser(user *model.User, value interface{}) error {
	return d.engine.Model(&user).Where("id = ? and is_del = ?", user.ID, enum.IS_DEL_UNABLE).Updates(value).Error
}

// 删除用户，有delete_on 和 is_del 字段则是软删除
func (d *Dao) DeleteUser(user *model.User) error {
	return d.engine.Model(&user).Where("id = ? and is_del = ?", user.ID, enum.IS_DEL_UNABLE).Delete(&user).Error
}

// 更新用户组字段
func (d *Dao) BatchUpdateUserOnGroupID(userIds []int, groupId uint32) error {
	user := model.NewUser()
	return d.engine.Model(&user).Where("id IN (?)", userIds).Updates(map[string]interface{}{"group_id": groupId}).Error
}

// 更新用户组长字段
func (d *Dao) UpdateUserOnLeader(id uint32, groupId uint32) error {
	user := model.NewUser()
	return d.engine.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{"group_id": groupId, "group_leader": groupId}).Error
}

// 根据用户组ID返回用户列表
func (d *Dao) GetUserListByGroupId(groupID uint32) ([]*model.User, error) {
	var userList []*model.User
	err := d.engine.Model(&model.User{}).Select("id,username").Where("group_id = ?", groupID).Find(&userList).Error
	return userList, err
}

// 根据用户组ID 返回组长信息
func (d *Dao) GetUserLeaderByGroupLeader(groupId uint32) (*model.User, error) {
	user := model.NewUser()
	err := d.engine.Model(&user).Where("group_leader = ?", groupId).First(&user).Error
	return user, err
}
