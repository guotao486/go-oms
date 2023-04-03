package service

import (
	"oms/internal/model"
	"oms/internal/request"
	"oms/internal/response"
	"oms/pkg/app"
	"oms/pkg/errcode"
	"strings"
)

// 根据id获取数据
func (s Service) GetMenusById(id uint32) (*model.Menus, error) {
	return s.dao.GetMenusById(id)
}

func (s Service) GetMenusListAll() ([]*response.MenusResponse, error) {
	menusList, err := s.dao.GetMenusListAll()
	if err != nil {
		return nil, err
	}
	return menusList, nil
}

func (s Service) GetParentMenusList() ([]*response.MenusResponse, error) {
	menusList, err := s.dao.GetParentMenusList()
	if err != nil {
		return nil, err
	}
	return menusList, nil
}

func (s Service) CreateMenus(param *request.CreateMenusRequest) error {
	menus := model.NewMenus()
	app.StructAssign(menus, param)
	menus.Role = strings.Join(param.Roles, ",")
	return s.dao.CreateMenus(menus)
}

func (s Service) UpdateMenus(param *request.UpdateMenusPostRequest) error {
	menus, err := s.GetMenusById(param.ID)
	if err != nil {
		return errcode.ErrorMenusNotFoundFail
	}

	app.StructAssign(menus, param)
	menus.Role = strings.Join(param.Roles, ",")
	return s.dao.UpdateMenus(menus)
}

func (s Service) UpdateMenusSort(param *request.UpdateMenusSortRequest) error {
	menus, err := s.GetMenusById(param.ID)
	if err != nil {
		return errcode.ErrorMenusNotFoundFail
	}

	app.StructAssign(menus, param)
	return s.dao.UpdateMenus(menus)
}

func (s Service) DeleteMenus(param *request.DeleteMenusRequest) error {

	menus, err := s.GetMenusById(param.ID)
	if err != nil {
		return errcode.ErrorMenusNotFoundFail
	}

	return s.dao.DeleteMenus(menus)
}
