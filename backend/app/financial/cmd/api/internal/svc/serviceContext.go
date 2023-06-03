package svc

import (
	model "Table/app/financial/cmd/api/cache"
	"Table/app/financial/cmd/api/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	WageModel model.WageModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		WageModel: model.NewWageModel(c.Mongo.Uri, c.Mongo.Database, c.Mongo.Collection, c.Cache),
	}
}
