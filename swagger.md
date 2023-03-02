
# swagger 使用
## swagger 安装
```
$ go get -u github.com/swaggo/swag/cmd/swag@v1.6.5
$ go get -u github.com/swaggo/gin-swagger@v1.2.0 
$ go get -u github.com/swaggo/files
$ go get -u github.com/alecthomas/template
```
安装成功后测试
```
$ swag -v
swag version v1.6.5
```
## 编写注解
API AND MAIN 注解编写完成后

路由引入
```
_ "oms/docs"
swaggerFiles "github.com/swaggo/files"
ginSwagger "github.com/swaggo/gin-swagger"
```
路由编写
```
r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```
生成文档
```
swag init
```


## swagger 注解

### 项目注解
```
// @title GO语言编程之旅 blog
// @version 1.0
// @description GO语言编程之旅 blog
// @termsOfService www.bookchat.net
```

### 方法注解
* 入参类型，header body(json),formData,query,path
  
|注解|描述|
|---|---|
|@Tags|文档模块划分|
|@Summary|  摘要|
|@Produce|响应类型，如JSON、XML、HTML|
|@Param|参数格式，从左到右分别是：参数名、入参类型、数据类型、是否必填和注释|
|@Success|响应成功，从左到右分别是：状态码、参数类型、数据类型和注释|
|@Failure|响应失败，从左到右分别是：状态码、参数类型、数据类型和注释|
|@Router|路由，从左到右分别是：路由地址和HTTP方法|

```
// @Tags Upload
// @Summary 文件上传
// @Produce json
// @Param type formData int true "上传类型" default(1)
// @Param file formData file true "file"
// @Success 200 {object} app.Success "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /upload/file [post]
func(u Upload) UploadFile(c *gin.Context){}


// @Tags modules
// @Summary 获取多个标签
// @Produce  json
// @param token header string false "token"
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {}

// @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {}

// @Summary 更新标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Param UpdateTagRequest body request.UpdateTagRequest true "UpdateTagRequest"
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {}

// @Summary 删除标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {}
```