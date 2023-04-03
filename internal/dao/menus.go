package dao

import (
	"oms/internal/model"
	"oms/internal/response"
	"oms/pkg/enum"
)

// 获取所有菜单
func (d Dao) GetMenusListAll() ([]*response.MenusResponse, error) {
	var menus []*response.MenusResponse
	err := d.engine.Scopes(IsDelToUnable, StateToUnable).Order("sort desc").Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

// 获取一级菜单
func (d Dao) GetParentMenusList() ([]*response.MenusResponse, error) {
	var menus []*response.MenusResponse
	err := d.engine.Where("parent_id = ?", enum.UNABLE).Scopes(IsDelToUnable, StateToUnable).Order("sort desc").Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (d Dao) CreateMenus(menus *model.Menus) error {
	return d.engine.Create(&menus).Error
}

func (d Dao) UpdateMenus(menus *model.Menus) error {
	return d.engine.Save(&menus).Error
}

func (d Dao) GetMenusById(id uint32) (*model.Menus, error) {
	menus := model.NewMenus()
	err := d.engine.Where("id = ?", id).First(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (d *Dao) DeleteMenus(menus *model.Menus) error {
	return d.engine.Model(&menus).Scopes(IsDelToUnable).Delete(&menus).Error
}
