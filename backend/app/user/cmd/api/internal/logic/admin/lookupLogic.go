package admin

import (
	"context"

	"Table/app/user/cmd/api/internal/svc"
	"Table/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLookupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookupLogic {
	return &LookupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LookupLogic) Lookup() (resp *types.LookupResp, err error) {
	// todo: add your logic here and delete this line

	return
}
