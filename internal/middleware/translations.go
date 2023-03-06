/*
 * @Author: GG
 * @Date: 2023-01-28 20:49:23
 * @LastEditTime: 2023-01-28 21:07:23
 * @LastEditors: GG
 * @Description: 参数验证国际化处理中间件
 * @FilePath: \oms\internal\middleware\translations.go
 *
 */
package middleware

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

// 在识别当前请求的语言类别上，我们通过 GetHeader 方法去获取约定的 header 参数 locale，用于判别当前请求的语言类别是 en 又或是 zh，如果有其它语言环境要求，也可以继续引入其它语言类别，因为 go-playground/locales 基本上都支持。

// 在后续的注册步骤，我们调用 RegisterDefaultTranslations 方法将验证器和对应语言类型的 Translator 注册进来，实现验证器的多语言支持。

// 同时将 Translator 存储到全局上下文中，便于后续翻译时的使用。
/**
 * @description: 参数验证国际化处理中间件
 * @return {*}
 */
func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		// 从Header中获取locale判断语言类别
		locale := c.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			// 调用 RegisterDefaultTranslations 方法将验证器和对应语言类型的 Translator 注册进来，实现验证器的多语言支持
			switch locale {
			case "zh":
				tagLabel(v)
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				_ = en_translations.RegisterDefaultTranslations(v, trans)
				break
			default:
				tagLabel(v)
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			}
			// 将 Translator 存储到全局上下文
			c.Set("trans", trans)
		}
		// 执行下一步
		c.Next()
	}
}

func tagLabel(v *validator.Validate) {
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})
}
