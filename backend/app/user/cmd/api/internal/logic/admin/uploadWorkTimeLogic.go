package admin

import (
	"context"

	"Table/app/user/cmd/api/internal/svc"
	"Table/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadWorkTimeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadWorkTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadWorkTimeLogic {
	return &UploadWorkTimeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadWorkTimeLogic) UploadWorkTime() (resp *types.UploadInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
