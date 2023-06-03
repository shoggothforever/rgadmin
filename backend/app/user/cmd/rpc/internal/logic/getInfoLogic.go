package logic

import (
	"context"

	"Table/app/user/cmd/rpc/internal/svc"
	"Table/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoLogic {
	return &GetInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInfoLogic) GetInfo(in *pb.UserInfoReq) (*pb.UserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UserInfoResp{}, nil
}
