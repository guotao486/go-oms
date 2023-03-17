/*
 * @Author: GG
 * @Date: 2023-01-28 21:07:58
 * @LastEditTime: 2023-03-17 15:16:59
 * @LastEditors: GG
 * @Description: 表单验证相关
 * @FilePath: \oms\pkg\app\form.go
 *
 */
package app

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
)

// ValidError
/**
 * @description: 验证错误对象
 * @return {*}
 */
type ValidError struct {
	Key     string
	Message string
}

// Error
/**
 * @description: 返回错误信息
 * @return {*}
 */
func (v *ValidError) Error() string {
	return v.Message
}

// 错误集合类型
type ValidErrors []*ValidError

/**
 * @description: 多个错误信息 string
 * @return {*}
 */
func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

/**
 * @description: 多个错误信息 []string
 * @return {*}
 */
func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

// BindAndValid
//
/**
 * @description: 绑定数据并开始校验
 * @param {*gin.Context} c
 * @param {interface{}} v
 * @return {*}
 */
func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(v) // 参数绑定
	if err != nil {
		v := c.Value("trans")                   // 取出中间件存储的Translator
		trans, _ := v.(ut.Translator)           // 类型断言，获取Translator
		verrs, ok := err.(val.ValidationErrors) // 获取参数校验错误信息
		if !ok {
			return false, errs
		}

		for key, value := range verrs.Translate(trans) { // 开始具体的翻译
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}
		return false, errs
	}
	return true, nil
}

func StructAssign(binding interface{}, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem() // 获取reflect.Type类型
	vVal := reflect.ValueOf(value).Elem()   // 获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		// 同类型
		valType := vTypeOfT.Field(i).Type

		if ok := bVal.FieldByName(name).IsValid() && bVal.FieldByName(name).Type() == valType; ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
}
