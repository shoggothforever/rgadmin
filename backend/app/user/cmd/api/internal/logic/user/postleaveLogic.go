package user

import (
	"Table/app/user/cmd/api/internal/svc"
	"Table/app/user/cmd/api/internal/types"
	"Table/app/user/cmd/api/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostleaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostleaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostleaveLogic {
	return &PostleaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostleaveLogic) Postleave(req *types.PostLeaveReq) (resp *types.PostLeaveResp, err error) {
	// todo: add your logic here and delete this line
	id := l.ctx.Value("payload").(string)
	fileter := bson.D{{"staffcode", id}}
	var user model.User
	err = l.svcCtx.Mongo.Collection(l.svcCtx.Config.Mongo.Collection).FindOne(l.ctx, fileter).Decode(&user)
	if err != nil {
		return &types.PostLeaveResp{Response: types.NewResponse(500, err.Error())}, err
	}
	var lr = model.Leavereq{
		ID:        primitive.ObjectID{},
		Status:    nil,
		Subject:   req.Subject,
		Reason:    req.Reason,
		StaffCode: user.StaffCode,
		Name:      user.Name,
		Date:      time.Now().Format("2006-01-02"),
	}
	_, err = l.svcCtx.Mongo.Collection("leaverequest").InsertOne(l.ctx, &lr)
	if err != nil {
		return &types.PostLeaveResp{Response: types.NewResponse(500, err.Error())}, err
	}
	resp = &types.PostLeaveResp{Response: types.Response{200, "事假申请已提交，请耐心等候审批"}}
	return resp, nil
}
