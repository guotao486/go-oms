/*
 * @Author: GG
 * @Date: 2023-01-30 14:34:14
 * @LastEditTime: 2023-03-14 13:29:13
 * @LastEditors: GG
 * @Description: dao 封装数据访问对象
 * @FilePath: \oms\internal\dao\dao.go
 *
 */
package dao

import "github.com/jinzhu/gorm"

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}
