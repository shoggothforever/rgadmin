package svc

import (
	model "Table/app/user/cmd/api/cache"
	"Table/app/user/cmd/api/internal/config"
	common "Table/common/middleware"
	"github.com/zeromicro/go-zero/core/stores/mon"
)

type ServiceContext struct {
	Config    config.Config
	Mongo     *mon.Model
	UserModel model.UserModel
	JwtModel  common.Claims
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Mongo:     mon.MustNewModel(c.Mongo.Uri, c.Mongo.Database, c.Mongo.Collection),
		UserModel: model.NewUserModel(c.Mongo.Uri, c.Mongo.Database, c.Mongo.Collection, c.Cache),
		JwtModel:  common.NewClaim(),
	}
}
