/*
 * @Author: GG
 * @Date: 2023-01-28 20:31:27
 * @LastEditTime: 2023-01-30 15:28:00
 * @LastEditors: GG
 * @Description: tag service
 * @FilePath: \oms\internal\service\tag.go
 *
 */
package service

import (
	"oms/internal/model/demo"
	"oms/internal/request"
	"oms/pkg/app"
)

// Tag service

// count
func (s *Service) CountTag(param *request.CountTagRequest) (int, error) {
	return s.dao.CountTag(param.Name, param.State)
}

// GetListTag
func (s *Service) GetListTag(param *request.TagListRequest, pager *app.Pager) ([]*demo.Tag, error) {
	return s.dao.GetListTag(param.Name, param.State, pager.Page, pager.PageSize)
}

// CreateTag
func (s *Service) CreateTag(param *request.CreateTagRequest) error {
	return s.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

// UpdateTag
func (s *Service) UpdateTag(param *request.UpdateTagRequest) error {
	return s.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

// DeleteTag
func (s *Service) DeleteTag(param *request.DeleteTagRequest) error {
	return s.dao.DeleteTag(param.ID)
}

// Tag service end...
