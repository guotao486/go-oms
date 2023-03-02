/*
 * @Author: GG
 * @Date: 2023-01-28 11:04:27
 * @LastEditTime: 2023-02-28 10:21:22
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\model\demo\tag.go
 *
 */
package demo

import (
	"oms/internal/model"
	"oms/pkg/app"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	*model.Model
	Name  string `json:"name"`  // 标签名称
	State uint8  `json:"state"` // 状态 0 禁用 1 启用
}

func (t Tag) TableName() string {
	return "blog_tag"
}

// NewTag
//
/**
 * @description: 实例化一个Model.Tag
 * @param {string} name
 * @param {uint8} State
 * @return {*}
 */
func NewTag() *Tag {
	return &Tag{}
}

// 定义一个针对swagger 的对象
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

// model action function

// Tag.Count()
/**
 * @description: 根据name 和 state 查询正常的标签数量
 * @param {*gorm.DB} db
 * @return {*}
 */
func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	db = db.Model(&t)
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)

	if err := db.Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// Tag.List()
/**
 * @description: 根据 name 和 state 分页查询正常的标签列表
 * @param {*gorm.DB} db
 * @param {*} pageOffest
 * @param {int} pageSize
 * @return {*}
 */
func (t Tag) List(db *gorm.DB, pageOffest, pageSize int) ([]*Tag, error) {
	var tags []*Tag

	// 分页查询
	if pageOffest >= 0 && pageSize > 0 {
		db = db.Offset(pageOffest).Limit(pageSize)
	}

	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}

	db = db.Where("state = ?", t.State)
	if err := db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

// Tag.Create()
/**
 * @description: 新增
 * @param {*gorm.DB} db
 * @return {*}
 */
func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, value interface{}) error {
	return db.Model(t).Where("id = ? and is_del = ?", t.ID, 0).Updates(value).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? and is_del = ?", t.ID, 0).Delete(&t).Error
}
