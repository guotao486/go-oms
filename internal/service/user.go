package service

import (
	"oms/internal/model"
	"oms/internal/request"
	"oms/internal/response"
	"oms/pkg/app"
	"oms/pkg/errcode"
)

// 获取用户分页列表
func (s *Service) GetUserListPager(param *request.GetListUserRequest, pager *app.Pager) ([]*response.UserResponse, error) {
	user := model.NewUser()
	user.Username = param.Username
	user.State = param.State
	pageOffset := app.GetPageOffset(pager.Page, pager.PageSize)

	var users []*response.UserResponse
	users, err := s.dao.GetUserListPages(user, pageOffset, pager.PageSize)
	// for _, u := range users {
	// 	u.CustomMarshal()
	// }
	return users, err
}

// 获取用户总数量
func (s *Service) GetUserCountList(param *request.GetListUserRequest) (int, error) {
	user := model.NewUser()
	user.Username = param.Username
	user.State = param.State
	return s.dao.GetUserListCount(user)
}

// 获取全部用户
func (s *Service) GetUserListAll() ([]*response.UserResponse, error) {
	user := model.NewUser()
	return s.dao.GetUserListAll(user)
}

// 根据id获取用户详情
func (s *Service) GetUserInfoById(id uint32) (*model.User, error) {
	user, err := s.dao.GetFirstById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 根据username获取用户详情
func (s *Service) GetUserInfoByUsername(Username string) (*model.User, error) {
	user, err := s.dao.GetFirstByUsername(Username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 创建用户
func (s *Service) CreateUser(param *request.CreateUserRequest) error {
	user := model.NewDefaultUser()
	user.Username = param.Username
	user.Password = param.Password
	user.SetSalt()
	err := user.SetPassword()
	if err != nil {
		return err
	}
	user.Level = param.Level
	user.Model = &model.Model{}
	return s.dao.CreateUser(user)
}

// 修改用户
func (s *Service) UpdateUser(param *request.UpdateUserPostRequest) error {
	// 检查是否有该用户
	user, _ := s.dao.GetFirstById(param.ID)
	if user == nil {
		return errcode.ErrorUserNotFound
	}
	// 检查新用户名是否存在
	uu, _ := s.dao.GetFirstByUsername(param.Username)
	if uu != nil && uu.ID != param.ID {
		return errcode.ErrorUserExists
	}
	newUser := model.NewUser()
	newUser.Level = param.Level
	newUser.State = param.State
	newUser.Username = param.Username

	return s.dao.UpdateUser(user, newUser)
}

// 删除用户
func (s *Service) DeleteUser(param *request.DeleteUserRequest) error {
	// 检查是否有该用户
	user, _ := s.dao.GetFirstById(param.ID)
	if user == nil {
		return errcode.ErrorUserNotFound
	}

	return s.dao.DeleteUser(user)
}
