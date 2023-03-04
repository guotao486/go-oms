/*
 * @Author: GG
 * @Date: 2023-02-28 10:54:16
 * @LastEditTime: 2023-02-28 11:03:07
 * @LastEditors: GG
 * @Description: enum
 * @FilePath: \oms\pkg\enum\common.go
 *
 */
package enum

var (
	UNABLE uint8 = 0 // 停止
	ENABLE uint8 = 1 // 开启

	IS_DEL_UNABLE uint8 = UNABLE // 未删除
	IS_DEL_ENABLE uint8 = ENABLE // 已删除

	DEFAULT        uint8 = ENABLE // uint 默认值
	DEFAULT_STATE  uint8 = ENABLE // 默认开启
	DEFAULT_IS_DEL uint8 = UNABLE // 默认没有删除
)

var (
	// value = 1
	USER_LEVEL_ADMIN      uint8  = 1
	USER_LEVEL_ADMIN_TEXT string = "管理员"
	// value = 2
	USER_LEVEL_STAFF      uint8  = 2
	USER_LEVEL_STAFF_TEXT string = "普通用户"
	// value = 2
	DEFAULT_USER_LEVEL = USER_LEVEL_STAFF
)
