package user

import (
	"Table/app/user/cmd/api/model"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"Table/app/user/cmd/api/internal/svc"
	"Table/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CalwageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCalwageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CalwageLogic {
	return &CalwageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CalwageLogic) Calwage(req *types.CalwageReq) (resp *types.CalwageResp, err error) {
	if req.WorkTime > 300 || req.WorkTime < 0 {
		return &types.CalwageResp{Response: types.NewResponse(401, "传入非法工时信息")}, errors.New("传入非法工时信息")
	}
	id := l.ctx.Value("payload").(string)
	//logx.WithContext(l.ctx).Info("获取到了id:", id)
	filter := bson.D{{"staffcode", id}}
	var user model.User
	err = l.svcCtx.Mongo.Collection("employee").FindOne(l.ctx, filter).Decode(&user)
	if err != nil {
		return &types.CalwageResp{Response: types.NewResponse(401, err.Error())}, err
	}
	//logx.WithContext(l.ctx).Info("查询到用户信息:", user)
	var base = model.GetBasewage(user.Role)
	base += base*req.WorkTime/1200 + model.ElseFee
	tax := model.CalTax(base)
	wage := model.Wage{
		StaffCode:     user.StaffCode,
		Name:          user.Name,
		Year:          time.Now().Year(),
		Month:         int(time.Now().Month()),
		WorkTime:      req.WorkTime,
		WageBeforeTax: base,
		Tax:           tax,
		ActualWage:    base - tax,
	}
	//logx.WithContext(l.ctx).Info("计算到薪水信息:", wage)
	_, err = l.svcCtx.Mongo.Collection("wage").InsertOne(l.ctx, &wage)
	if err != nil {
		return &types.CalwageResp{Response: types.NewResponse(401, err.Error())}, err
	}
	//logx.WithContext(l.ctx).Info("插入mongoDB成功，id:", res.InsertedID)
	resp = &types.CalwageResp{
		Response:      types.Response{200, "计算当月工资信息成功"},
		StaffCode:     user.StaffCode,
		Name:          user.Name,
		Year:          wage.Year,
		Month:         wage.Month,
		WageBeforeTax: wage.WageBeforeTax,
		Tax:           wage.Tax,
		ActualWage:    wage.ActualWage,
	}
	return resp, nil
}
