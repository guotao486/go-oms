package errcode

var (
	Success                   = NewError(0, "完成")
	ServerError               = NewError(10000000, "服务器内部错误")
	InvalidParams             = NewError(10000001, "入参错误")
	NotFound                  = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout  = NewError(10000005, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败，Token生成失败")
	TooManyRequests           = NewError(10000007, "请求过多") // 请求过多
	NotLogin                  = NewError(10000008, "请登录")
	NotRoleAuth               = NewError(10000009, "没有角色权限")
)
