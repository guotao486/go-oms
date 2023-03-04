/*
 * @Author: GG
 * @Date: 2023-01-28 11:04:25
 * @LastEditTime: 2023-03-02 16:13:26
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\pkg\app\pagination.go
 *
 */
package app

import (
	"oms/global"
	"oms/pkg/convert"

	"github.com/gin-gonic/gin"
)

/**
* @Author $
* @Description //TODO $
* @Date $ $
* @Param $
* @return $
**/

// GetPage
//
//	@Description: 返回页码参数
//	@params c 上下文
//	@return int
func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

// GetPageSize
//
//	@Description: 返回每页数量
//	@params c
//	@return int
func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("limit")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}

	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}

	return pageSize
}

// GetPageOffset
//
//	@Description: 分页查询偏移量
//	@params page
//	@params pageSize
//	@return int
func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
