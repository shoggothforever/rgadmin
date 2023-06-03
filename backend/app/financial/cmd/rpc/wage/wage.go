// Code generated by goctl. DO NOT EDIT.
// Source: wage.proto

package wage

import (
	"context"

	"Table/app/financial/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CalReq  = pb.CalReq
	CalResp = pb.CalResp

	Wage interface {
		CalWage(ctx context.Context, in *CalReq, opts ...grpc.CallOption) (*CalResp, error)
	}

	defaultWage struct {
		cli zrpc.Client
	}
)

func NewWage(cli zrpc.Client) Wage {
	return &defaultWage{
		cli: cli,
	}
}

func (m *defaultWage) CalWage(ctx context.Context, in *CalReq, opts ...grpc.CallOption) (*CalResp, error) {
	client := pb.NewWageClient(m.cli.Conn())
	return client.CalWage(ctx, in, opts...)
}
