/*
 * @Author: GG
 * @Date: 2023-01-30 15:28:37
 * @LastEditTime: 2023-03-28 15:18:03
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\pkg\errcode\module_code.go
 *
 */
package errcode

var (
	ErrorLoginFail         = NewError(20000001, "登录失败")
	ErrorLoginPasswordFail = NewError(20000002, "用户或密码错误")

	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(20010002, "新增标签失败")
	ErrorUpdateTagFail  = NewError(20010003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20010004, "删除标签失败")
	ErrorCountTagFail   = NewError(20010005, "统计标签失败")

	ErrorUploadFileFail = NewError(20100001, "上传文件失败")

	ErrorGetUserListFail           = NewError(20020001, "获取用户列表失败")
	ErrorCreateUserFail            = NewError(20020002, "新增用户失败")
	ErrorUpdateUserFail            = NewError(20020003, "更新用户失败")
	ErrorDeleteUserFail            = NewError(20020004, "删除用户失败")
	ErrorCountUserFail             = NewError(20020005, "统计用户失败")
	ErrorUserNotFound              = NewError(20020006, "用户不存在")
	ErrorUserExists                = NewError(20020007, "用户名已存在")
	ErrorUpdateUserGroupIDFail     = NewError(20020008, "设置用户组成员失败")
	ErrorUpdateUserGroupLeaderFail = NewError(20020009, "设置用户组组长失败")
	ErrorGetUserGroupUserListFail  = NewError(20020010, "获取用户组成员失败")
	ErrorGetUserGroupLeaderFail    = NewError(20020011, "获取用户组组长失败")

	ErrorGetUserGroupListFail  = NewError(20030001, "获取用户组列表失败")
	ErrorCreateUserGroupFail   = NewError(20030002, "新增用户组失败")
	ErrorUpdateUserGroupFail   = NewError(20030003, "更新用户组失败")
	ErrorDeleteUserGroupFail   = NewError(20030004, "删除用户组失败")
	ErrorCountUserGroupFail    = NewError(20030005, "统计用户组失败")
	ErrorUserGroupNotFoundFail = NewError(20030006, "用户组不存在")
	ErrorUserGroupExistsFail   = NewError(20030007, "用户组已存在")

	ErrorGetOrderListFail         = NewError(20040001, "获取订单列表失败")
	ErrorCreateOrderFail          = NewError(20040002, "新增订单失败")
	ErrorUpdateOrderFail          = NewError(20040003, "更新订单失败")
	ErrorDeleteOrderFail          = NewError(20040004, "删除订单失败")
	ErrorCountOrderFail           = NewError(20040005, "统计订单失败")
	ErrorOrderNotFoundFail        = NewError(20040006, "订单不存在")
	ErrorOrderExistsFail          = NewError(20040007, "订单已存在")
	ErrorAddOrderProductFail      = NewError(20040008, "订单商品添加失败")
	ErrorOrderProductNotEmptyFail = NewError(20040009, "订单商品不能为空")

	ErrorGetMenusListFail  = NewError(20050001, "获取菜单列表失败")
	ErrorCreateMenusFail   = NewError(20050002, "创建菜单失败")
	ErrorUpdateMenusFail   = NewError(20050003, "更新菜单失败")
	ErrorDeleteMenusFail   = NewError(20050004, "删除菜单失败")
	ErrorMenusNotFoundFail = NewError(20050005, "菜单不存在")
)
