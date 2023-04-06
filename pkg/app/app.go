/*
 * @Author: GG
 * @Date: 2023-01-28 11:04:27
 * @LastEditTime: 2023-03-02 16:33:04
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\pkg\app\app.go
 *
 */
package app

import (
	"net/http"
	"oms/pkg/errcode"

	"github.com/gin-gonic/gin"
)

/**
* @Author $
* @Description //TODO $
* @Date $ $
* @Param $
* @return $
**/

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

type Success struct {
	Code    uint8       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccess() *Success {
	return &Success{
		Code:    http.StatusOK,
		Message: "success",
	}
}

// ToResponse
//
//	@Description: 返回响应
//	@receiver r
//	@params data
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = NewSuccess()
	}
	r.Ctx.JSON(http.StatusOK, data)
}

// ToSuccessResponse
//
/**
 * @description: 返回success响应
 * @param {interface{}} data
 * @return {*}
 */
func (r *Response) ToSuccessResponse(data interface{}) {
	success := NewSuccess()
	if data != nil {
		success.Data = data
	}
	r.ToResponse(success)
}

// ToResponseList
//
//	@Description: 返回响应列表
//	@receiver r
//	@params list
//	@params totalRows
func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.ToSuccessResponse(gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

// ToErrorResponse
//
//	@Description: 返回错误响应
//	@receiver r
//	@params err
func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{
		"code":    err.Code(),
		"message": err.Msg(),
	}
	details := err.Details()

	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(err.StatusCode(), response)
}
