/*
 * @Author: GG
 * @Date: 2023-02-28 17:10:15
 * @LastEditTime: 2023-04-06 15:04:52
 * @LastEditors: GG
 * @Description: 菜单工具类
 * @FilePath: \oms\pkg\menus\menus.go
 *
 */
package menus

import (
	"oms/global"
	"oms/internal/response"
	"oms/internal/service"

	"github.com/gin-gonic/gin"
)

func GetMenus(c *gin.Context) []*response.MenusResponse {
	svc := service.New(c)
	list, err := svc.GetMenusListAll()
	if err != nil {
		global.Logger.Errorf(c, "svc.GetMenusListAll err: %v", err)
		return nil
	}
	return list
}

func GetMenusTree(c *gin.Context) []*response.MenusResponse {
	m := GetMenus(c)
	return getTreeRecursive(m, 0)
}
func getTreeRecursive(list []*response.MenusResponse, parentId uint32) []*response.MenusResponse {
	res := make([]*response.MenusResponse, 0)

	for _, v := range list {
		if v.ParentID == parentId {
			v.ChildNode = getTreeRecursive(list, v.ID)
			res = append(res, v)
		}
	}
	return res
}

func GetCurrent(c *gin.Context) *response.MenusResponse {
	m := GetMenusTree(c)
	urlPath := c.Request.URL.Path
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
