package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

// NewError
//
//	@Description: 实例化一个错误对象
//	@params code
//	@params msg
//	@return *Error
func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

// Error
//
//	@Description: 返回错误响应信息
//	@receiver e
//	@return string
func (e *Error) Error() string {
	return fmt.Sprintf("错误码 %d ,错误信息：%s", e.Code(), e.Msg())
}

// Code
//
//	@Description: 返回错误码
//	@receiver e
//	@return int
func (e *Error) Code() int {
	return e.code
}

// Msg
//
//	@Description: 返回错误信息
//	@receiver e
//	@return string
func (e *Error) Msg() string {
	return e.msg
}

// Msgf
//
//	@Description: 返回格式化的错误信息
//	@receiver e
//	@params args
//	@return string
func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

// Details
//
//	@Description: 返回响应详情
//	@receiver e
//	@return []string
func (e *Error) Details() []string {
	return e.details
}

// WithDetails
//
//	@Description: 多条结果的响应详情
//	@receiver e
//	@params details
//	@return *Error
func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}

	return &newError
}

// StatusCode
//
//	@Description: 将内部错误码转换成http
//	@receiver e
//	@return int
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case NotFound.Code():
		return http.StatusNotFound
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
