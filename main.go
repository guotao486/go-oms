/*
 * @Author: GG
 * @Date: 2023-02-28 08:57:23
 * @LastEditTime: 2023-04-06 16:41:01
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\main.go
 *
 */
/*
 * @Author: GG
 * @Date: 2023-02-28 08:57:23
 * @LastEditTime: 2023-03-30 10:33:42
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\main.go
 *
 */
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"oms/global"
	"oms/internal/model"
	"oms/internal/routers"
	"oms/pkg/cache"
	"oms/pkg/logger"
	pkg_setting "oms/pkg/setting"
	"oms/pkg/tracer"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// build 时命令行参数
var (
	port    string
	runMode string
	config  string

	// 版本信息
	/**
	 * 编译命令
	 * go build -ldflags \
	 * "-X main.buildTime=`date +%Y-%m-%d,%H:%M:%S` -X main.buildVersion=1.0.0 -X main.gitCommitID=`git rev-parse HEAD`"
	 *
	 * 查看信息
	 * ./build.exe -version
	 */
	isVersion    bool
	buildTime    string
	buildVersion string
	gitCommitID  string
)

// init 初始化操作
func init() {
	setupFlag()
	// 读取配置
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	// 日志
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	// 数据库连接
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	// 缓存
	err = setupCacheStore()
	if err != nil {
		log.Fatalf("init.setupCacheStore err: %v", err)
	}
	// 自定义验证器
	err = setupValidator()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	// 链路追踪
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
}

// @title GO语言编程之旅 blog
// @version 1.0
// @description GO语言编程之旅 blog
// @termsOfService www.bookchat.net
func main() {
	// 第一版
	//r := gin.Default()
	//r.GET("/ping", func(context *gin.Context) {
	//	context.JSON(200, gin.H{"message": "pong"})
	//})
	//r.Run()

	// 第二版
	//router := routers.NewRouter()
	//s := &http.Server{
	//	Addr:           ":8080",          // 监听的端口
	//	Handler:        router,           // 处理的程序
	//	ReadTimeout:    10 * time.Second, // 读取最大时间
	//	WriteTimeout:   10 * time.Second, // 写入最大时间
	//	MaxHeaderBytes: 1 << 20,          // 请求头最大字节数
	//}
	//// 开始监听
	//s.ListenAndServe()

	// 编译版本信息
	if isVersion {
		fmt.Printf("buildTime: %v\n", buildTime)
		fmt.Printf("buildVersion: %v\n", buildVersion)
		fmt.Printf("gitCommitID: %v\n", gitCommitID)
		return
	}
	// 第三版
	gin.SetMode(global.ServerSetting.RunMode)

	router := routers.NewRouter()
	global.Logger.Info(context.Background(), "guotao", "oms")

	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	// s.ListenAndServe()

	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			global.Logger.Fatalf(context.Background(), "s.ListenAndServe err: %v", err)
		}
	}()

	// 优雅重启和停止
	// 等待信号中断
	quit := make(chan os.Signal)
	// 接收syscall.SIGINT 和 syscall.SIGTERM 信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shuting down server...")

	// 最大时间控制，通知服务有5秒时间来处理原有的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		global.Logger.Fatal(ctx, "server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}

// setupFlag
//
// @Description: 获取命令行参数
// @return error
func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()
	return nil
}

// setupSetting
//
//	@Description: 读取配置步骤
//	@return error
func setupSetting() error {
	// 读取配置
	setting, err := pkg_setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}

	// 读取Server区域配置
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	// 读取App区域配置
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	// 读取Database区域配置
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Cache", &global.CacheSetting)
	if err != nil {
		return err
	}

	// 服务配置读取/写入时间
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	// JWT配置
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second

	// email配置
	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	// 将build 命令行的参数覆盖文本配置
	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}
	return nil
}

// setupDBEngine
//
//	@Description: 数据库连接
//	@return error
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	if len(global.ModelAutoMigrate) > 0 {
		global.DBEngine.AutoMigrate(global.ModelAutoMigrate...)
	}

	if len(global.ModeInitData) > 0 {
		for _, f := range global.ModeInitData {
			f()
		}
	}
	return nil
}

func setupCacheStore() error {
	var err error
	global.CacheStore, err = cache.NewCacheStore(global.CacheSetting)
	if err != nil {
		return err
	}
	return nil
}

// setupLogger
//
//	@Description: 日志
//	@return error
func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    600,  // 最大占用空间 600M
		MaxAge:     10,   // 最大生存周期 10 天
		MaxBackups: 10,   //
		LocalTime:  true, // 文件名的时间格式为本地时间
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

// setupValidator
//
// @Description: 验证器
// @return error
func setupValidator() error {
	// 将你所自定义的 validator 写入

	//自定义验证规则
	var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
		return false
	}

	//注册自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//绑定第一个参数是校验规则的函数第二个参数是自定义的验证函数
		v.RegisterValidation("bookabledate", bookableDate)
	}

	// binding.Validator = *global.Validator
	return nil
}

// setupTracer
//
// @Description: jaeger tracer 链路追踪
// @return error
func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJagerTracer(
		"oms",
		"110.40.208.203:6831",
	)
	if err != nil {
		return err
	}

	global.Tracer = jaegerTracer
	return nil
}
