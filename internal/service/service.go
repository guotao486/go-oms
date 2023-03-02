package service

import (
	"context"
	"oms/global"
	"oms/internal/dao"

	otgorm "github.com/eddycjy/opentracing-gorm"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) *Service {
	svc := &Service{ctx: ctx}
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
	// return &Service{ctx: ctx, dao: dao.New(global.DBEngine)}
}
