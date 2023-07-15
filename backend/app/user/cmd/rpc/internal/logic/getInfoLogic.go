package logic

import (
	"Table/app/user/cmd/api/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"

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

	fileter := bson.D{{"staffcode", in.Staffcode}}
	var user model.User
	//err = l.svcCtx.Mongo.Collection(l.svcCtx.Config.Mongo.Collection).FindOne(l.ctx, fileter).Decode(&user)
	//if err != nil {
	//	return &types.UserInfoResp{Response: types.NewResponse(500, err.Error())}, err
	//}
	l.svcCtx.Mongo.Collection("employee").FindOne(l.ctx, fileter).Decode(&user)
	resp := pb.UserInfoResp{Id: user.ID.String(), Name: user.Name, Staffcode: user.StaffCode, Role: user.Role, Basewage: user.BaseWage, Elsefee: user.ElseFee}
	return &resp, nil
}
