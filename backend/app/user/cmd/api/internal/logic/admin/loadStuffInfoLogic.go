package admin

import (
	"context"

	"Table/app/user/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoadStuffInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoadStuffInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoadStuffInfoLogic {
	return &LoadStuffInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoadStuffInfoLogic) LoadStuffInfo() error {
	// todo: add your logic here and delete this line

	return nil
}
