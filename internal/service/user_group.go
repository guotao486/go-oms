package service

import (
	"oms/global"
	"oms/internal/model"
	"oms/internal/request"
	"oms/pkg/errcode"

	"github.com/jinzhu/gorm"
)

// 创建用户组并更新用户数据
func (s *Service) CreateUserGroup(param *request.CreateUserGroupRequest) error {
	return global.DBEngine.Transaction(func(tx *gorm.DB) error {
		s.SetDao(tx)
		userGroup := model.NewUserGroup()
		userGroup.Title = param.Title
		userGroup.State = param.State
		userGroup.Leader = param.Leader
		// 检查是否有同名用户组
		_, err := s.dao.GetUserGroupByTitle(param.Title)
		if err == nil {
			return errcode.ErrorUserGroupExistsFail
		}
		// 创建用户组
		err = s.dao.CreateUserGroup(userGroup)
		if err != nil {
			return errcode.ErrorCreateUserGroupFail
		}
		// 更新用户 group_id 字段数据
		err = s.dao.BatchUpdateUserOnGroupID(param.UserIds, userGroup.ID)
		if err != nil {
			return errcode.ErrorUpdateUserGroupIDFail
		}

		// 查找组长数据
		user, err := s.GetUserInfoById(uint32(param.Leader))
		if err != nil {
			return errcode.ErrorUserNotFound
		}

		user.GroupID = userGroup.ID
		user.GroupLeader = param.Leader
		// 更新组长数据
		err = s.dao.UpdateUserOnLeader(user)
		if err != nil {
			return errcode.ErrorUpdateUserGroupLeaderFail
		}
		return nil
	})
}
