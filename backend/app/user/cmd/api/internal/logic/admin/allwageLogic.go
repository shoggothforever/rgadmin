package admin

import (
	"context"

	"Table/app/user/cmd/api/internal/svc"
	"Table/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllwageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllwageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllwageLogic {
	return &AllwageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllwageLogic) Allwage(req *types.WageExcelReq) (resp *types.WageExcelResp, err error) {
	// todo: add your logic here and delete this line

	return
}
