/*
 * @Author: GG
 * @Date: 2023-01-30 14:42:29
 * @LastEditTime: 2023-01-31 16:02:58
 * @LastEditors: GG
 * @Description: tag dao action
 * @FilePath: \oms\internal\dao\tag.go
 *
 */
package dao

import (
	"oms/internal/model"
	"oms/internal/model/demo"
	"oms/pkg/app"
)

func (d Dao) CountTag(name string, state uint8) (int, error) {
	tag := demo.NewTag()
	tag.Name = name
	tag.State = state
	return tag.Count(d.engine)
}

func (d Dao) GetListTag(name string, state uint8, page, pageSize int) ([]*demo.Tag, error) {
	tag := demo.NewTag()
	tag.Name = name
	tag.State = state
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := demo.NewTag()
	tag.Name = name
	tag.State = state
	// tag.Model = &model.Model{CreatedBy: createdBy}
	// tag.CreatedBy = createdBy

	return tag.Create(d.engine)
}

func (d Dao) UpdateTag(id uint32, name string, state uint8, modifieBy string) error {
	tag := demo.NewTag()
	tag.Model = &model.Model{ID: id}
	value := map[string]interface{}{
		"state": state,
		// "modified_by": modifieBy,
	}
	if name != "" {
		value["name"] = name
	}
	return tag.Update(d.engine, value)
}

func (d Dao) DeleteTag(id uint32) error {
	tag := demo.NewTag()
	tag.Model = &model.Model{ID: id}
	return tag.Delete(d.engine)
}
