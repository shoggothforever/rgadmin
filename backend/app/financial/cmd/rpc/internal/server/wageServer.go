// Code generated by goctl. DO NOT EDIT.
// Source: wage.proto

package server

import (
	"context"

	"Table/app/financial/cmd/rpc/internal/logic"
	"Table/app/financial/cmd/rpc/internal/svc"
	"Table/app/financial/cmd/rpc/pb"
)

type WageServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedWageServer
}

func NewWageServer(svcCtx *svc.ServiceContext) *WageServer {
	return &WageServer{
		svcCtx: svcCtx,
	}
}

func (s *WageServer) CalWage(ctx context.Context, in *pb.CalReq) (*pb.CalResp, error) {
	l := logic.NewCalWageLogic(ctx, s.svcCtx)
	return l.CalWage(in)
}
