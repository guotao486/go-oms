package service

import (
	"oms/internal/model"
	"oms/internal/request"
)

// 创建用户组并更新用户数据
func (s *Service) CreateUserGroup(param *request.CreateUserGroupRequest) error {
	userGroup := model.NewUserGroup()
	userGroup.Title = param.Title
	userGroup.State = param.State
	userGroup.Leader = param.Leader
	// 创建用户组
	err := s.dao.CreateUserGroup(userGroup)
	if err != nil {
		return err
	}
	// 更新用户 group_id 字段数据
	err = s.dao.BatchUpdateUserOnGroupID(param.UserIds, userGroup.ID)
	if err != nil {
		return err
	}

	// 查找组长数据
	user, err := s.GetUserInfoById(uint32(param.Leader))
	if err != nil {
		return err
	}

	user.GroupID = userGroup.ID
	user.GroupLeader = param.Leader
	// 更新组长数据
	err = s.dao.UpdateUserOnLeader(user)
	if err != nil {
		return err
	}
	return nil
}
