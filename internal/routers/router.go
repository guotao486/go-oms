/*
 * @Author: GG
 * @Date: 2023-01-28 11:04:27
 * @LastEditTime: 2023-03-02 14:42:32
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\routers\router.go
 *
 */
package routers

import (
	"net/http"
	_ "oms/docs"
	"oms/global"
	"oms/internal/controller"
	"oms/internal/middleware"
	"oms/pkg/app"
	"oms/pkg/limiter"
	"time"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var methodLimiter = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

// gin 默认单模板，要使用block template 的话需要这个包
func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	commonTemplate := []string{"templates/common/base.html", "templates/common/head.html", "templates/common/header.html", "templates/common/menus.html", "templates/common/footer.html", "templates/common/footer_script.html"}
	r.AddFromFiles("login", "templates/login/index.html")

	r.AddFromFiles("user/create", append(commonTemplate, "templates/user/create.html")...)
	return r
}

func NewRouter() *gin.Engine {
	r := gin.New()
	// r.LoadHTMLGlob("templates/**/*") // gin 默认单模板,继承会发生block覆盖
	// r.HTMLRender = createMyRender() //
	r.HTMLRender = app.LoadTemplateFiles()
	r.Static("/assets", "./assets")
	if global.ServerSetting.RunMode == "debug" {
		// 默认
		r.Use(gin.Logger())   // 输出请求日志中间件
		r.Use(gin.Recovery()) // 异常捕获中间件
	} else {
		// 自定义
		r.Use(middleware.AccessLog()) // 响应日志
		r.Use(middleware.Recovery())  // 异常捕获
	}

	r.Use(middleware.Tracing())                        // 链路追踪
	r.Use(middleware.RateLimiter(methodLimiter))       // 限流器
	r.Use(middleware.ContextTimeout(60 * time.Second)) // 超时控制
	r.Use(middleware.Translations())                   // 参数验证国际化处理中间件

	// 资源路由
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	// swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 文件上传路由
	upload := NewUpload()
	r.POST("upload/file", upload.UploadFile)

	homeC := controller.NewIndex()
	r.GET("/home", homeC.Home)
	authC := controller.NewAuth()
	r.GET("/login", authC.Login)

	userC := controller.NewUser()
	userR := r.Group("/user")
	userR.GET("/list", userC.List)
	userR.GET("/create", userC.Create)
	userR.POST("/create", userC.Create)
	r.GET("/order/index", nil)
	// auth 路由
	// r.POST("/auth", api.GetAuth)

	// // 业务路由
	// tag := v1.NewTag()
	// article := v1.NewArticle()
	// apiv1 := r.Group("/api/v1/")
	// apiv1.Use(middleware.JWT())
	// {
	// 	apiv1.POST("/tags", tag.Create)
	// 	apiv1.DELETE("/tags/:id", tag.Delete)
	// 	apiv1.PUT("/tags/:id", tag.Update)
	// 	apiv1.PATCH("/tags/:id/state", tag.Update)
	// 	apiv1.GET("/tags", tag.List)

	// 	apiv1.POST("/articles", article.Create)
	// 	apiv1.DELETE("/articles/:id", article.Delete)
	// 	apiv1.PUT("/articles/:id", article.Update)
	// 	apiv1.PATCH("/articles/:id/state", article.Update)
	// 	apiv1.GET("/articles/:id", article.Get)
	// 	apiv1.GET("/articles", article.List)
	// }
	return r
}
