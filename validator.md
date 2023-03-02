<!--
 * @Author: GG
 * @Date: 2023-01-28 20:25:48
 * @LastEditTime: 2023-01-28 22:04:58
 * @LastEditors: GG
 * @Description: 
 * @FilePath: \oms\validator.md
 * 
-->
# validator 使用

## 安装
```
$ go get -u github.com/go-playground/validator/v10
```

## 校验规则

|标签|	含义|
|---|---|
|required|	必填|
|gt	|大于|
|gte|	大于等于|
|lt	|小于|
|lte	|小于等于|
|min	|最小值|
|max	|最大值|
|oneof|	参数集内的其中之一|
|len|	长度要求与 len 给定的一致|

## 国际化处理

* go-playground/locales：多语言包，是从 CLDR 项目（Unicode 通用语言环境数据存储库）生成的一组多语言环境，主要在 i18n 软件包中使用，该库是与 universal-translator 配套使用的。
* go-playground/universal-translator：通用翻译器，是一个使用 CLDR 数据 + 复数规则的 Go 语言 i18n 转换器。
* go-playground/validator/v10/translations：validator 的翻译器。
  
### 编写中间件
```
package middleware

import (
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
 * @description: 中间件，国际化处理
 * @return {*}
 */
func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		locale := c.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "zh":
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				_ = en_translations.RegisterDefaultTranslations(v, trans)
				break
			default:
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			}
			// 存入上下文
			c.Set("trans", trans)
		}
		// 执行下一步
		c.Next()
	}
}

```

### 验证
```
import (
    ...
    ut "github.com/go-playground/universal-translator"
    val "github.com/go-playground/validator/v10"
)
type ValidError struct {
    Key     string
    Message string
}
type ValidErrors []*ValidError
func (v *ValidError) Error() string {
    return v.Message
}
func (v ValidErrors) Error() string {
    return strings.Join(v.Errors(), ",")
}
func (v ValidErrors) Errors() []string {
    var errs []string
    for _, err := range v {
        errs = append(errs, err.Error())
    }
    return errs
}
func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
    var errs ValidErrors
    err := c.ShouldBind(v)
    if err != nil {
        v := c.Value("trans")
        trans, _ := v.(ut.Translator)
        verrs, ok := err.(val.ValidationErrors)
        if !ok {
            return false, errs
        }
        for key, value := range verrs.Translate(trans) {
            errs = append(errs, &ValidError{
                Key:     key,
                Message: value,
            })
        }
        return false, errs
    }
    return true, nil
}
```

### 使用
```
func (t Tag) List(c *gin.Context) {
    param := struct {
        Name  string `form:"name" binding:"max=100"`
        State uint8  `form:"state,default=1" binding:"oneof=0 1"`
    }{}
    response := app.NewResponse(c)
    valid, errs := app.BindAndValid(c, &param)
    if !valid {
        global.Logger.Errorf("app.BindAndValid errs: %v", errs)
        response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
        return
    }
    response.ToResponse(gin.H{})
    return
}
```

* required，有参数才验证，没有该入参的话，是默认无校验的

### 自定义错误信息
```
import (
    "gopkg.in/go-playground/validator.v8"
)

// 绑定模型
type LoginRequest struct {
    Mobile string   `form:"mobile" json:"mobile" binding:"required"`
    Code string `form:"code" json:"code" binding:"required"`
}

// 绑定模型获取验证错误的方法
func (r *LoginRequest) GetError (err validator.ValidationErrors) string {

    // 这里的 "LoginRequest.Mobile" 索引对应的是模型的名称和字段
    if val, exist := err["LoginRequest.Mobile"]; exist {
        if val.Field == "Mobile" {
            switch val.Tag{
                case "required":
                    return "请输入手机号码"
            }
        }
    }
    if val, exist := err["LoginRequest.Code"]; exist {
        if val.Field == "Code" {
            switch val.Tag{
                case "required":
                    return "请输入验证码"
            }
        }
    }
    return "参数错误"
}
```

### 自定义错误信息使用
```
import (
    "github.com/gin-gonic/gin"
    "net/http"
    "gopkg.in/go-playground/validator.v8"
)


func Login(c *gin.Context) {
    var loginRequest LoginRequest

    if err := c.ShouldBind(&loginRequest); err == nil { 
        // 参数接收正确, 进行登录操作

        c.JSON(http.StatusOK, loginRequest)
    }else{
        // 验证错误
        c.JSON(http.StatusUnprocessableEntity, gin.H{
            "message": loginRequest.GetError(err.(validator.ValidationErrors)), // 注意这里要将 err 进行转换
        })
    }
}

```