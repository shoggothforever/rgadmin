package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"

	"Table/app/user/cmd/api/internal/svc"
	"Table/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsergetleaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUsergetleaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsergetleaveLogic {
	return &UsergetleaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UsergetleaveLogic) Usergetleave() (resp *types.GetLeaveResp, err error) {
	// todo: add your logic here and delete this line
	id := l.ctx.Value("payload").(string)
	fileter := bson.D{{"staffcode", id}}
	resp = new(types.GetLeaveResp)
	cur, err := l.svcCtx.Mongo.Collection("leaverequest").Find(l.ctx, fileter)
	if err != nil {
		return &types.GetLeaveResp{Response: types.NewResponse(500, err.Error())}, err
	}
	defer cur.Close(l.ctx)
	var val types.LeaveResp
	for cur.Next(l.ctx) {
		cur.Decode(&val)
		resp.LeaveResps = append(resp.LeaveResps, val)
	}

	resp.Response.Code = 200
	resp.Msg = "获取所有事假请求成功"
	return resp, nil
}
