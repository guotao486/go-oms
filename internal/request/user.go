package request

// 登录
type LoginRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// 创建
type CreateUserRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"username" binding:"required"`
	Level    uint8  `form:"level" binding:"required"`
}

// 修改
type UpdateUserRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"username" binding:"required"`
	Level    uint8  `form:"level" binding:"required"`
	State    uint8  `form:"state" binding:"required"`
}
