/*
 * @Author: GG
 * @Date: 2023-02-28 17:10:15
 * @LastEditTime: 2023-03-06 14:32:45
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\pkg\menus\menus.go
 *
 */
package menus

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
					Role:      []uint8{1, 2},
					ChildNode: nil,
				},
				{
					Title:     "编辑用户",
					Router:    "/user/update",
					Sort:      1,
					Role:      []uint8{1, 2},
					ChildNode: nil,
				},
			},
		},
		{
			Title:     "用户组",
			Router:    "/group",
			Sort:      1,
			Role:      []uint8{1, 2},
			ChildNode: nil,
		},
		{
			Title:     "订单管理",
			Router:    "/order",
			Sort:      0,
			Role:      []uint8{1, 2},
			ChildNode: nil,
		},
	}
	return MenusTree
}

func GetCurrent(urlPath string) *Menus {
	m := GetMenus()
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
