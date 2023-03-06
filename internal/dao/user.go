/*
 * @Author: GG
 * @Date: 2023-02-28 11:41:40
 * @LastEditTime: 2023-03-06 11:01:02
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
