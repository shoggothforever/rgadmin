package logic

import (
	"Table/app/user/cmd/api/model"
	dbmongo "Table/db/mongo"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"Table/app/user/cmd/api/internal/svc"
	"Table/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	fileter := bson.D{{"staffcode", req.StaffCode}}
	var user model.User
	err = l.svcCtx.Mongo.Collection(dbmongo.UserCollection).FindOne(l.ctx, fileter).Decode(&user)
	if err != nil {
		return &types.UserLoginResp{Response: types.NewResponse(401, err.Error())}, err
	}
	if req.Password != user.Password {
		return &types.UserLoginResp{Response: types.NewResponse(401, "登录失败，请检查用户名或密码是否正确")}, errors.New("登录失败，请检查用户名或密码是否正确")
	}
	//jwt, err := l.svcCtx.JwtModel.GeneraterJwt(user.StuffCode, user.Name, user.Password, time.Duration(l.svcCtx.Config.JwtAuth.AccessExpire))
	jwt, err := types.GetJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.JwtAuth.AccessExpire, req.StaffCode)
	resp = &types.UserLoginResp{
		Response:    types.Response{Code: 200, Msg: "登录成功"},
		Role:        user.Role,
		AccessToken: jwt,
	}
	return resp, nil
}
