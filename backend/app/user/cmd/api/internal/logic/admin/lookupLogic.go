package admin

import (
	dbmongo "Table/db/mongo"
	dbredis "Table/db/redis"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"time"

	"Table/app/user/cmd/api/internal/svc"
	"Table/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLookupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookupLogic {
	return &LookupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LookupLogic) Lookup() (resp *types.LookupResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.LookupResp)
	rdb := l.svcCtx.Redis
	mdb := l.svcCtx.Mongo
	var cnt int64
	var curkey = fmt.Sprintf(dbredis.CurrentNumKey, time.Now().Year())
	resp.CurrentCount, err = rdb.Get(l.ctx, curkey).Result()
	if err != nil && err != redis.Nil {
		resp.Response = types.NewResponse(500, "获取提交用户数据失败"+err.Error())
		return resp, err
	} else if err == redis.Nil {
		cnt, err = mdb.Collection(dbmongo.WageCollection).CountDocuments(l.ctx, bson.D{{"year", time.Now().Year()}, {"month", time.Now().Month()}})
		if err != nil {
			resp.Response = types.NewResponse(500, "获取提交用户数据失败"+err.Error())
			return resp, err
		}
		_, err = rdb.Set(l.ctx, curkey, cnt, time.Second*60).Result()
		if err != nil {
			resp.Response = types.NewResponse(500, "添加缓存信息失败"+err.Error())
			return resp, err
		}
		resp.CurrentCount = strconv.FormatInt(cnt, 10)
	}
	resp.AllCount, err = rdb.Get(l.ctx, dbredis.AllNumKey).Result()
	if err != nil && err != redis.Nil {
		resp.Response = types.NewResponse(500, "获取提交用户数据失败"+err.Error())
		return
	} else if err == redis.Nil {
		cnt, err = mdb.Collection(dbmongo.UserCollection).CountDocuments(l.ctx, bson.D{})
		if err != nil {
			resp.Response = types.NewResponse(500, "获取用户数据失败"+err.Error())
			return resp, err
		}
		_, err = rdb.Set(l.ctx, dbredis.AllNumKey, cnt, time.Second*60).Result()
		if err != nil {
			resp.Response = types.NewResponse(500, "添加缓存信息失败"+err.Error())
			return resp, err
		}
		resp.AllCount = strconv.FormatInt(cnt, 10)
	}

	resp.Response = types.NewResponse(200, "成功获取计数信息")
	return resp, err
}
