/*
 * @Author: GG
 * @Date: 2023-01-30 15:28:37
 * @LastEditTime: 2023-01-30 15:33:59
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\pkg\errcode\module_code.go
 *
 */
package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(20010002, "新增标签失败")
	ErrorUpdateTagFail  = NewError(20010003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20010004, "删除标签失败")
	ErrorCountTagFail   = NewError(20010005, "统计标签失败")

	ErrorGetUserListFail = NewError(20020001, "获取用户列表失败")
	ErrorCreateUserFail  = NewError(20020002, "新增用户失败")
	ErrorUpdateUserFail  = NewError(20020003, "更新用户失败")
	ErrorDeleteUserFail  = NewError(20020004, "删除用户失败")
	ErrorCountUserFail   = NewError(20020005, "统计用户失败")

	ErrorUploadFileFail = NewError(20030001, "上传文件失败")
)
