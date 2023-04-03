/*
 * @Author: GG
 * @Date: 2023-01-30 14:34:14
 * @LastEditTime: 2023-03-22 10:53:19
 * @LastEditors: GG
 * @Description: dao 封装数据访问对象
 * @FilePath: \oms\internal\dao\dao.go
 *
 */
package dao

import (
	"oms/pkg/enum"

	"github.com/jinzhu/gorm"
)

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}

func (d *Dao) WhereLikeString(str string) string {
	return "%" + str + "%"
}

// 预加载所有数据
func PreloadAll(db *gorm.DB) *gorm.DB {
	return db.Set("gorm:auto_preload", true)
}

// 未删除
// is_del = 0
func IsDelToUnable(db *gorm.DB) *gorm.DB {
	return db.Where("is_del = ?", enum.DEFAULT_IS_DEL)
}

func StateToUnable(db *gorm.DB) *gorm.DB {
	return db.Where("state = ?", enum.DEFAULT_STATE)
}
