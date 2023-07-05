package admin

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"

	"Table/app/user/cmd/api/internal/svc"
	"Table/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApproveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApproveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApproveLogic {
	return &ApproveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApproveLogic) Approve(req *types.ApproveReq) (resp *types.ApproveResp, err error) {
	// todo: add your logic here and delete this line
	logx.Info("获取到的请求体内容是", req)
	filter := bson.D{{"staffcode", req.StaffCode}, {"subject", req.Subject}, {"date", req.Date}, {"status", false}}
	update := bson.D{{"$set", bson.D{{"status", req.Status}}}}
	_, err = l.svcCtx.Mongo.Collection("leaverequest").UpdateOne(l.ctx, filter, update)
	if err != nil {
		return &types.ApproveResp{Response: types.NewResponse(401, err.Error())}, err
	}
	resp = &types.ApproveResp{types.Response{Code: 200, Msg: "审批事假成功"}}
	return
}
