/*
 * @Author: GG
 * @Date: 2023-01-28 11:04:27
 * @LastEditTime: 2023-03-28 14:24:21
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

type Controller struct {
	IndexC     *controller.IndexController
	AuthC      *controller.AuthController
	UserC      *controller.UserController
	UserGroupC *controller.UserGroupController
	OrderC     *controller.OrderController
}

func NewController() *Controller {
	return &Controller{
		UserC:      controller.NewUser(),
		UserGroupC: controller.NewUserGroup(),
		OrderC:     controller.NewOrder(),
	}
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

	// 中间件
	r.Use(middleware.Sessions())
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
	r.POST("upload/image_file", upload.ImageFiles)

	controller := NewController()
	r.GET("/home", controller.IndexC.Home)
	r.GET("/login", controller.AuthC.Login)
	r.POST("/login", controller.AuthC.Login)
	r.GET("/logout", controller.AuthC.Logout)

	// user module
	userR := r.Group("/user")
	userR.GET("/", controller.UserC.Index)
	userR.GET("/list", controller.UserC.List)
	userR.GET("/create", controller.UserC.Create)
	userR.POST("/create", controller.UserC.Create)
	userR.GET("/update", controller.UserC.Update)
	userR.POST("/update", controller.UserC.Update)
	userR.DELETE("/delete/:id", controller.UserC.Delete)
	// user module end

	// userGroup module
	userGroupR := r.Group("/group")
	userGroupR.GET("/", controller.UserGroupC.Index)
	userGroupR.GET("/list", controller.UserGroupC.List)
	userGroupR.GET("/create", controller.UserGroupC.Create)
	userGroupR.POST("/create", controller.UserGroupC.Create)
	userGroupR.GET("/update", controller.UserGroupC.Update)
	userGroupR.POST("/update", controller.UserGroupC.Update)
	userGroupR.DELETE("/delete/:id", controller.UserGroupC.Delete)
	// userGroup end

	// order
	orderR := r.Group("/order")
	orderR.GET("/", controller.OrderC.Index)
	orderR.GET("/list", controller.OrderC.List)
	orderR.GET("/create", controller.OrderC.Create)
	orderR.POST("/create", controller.OrderC.Create)
	orderR.GET("/update", controller.OrderC.Update)
	orderR.POST("/update", controller.OrderC.Update)
	orderR.DELETE("/delete/:id", controller.OrderC.Delete)
	orderR.GET("/ajax_update/payment", controller.OrderC.AjaxUpdatePayment)
	orderR.POST("/ajax_update/payment", controller.OrderC.AjaxUpdatePayment)
	orderR.GET("/ajax_update/status", controller.OrderC.AjaxUpdateStatus)
	orderR.POST("/ajax_update/status", controller.OrderC.AjaxUpdateStatus)
	orderR.GET("/ajax_update/shipping", controller.OrderC.AjaxUpdateShipping)
	orderR.POST("/ajax_update/shipping", controller.OrderC.AjaxUpdateShipping)

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
