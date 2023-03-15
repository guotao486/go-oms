/*
 * @Author: GG
 * @Date: 2023-02-28 08:57:38
 * @LastEditTime: 2023-03-14 11:37:28
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\service\service.go
 *
 */
package service

import (
	"context"
	"oms/global"
	"oms/internal/dao"

	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/jinzhu/gorm"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) *Service {
	svc := &Service{ctx: ctx}
	svc.SetDao(global.DBEngine)
	// svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
	// return &Service{ctx: ctx, dao: dao.New(global.DBEngine)}
}

func (s *Service) SetDao(db *gorm.DB) {
	s.dao = dao.New(otgorm.WithContext(s.ctx, db))
}
