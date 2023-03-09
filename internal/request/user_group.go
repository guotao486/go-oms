/*
 * @Author: GG
 * @Date: 2023-03-06 15:26:15
 * @LastEditTime: 2023-03-09 14:21:18
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\request\user_group.go
 *
 */
package request

type CreateUserGroupRequest struct {
	Title   string `form:"title" binding:"required,min=2,max=20" label:"用户组名称"`
	State   uint8  `form:"state" binding:"required,oneof=0 1" label:"状态"`
	Leader  uint8  `form:"leader" binding:"required" label:"组长"`
	UserIds []int  `form:"userIds" binding:"required" label:"成员"`
}

type UpdateUserGroupGetRequest struct {
	ID uint32 `form:"id" binding:"required,number"`
}

type UpdateUserGroupPostRequest struct {
	ID      uint32 `form:"id" binding:"required,number"`
	Title   string `form:"title" binding:"required,min=2,max=20" label:"用户组名称"`
	State   uint8  `form:"state" binding:"required,oneof=0 1" label:"状态"`
	Leader  uint8  `form:"leader" binding:"required" label:"组长"`
	UserIds []int  `form:"userIds" binding:"required" label:"成员"`
}

type DeleteUserGroupRequest struct {
	ID uint32 `form:"id" form:"id" json:"id" binding:"required,gte=1"`
}

// 用户搜索列表
type GetListUserGroupRequest struct {
	Title    string `form:"title" binding:"max=100"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
	Useranme string `form:"username" binding:"max=100"`
}
