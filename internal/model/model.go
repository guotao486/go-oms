/*
 * @Author: GG
 * @Date: 2023-01-28 11:04:27
 * @LastEditTime: 2023-03-22 15:53:05
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\model\model.go
 *
 */
package model

import (
	"fmt"
	"oms/global"
	"oms/pkg/enum"
	"oms/pkg/setting"
	"time"

	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID uint32 `gorm:"primary_key" json:"id"` // id
	// CreatedBy  string `json:"created_by"`            // 创建人
	// ModifiedBy string `json:"modified_by"`           // 修改人
	CreatedOn  uint32 `json:"created_on"`  // 创建时间
	ModifiedOn uint32 `json:"modified_on"` // 修改时间
	DeletedOn  uint32 `json:"deleted_on"`  // 删除时间
	IsDel      uint8  `json:"is_del"`      // 是否删除 0 未删除 1 已删除
}

// NewDBEngine
//
//	@Description: 数据库连接
//	@params databaseSetting
//	@return *gorm.DB
//	@return error
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))

	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}

	// 单数表
	db.SingularTable(true)
	// 设置连接池
	// 最大空闲连接数
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	// 最大连接数
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	// 注册并替换回调函数,Replace 替换原有的回调函数
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	// 链路追踪 SQL追踪
	otgorm.AddGormCallbacks(db)
	return db, nil
}

// model callback 回调处理，替换gorm原有回调

// updateTimeStampForCreateCallback
//
// 通过调用 scope.FieldByName 方法，获取当前是否包含所需的字段。
//
// 通过判断 Field.IsBlank 的值，可以得知该字段的值是否为空。
//
// 若为空，则会调用 Field.Set 方法给该字段设置值，入参类型为 interface{}，内部也就是通过反射进行一系列操作赋值。
//
/**
 * @description: 新增行为的回调
 * @param {*gorm.Scope} scope
 * @return {*}
 */
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}

		if updateTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if updateTimeField.IsBlank {
				_ = updateTimeField.Set(nowTime)
			}
		}

	}
}

// updateTimeStampForUpdateCallback
//
// 通过调用 scope.Get("gorm:update_column") 去获取当前设置了标识 gorm:update_column 的字段属性。
//
// 若不存在，也就是没有自定义设置 update_column，那么将会在更新回调内设置默认字段 ModifiedOn 的值为当前的时间戳。
//
/**
 * @description: 更新行为的回调
 * @param {*gorm.Scope} scope
 * @return {*}
 */
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// deleteCallback
//
// 通过调用 scope.Get("gorm:delete_option") 去获取当前设置了标识 gorm:delete_option 的字段属性。
//
// 判断是否存在 DeletedOn 和 IsDel 字段，若存在则调整为执行 UPDATE 操作进行软删除（修改 DeletedOn 和 IsDel 的值），否则执行 DELETE 进行硬删除。
//
// 调用 scope.QuotedTableName 方法获取当前所引用的表名，并调用一系列方法针对 SQL 语句的组成部分进行处理和转移，
//
// scope.Quote(deletedOnField.DBName) 获取字段名
//
// scope.AddToVars(1) 设置值为1
//
// 最后在完成一些所需参数设置后调用 scope.CombinedConditionSql 方法完成 SQL 语句的组装
//
/**
 * @description:
 * @param {*gorm.Scope} scope
 * @return {*}
 */
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		deletedOnField, hasDeleteOnField := scope.FieldByName("DeletedOn")
		isDelField, hasIsDelField := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeleteOnField && hasIsDelField {
			now := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(enum.IS_DEL_ENABLE),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}

}

// addExtraSpaceIfExist
//
/**
 * @description: 删除回调使用的字符串处理，前面加空格
 * @param {string} str
 * @return {*}
 */
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

// model callback end...
