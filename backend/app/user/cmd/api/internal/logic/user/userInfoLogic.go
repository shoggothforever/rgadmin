package user

import (
	model "Table/app/user/cmd/api/cache"
	"Table/app/user/cmd/api/internal/svc"
	"Table/app/user/cmd/api/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line
	id := l.ctx.Value("payload").(string)
	fileter := bson.D{{"stuffcode", id}}
	var user model.User
	err = l.svcCtx.Mongo.Collection(l.svcCtx.Config.Mongo.Collection).FindOne(l.ctx, fileter).Decode(&user)
	if err != nil {
		return nil, err
	}
	resp = &types.UserInfoResp{Response: types.NewResponse(200, "获取用户信息成功"), Id: user.ID, StuffCode: id, Name: user.Name, Role: user.Role}
	return resp, nil
}
