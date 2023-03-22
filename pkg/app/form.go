/*
 * @Author: GG
 * @Date: 2023-01-28 21:07:58
 * @LastEditTime: 2023-03-21 17:33:05
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

// 2个struct的同名字段赋值
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

// 获取二维数组类型的post数据
// p := app.GetPostMapForm(data, "product")
// fmt.Printf("p: %v\n", p)
// for k, v := range p {
// 	Product := &request.CreateOrderProductRequest{}
// 	app.GetPostMapFormItem(data, v["i"], v["j"], k, Product)
//
// }
// for k, v := range p {
// 	fmt.Printf("k: %v\n", k)
// 	fmt.Printf("v: %v\n", v)
// }
//
// k: 3
// v: map[i:7 j:13]

func GetPostMapForm(formData map[string][]string, fieldName string) map[string]map[string]int {
	var key string
	var i, j int
	dicts := make(map[string]map[string]int)
	for k, _ := range formData {
		if i = strings.IndexByte(k, '['); i >= 1 && k[0:i] == fieldName {
			if j = strings.IndexByte(k[i+1:], ']'); j >= 1 {
				key = k[i+1:][:j]
				if dicts[key] == nil {
					dicts[key] = make(map[string]int)
					dicts[key]["i"] = i
					dicts[key]["j"] = j
				}
			}
		}
	}

	return dicts
}

func GetPostMapFormItem(formData map[string][]string, pi, pj int, parentName, fieldName string, entity interface{}) {
	var key string
	var i, j, i2, j2 int
	for k, v := range formData {
		if len(k) >= pi && k[0:pi] == parentName {
			if pj > strings.IndexByte(k[pi+1:], ']') {
				continue
			}
			if i = strings.IndexByte(k[pi:], '['); i >= 1 {
				continue
			}
			if j = strings.IndexByte(k[pi:], ']'); j < 1 {
				continue
			}
			if k[pi:][i+1:j] != fieldName {
				continue
			}
			if i2 = strings.IndexByte(k[pi:][j+1:], '['); i2 < 0 {
				continue
			}

			if j2 = strings.IndexByte(k[pi:][j+1:], ']'); j2 < 1 {
				continue
			}

			key = k[pi:][j+1:][i2+1 : j2]

			refEntityV := reflect.ValueOf(entity).Elem()
			if refEntityV.FieldByName(key).IsValid() {
				refEntityV.FieldByName(key).Set(reflect.ValueOf(v[0]))
				continue
			}
			for i := 0; i < refEntityV.NumField(); i++ {
				field := refEntityV.Type().Field(i)
				if field.Tag.Get("form") == key {
					refEntityV.Field(i).Set(reflect.ValueOf(v[0]))
					break
				}
			}

		}
	}
}
