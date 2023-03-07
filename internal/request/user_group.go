package request

type CreateUserGroupRequest struct {
	Title   string `form:"title" binding:"required,min=2,max=20" label:"用户组名称"`
	State   uint8  `form:"state" binding:"required,oneof=0 1" label:"状态"`
	Leader  uint8  `form:"leader" binding:"required" label:"组长"`
	UserIds []int  `form:"userIds" binding:"required" label:"成员"`
}

type UpdateUserGroupRequest struct{}

type DeleteUserGroupRequest struct{}
