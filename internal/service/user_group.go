package service

import (
	"oms/global"
	"oms/internal/model"
	"oms/internal/request"
	"oms/internal/response"
	"oms/pkg/app"
	"oms/pkg/errcode"

	"github.com/jinzhu/gorm"
)

// 根据ID 获取用户组详情
func (s *Service) GetUserGroupById(id uint32) (*model.UserGroup, error) {
	return s.dao.GetUserGroupById(id)
}

// 获取用户总数量
func (s *Service) GetUserGroupCountList(param *request.GetListUserGroupRequest) (int, error) {
	userGroup := model.NewUserGroup()
	userGroup.Title = param.Title
	userGroup.State = param.State

	user := model.NewUser()
	user.Username = param.Useranme
	return s.dao.GetUserGroupListCount(userGroup, user)
}

// 用户组分页列表
func (s *Service) GetUserGroupListPager(param *request.GetListUserGroupRequest, pager *app.Pager) ([]*response.UserGroupResponse, error) {
	userGroup := model.NewUserGroup()
	userGroup.Title = param.Title
	userGroup.State = param.State

	user := model.NewUser()
	user.Username = param.Useranme

	pageOffset := app.GetPageOffset(pager.Page, pager.PageSize)
	return s.dao.GetUserGroupListPage(userGroup, user, pageOffset, pager.PageSize)
}

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
		err = s.dao.UpdateUserOnLeader(user.ID, user.GroupID)
		if err != nil {
			return errcode.ErrorUpdateUserGroupLeaderFail
		}
		return nil
	})
}

// 编辑用户组
func (s *Service) UpdateUserGroup(param *request.UpdateUserGroupPostRequest) error {
	return global.DBEngine.Transaction(func(tx *gorm.DB) error {
		s.SetDao(tx)
		newUserGroup := model.NewUserGroup()
		newUserGroup.Title = param.Title
		newUserGroup.State = param.State
		newUserGroup.Leader = param.Leader

		// 获取详情
		uGroup, _ := s.dao.GetUserGroupById(param.ID)
		if uGroup != nil && uGroup.ID != param.ID {
			return errcode.ErrorUserGroupExistsFail
		}

		if uGroup.Title != param.Title {
			// 检查是否有同名用户组
			uGroup, _ := s.dao.GetUserGroupByTitle(param.Title)
			if uGroup != nil && uGroup.ID != param.ID {
				return errcode.ErrorUserGroupExistsFail
			}
		}

		// 清空之前的用户组员
		userList, err := s.dao.GetUserListByGroupId(uGroup.ID)
		if err != nil {
			return errcode.ErrorGetUserListFail
		}
		clearUserIds := []int{}

		for _, v := range userList {
			clearUserIds = append(clearUserIds, int(v.ID))
		}
		err = s.dao.BatchUpdateUserOnGroupID(clearUserIds, 0)
		if err != nil {
			return errcode.ErrorUpdateUserGroupIDFail
		}

		// 组长信息变更
		if param.Leader != uGroup.Leader {
			// 清空旧组长信息
			userLeader, err := s.dao.GetUserLeaderByGroupLeader(uGroup.ID)
			if err != nil {
				return errcode.ErrorGetUserGroupLeaderFail
			}
			err = s.dao.UpdateUserOnLeader(userLeader.ID, 0)
			if err != nil {
				return errcode.ErrorUpdateUserGroupLeaderFail
			}

			// 查找新组长数据
			user, err := s.GetUserInfoById(uint32(param.Leader))
			if err != nil {
				return errcode.ErrorUserNotFound
			}

			user.GroupID = param.ID
			user.GroupLeader = param.Leader
			// 更新新组长数据
			err = s.dao.UpdateUserOnLeader(user.ID, user.GroupID)
			if err != nil {
				return errcode.ErrorUpdateUserGroupLeaderFail
			}
		}

		// 更新新的用户 group_id 字段数据
		err = s.dao.BatchUpdateUserOnGroupID(param.UserIds, uGroup.ID)
		if err != nil {
			return errcode.ErrorUpdateUserGroupIDFail
		}

		// 更新用户组
		err = s.dao.UpdateUserGroup(uGroup, newUserGroup)
		if err != nil {
			return errcode.ErrorCreateUserGroupFail
		}

		return nil
	})
}

func (s *Service) DeleteUserGroup(param *request.DeleteUserGroupRequest) error {
	// 检查是否有该用户
	userGroup, _ := s.dao.GetUserGroupById(param.ID)
	if userGroup == nil {
		return errcode.ErrorUserGroupNotFoundFail
	}

	return s.dao.DeleteUserGroup(userGroup)
}
