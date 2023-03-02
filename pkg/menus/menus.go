package menus

type Menus struct {
	Title     string
	Router    string
	Sort      uint8
	Role      []uint8
	ChildNode *Menus
}

func GetMenus() []*Menus {
	var MenusTree = []*Menus{
		&Menus{
			Title:     "首页",
			Router:    "/home",
			Sort:      0,
			Role:      []uint8{1, 2},
			ChildNode: nil,
		},
		&Menus{
			Title:     "用户",
			Router:    "/user",
			Sort:      1,
			Role:      []uint8{1, 2},
			ChildNode: nil,
		},
		&Menus{
			Title:     "订单",
			Router:    "/order",
			Sort:      0,
			Role:      []uint8{1, 2},
			ChildNode: nil,
		},
	}
	return MenusTree
}
