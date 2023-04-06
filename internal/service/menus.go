package service

import (
	"oms/global"
	"oms/internal/model"
	"oms/internal/request"
	"oms/internal/response"
	"oms/pkg/app"
	"oms/pkg/convert"
	"oms/pkg/errcode"
	"strings"
)

// 根据id获取数据
func (s Service) GetMenusById(id uint32) (*model.Menus, error) {
	return s.dao.GetMenusById(id)
}

func (s Service) GetMenusListAll() ([]*response.MenusResponse, error) {
	cache := global.CacheStore.Engine
	var menusList []*response.MenusResponse
	// 获取缓存
	cacheList, _ := cache.Get(model.CacheMenusListKey)
	// 若缓存不存在
	if cacheList == nil {
		menusList, err := s.dao.GetMenusListAll()
		if err != nil {
			return nil, err
		}
		st := &convert.StructTo{V: menusList}
		buf, err := st.StructToBytes()
		if err != nil {
			return menusList, err
		}
		err = cache.Set(model.CacheMenusListKey, buf)
		return menusList, err
	}

	// 缓存存在
	bs := &convert.ByteTo{V: &menusList}
	err := bs.ByteToStruct(cacheList)
	if err != nil {
		return nil, err
	}
	return menusList, nil
}

func (s Service) GetParentMenusList() ([]*response.MenusResponse, error) {
	cache := global.CacheStore.Engine
	var menusList []*response.MenusResponse
	// 获取缓存
	cacheList, _ := cache.Get(model.CacheParentMenusListKey)
	// 若缓存不存在
	if cacheList == nil {
		menusList, err := s.dao.GetParentMenusList()
		if err != nil {
			return nil, err
		}
		st := &convert.StructTo{V: menusList}
		buf, err := st.StructToBytes()
		if err != nil {
			return menusList, err
		}
		err = cache.Set(model.CacheParentMenusListKey, buf)
		return menusList, err
	}

	// 缓存存在
	bs := &convert.ByteTo{V: &menusList}
	err := bs.ByteToStruct(cacheList)
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
