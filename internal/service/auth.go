package service

import (
	"oms/internal/request"
	"oms/pkg/errcode"
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
