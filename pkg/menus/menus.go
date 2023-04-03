/*
 * @Author: GG
 * @Date: 2023-02-28 17:10:15
 * @LastEditTime: 2023-04-03 16:48:43
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\pkg\menus\menus.go
 *
 */
package menus

import "fmt"

type Menus struct {
	Title     string
	Router    string
	Sort      uint8
	Role      []uint8
	ChildNode []*Menus
}

func GetMenus() []*Menus {
	var MenusTree = []*Menus{
		{
			Title:     "首页",
			Router:    "/home",
			Sort:      0,
			Role:      []uint8{1, 2},
			ChildNode: nil,
		},
		{
			Title:  "用户管理",
			Router: "/user/",
			Sort:   1,
			Role:   []uint8{1, 2},
			ChildNode: []*Menus{
				{
					Title:     "新增用户",
					Router:    "/user/create",
					Sort:      1,
					Role:      []uint8{1},
					ChildNode: nil,
				},
				{
					Title:     "编辑用户",
					Router:    "/user/update",
					Sort:      1,
					Role:      []uint8{1},
					ChildNode: nil,
				},
				{
					Title:     "删除用户",
					Router:    "/user/delete",
					Sort:      1,
					Role:      []uint8{1},
					ChildNode: nil,
				},
			},
		},
		{
			Title:  "用户组",
			Router: "/group/",
			Sort:   1,
			Role:   []uint8{1, 2},
			ChildNode: []*Menus{
				{
					Title:     "新增用户组",
					Router:    "/group/create",
					Sort:      1,
					Role:      []uint8{1},
					ChildNode: nil,
				},
				{
					Title:     "编辑用户组",
					Router:    "/group/update",
					Sort:      1,
					Role:      []uint8{1},
					ChildNode: nil,
				},
				{
					Title:     "删除用户组",
					Router:    "/group/delete",
					Sort:      1,
					Role:      []uint8{1},
					ChildNode: nil,
				},
			},
		},
		{
			Title:  "订单管理",
			Router: "/order/",
			Sort:   0,
			Role:   []uint8{1, 2},
			ChildNode: []*Menus{
				{
					Title:     "新增订单",
					Router:    "/order/create",
					Sort:      1,
					Role:      []uint8{1},
					ChildNode: nil,
				},
				{
					Title:     "编辑订单",
					Router:    "/order/update",
					Sort:      1,
					Role:      []uint8{1},
					ChildNode: nil,
				},
				{
					Title:     "删除订单",
					Router:    "/order/delete",
					Sort:      1,
					Role:      []uint8{1},
					ChildNode: nil,
				},
			},
		},
		{
			Title:  "菜单管理",
			Router: "/menus/",
			Sort:   1,
			Role:   []uint8{1, 2},
			ChildNode: []*Menus{
				{
					Title:     "新增菜单",
					Router:    "/menus/create",
					Sort:      1,
					Role:      []uint8{1},
					ChildNode: nil,
				},
				{
					Title:     "编辑菜单",
					Router:    "/menus/update",
					Sort:      1,
					Role:      []uint8{1},
					ChildNode: nil,
				},
				{
					Title:     "删除菜单",
					Router:    "/menus/delete",
					Sort:      1,
					Role:      []uint8{1},
					ChildNode: nil,
				},
			},
		},
	}
	return MenusTree
}

func GetCurrent(urlPath string) *Menus {
	m := GetMenus()
	fmt.Printf("urlPath: %v\n", urlPath)
	for _, menu := range m {
		if urlPath == menu.Router {
			return menu
		}

		for _, cMenu := range menu.ChildNode {
			if urlPath == cMenu.Router {
				return cMenu
			}
		}
	}
	return nil
}
