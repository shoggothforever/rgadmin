package logic

import (
	"context"

	"Table/app/financial/cmd/rpc/internal/svc"
	"Table/app/financial/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CalWageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCalWageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CalWageLogic {
	return &CalWageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CalWageLogic) CalWage(in *pb.CalReq) (*pb.CalResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CalResp{}, nil
}
