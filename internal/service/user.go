package service

import (
	"oms/internal/model"
	"oms/internal/request"
)

// CreateUser
func (s *Service) CreateUser(param *request.CreateUserRequest) error {
	user := model.NewUser()
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
