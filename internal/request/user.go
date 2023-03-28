package request

// 登录
type LoginRequest struct {
	Username string `form:"username" binding:"required,min=2" label:"用户名"`
	Password string `form:"password" binding:"required,min=6" label:"密码"`
}

// 创建
type CreateUserRequest struct {
	Username string `form:"username" binding:"required,min=2,max=25" label:"用户名"`
	Password string `form:"password" binding:"required,min=6,max=25" label:"密码"`
	Level    uint8  `form:"level" binding:"required"`
}

// 修改详情
type UpdateUserGetRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

// 修改提交
type UpdateUserPostRequest struct {
	ID       uint32 `form:"id" binding:"required,gte=1"`
	Username string `form:"username" binding:"required,min=2,max=25" label:"用户名"`
	Level    uint8  `form:"level" binding:"required" label:"用户权限"`
	State    uint8  `form:"state" binding:"required,oneof=1 2" label:"用户状态"`
}

// 删除
type DeleteUserRequest struct {
	ID uint32 `uri:"id" form:"id" json:"id" binding:"required,gte=1"`
}

// 用户搜索列表
type GetListUserRequest struct {
	Username string `form:"username" binding:"max=100"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
}
