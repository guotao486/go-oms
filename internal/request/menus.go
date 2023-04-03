/*
 * @Author: GG
 * @Date: 2023-03-30 14:57:40
 * @LastEditTime: 2023-03-31 16:46:16
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\request\menus.go
 *
 */
package request

type CreateMenusRequest struct {
	Title    string   `form:"title" binding:"required" label:"订单编号"`
	Router   string   `form:"router" binding:"required" label:"url"`
	Sort     uint8    `form:"sort,default=0" label:"排序"`
	Roles    []string `form:"role" binding:"required" label:"权限"`
	ParentID uint32   `form:"parent_id,default=0" label:"父级菜单"`
	State    uint8    `form:"state,default=1" binding:"required" label:"状态"`
}

type UpdateMenusGetRequest struct {
	ID uint32 `form:"id" binding:"required" label:"id"`
}
type UpdateMenusPostRequest struct {
	ID       uint32   `form:"id" binding:"required" label:"id"`
	Title    string   `form:"title" binding:"required" label:"订单编号"`
	Router   string   `form:"router" binding:"required" label:"url"`
	Sort     uint8    `form:"sort,default=0" label:"排序"`
	Roles    []string `form:"role" binding:"required" label:"权限"`
	ParentID uint32   `form:"parent_id,default=0" label:"父级菜单"`
	State    uint8    `form:"state,default=1" binding:"required" label:"状态"`
}

type DeleteMenusRequest struct {
	ID uint32 `form:"id" binding:"required" label:"id"`
}

type UpdateMenusSortRequest struct {
	ID   uint32 `form:"id" binding:"required" label:"id"`
	Sort uint8  `form:"sort" label:"排序"`
}
