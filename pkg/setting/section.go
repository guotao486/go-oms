/*
 * @Author: GG
 * @Date: 2023-01-28 11:04:24
 * @LastEditTime: 2023-01-31 16:39:42
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\pkg\setting\section.go
 *
 */
package setting

import "time"

/**
* @Author $
* @Description //TODO $
* @Date $ $
* @Param $
* @return $
**/

// ServerSettingS
//
// @Description: 服务配置
type ServerSettingS struct {
	RunMode      string        // gin的运行模式
	HttpPort     string        // http监听端口
	ReadTimeout  time.Duration // http读取最大时间
	WriteTimeout time.Duration // http写入最大时间
}

// AppSettingS
//
// @Description: 应用配置
type AppSettingS struct {
	// 模板配置
	TemplatePath       string // 模板文件目录
	TemplateLayoutPath string // 根模板文件目录
	TemplateStuffix    string // 模板文件后缀

	// 分页配置
	DefaultPageSize int    // 默认分页数量
	MaxPageSize     int    // 默认最大分页数量
	LogSavePath     string // 默认应用日志存储地址
	LogFileName     string // 应用日志文件名称
	LogFileExt      string // 应用日志文件后缀

	// 文件上传配置
	UploadSavePath       string   // 文件上传保存目录
	UploadServerUrl      string   // 文件上传后展示地址
	UploadImageMaxSize   int      // 文件上传允许的大小
	UploadImageAllowExts []string // 文件上传允许的后缀
}

// DatabaseSettingS
//
// @Description: 数据库配置
type DatabaseSettingS struct {
	DBType       string // 数据库类型
	UserName     string // 数据库账号
	Password     string // 数据库密码
	Host         string // 数据库地址
	DBName       string // 数据库名称
	TablePrefix  string // 数据库表前缀
	Charset      string // 数据库编码
	ParseTime    bool   // time.Time类型自动转换
	MaxIdleConns int    // 最大空闲连接数量
	MaxOpenConns int    // 最大打开的连接数
}

// JWTSettingS
//
// @Description: JWT配置
type JWTSettingS struct {
	Secret string        // 秘钥
	Issuer string        // 发行人
	Expire time.Duration // 过期时效
}

// EmailSettingS
//
// @Description: email配置
type EmailSettingS struct {
	Host     string
	Port     int
	Username string
	Password string
	IsSSL    bool
	From     string
	To       []string
}

// ReadSection
//
//	@Description: 读取区段配置
//	@receiver s
//	@params k
//	@params v
//	@return error
func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
