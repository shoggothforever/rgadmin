package user

import (
	"context"
	"github.com/golang-jwt/jwt/v4"

	"Table/app/user/cmd/api/internal/svc"
	"Table/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.UserLogoutReq) (resp *types.UserLogoutResp, err error) {
	// todo: add your logic here and delete this line
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = make(jwt.MapClaims)
	resp = &types.UserLogoutResp{
		Response: types.Response{Code: 200, Msg: "退出成功"},
	}
	return resp, nil
}
