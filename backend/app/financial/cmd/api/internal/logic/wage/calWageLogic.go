package wage

import (
	"context"

	"Table/app/financial/cmd/api/internal/svc"
	"Table/app/financial/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CalWageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCalWageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CalWageLogic {
	return &CalWageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CalWageLogic) CalWage(req *types.CalReq) (resp *types.CalResp, err error) {
	// todo: add your logic here and delete this line

	return
}
