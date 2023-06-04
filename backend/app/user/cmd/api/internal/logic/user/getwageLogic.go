package user

import (
	"Table/app/user/cmd/api/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"Table/app/user/cmd/api/internal/svc"
	"Table/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetwageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetwageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetwageLogic {
	return &GetwageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetwageLogic) Getwage(req *types.QuerywageReq) (resp *types.QuerywageResp, err error) {
	// todo: add your logic here and delete this line
	id := l.ctx.Value("payload").(string)
	var wage model.Wage
	var filter bson.D
	if req.Month > 0 && req.Month <= 12 {
		filter = bson.D{{"staffcode", id}, {"year", req.Year}, {"month", req.Month}}
	} else {
		filter = bson.D{{"staffcode", id}, {"year", req.Year}}
	}
	findoption := options.Find()
	cur, err := l.svcCtx.Mongo.Collection("wage").Find(l.ctx, filter, findoption)
	if err != nil {
		return &types.QuerywageResp{Response: types.NewResponse(404, "查询用户工资信息失败"+err.Error())}, err
	}
	defer cur.Close(l.ctx)
	resp = &types.QuerywageResp{
		Response:      types.Response{200, "查询用户工资信息成功"},
		WageBeforeTax: 0,
		Tax:           0,
		ActualWage:    0,
	}
	for cur.Next(l.ctx) {
		err = cur.Decode(&wage)
		if err != nil {
			return &types.QuerywageResp{Response: types.NewResponse(404, "查询用户工资信息失败"+err.Error())}, err
		}
		resp.WageBeforeTax += wage.WageBeforeTax
		resp.Tax += wage.Tax
		resp.ActualWage += wage.ActualWage
	}

	return resp, nil
}
