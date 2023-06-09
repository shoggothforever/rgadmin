package user

import (
	"Table/app/user/cmd/api/internal/svc"
	"Table/app/user/cmd/api/internal/types"
	"Table/app/user/cmd/api/model"
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
	fileter := bson.D{{"staffcode", id}}
	var user model.User
	err = l.svcCtx.Mongo.Collection(l.svcCtx.Config.Mongo.Collection).FindOne(l.ctx, fileter).Decode(&user)
	if err != nil {
		return &types.UserInfoResp{Response: types.NewResponse(500, err.Error())}, err
	}
	resp = &types.UserInfoResp{Response: types.NewResponse(200, "获取用户信息成功"),
		Id:        user.ID.Hex(),
		StaffCode: user.StaffCode,
		Name:      user.Name,
		Role:      user.Role,
		BaseWage:  user.BaseWage,
		ElseFee:   user.ElseFee}
	return resp, nil
}
