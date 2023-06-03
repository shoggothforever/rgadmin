package admin

import (
	"context"

	"Table/app/user/cmd/api/internal/svc"
	"Table/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoadInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoadInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoadInfoLogic {
	return &LoadInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoadInfoLogic) LoadInfo(req *types.LoadInfoReq) (resp *types.LoadInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
