package logic

import (
	"context"

	"Table/app/user/cmd/rpc/internal/svc"
	"Table/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoadInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoadInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoadInfoLogic {
	return &LoadInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoadInfoLogic) LoadInfo(in *pb.LoadInfoReq) (*pb.LoadInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.LoadInfoResp{}, nil
}
