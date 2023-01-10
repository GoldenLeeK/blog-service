package service

import (
	"context"
	"github.com/GoldenLeeK/blog-service/global"
	"github.com/GoldenLeeK/blog-service/interal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) *Service {
	svc := &Service{
		ctx: ctx,
	}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
