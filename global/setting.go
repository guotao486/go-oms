/*
 * @Author: GG
 * @Date: 2023-01-28 12:03:56
 * @LastEditTime: 2023-02-06 15:31:26
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\global\setting.go
 *
 */
package global

import (
	"oms/pkg/logger"
	"oms/pkg/setting"

	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	IS_DEL_DISABLE = iota
	IS_Del_ENABLE
)

var ModelAutoMigrate []interface{}

var ModeInitData []func()

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS

	DBEngine *gorm.DB

	Logger *logger.Logger

	Validator *binding.StructValidator

	JWTSetting *setting.JWTSettingS

	EmailSetting *setting.EmailSettingS
)
