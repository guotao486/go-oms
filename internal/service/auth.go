package service

import (
	"oms/internal/model"
	"oms/internal/request"
	"oms/pkg/errcode"
	"oms/pkg/util"
)

// CheckAuth
//
/**
 * @description: 检查auth信息
 * @param {*request.AuthRequest} param
 * @return {*}
 */
func (svc *Service) CheckAuth(param *request.AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return nil
	}

	// errors.New('auth info does not exist.')
	return errcode.NotFound
}

func (svc *Service) Login(param *request.LoginRequest) (*model.User, error) {
	user, err := svc.GetUserInfoByUsername(param.Username)
	if err != nil {
		return nil, errcode.ErrorLoginPasswordFail
	}

	password := param.Password + user.Salt

	if !util.CheckPasswordHash(password, user.Password) {
		return nil, errcode.ErrorLoginPasswordFail
	}

	return user, nil
}
