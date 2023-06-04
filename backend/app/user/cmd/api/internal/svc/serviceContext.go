package svc

import (
	"Table/app/user/cmd/api/internal/config"
	"Table/app/user/cmd/api/internal/types"
	"Table/app/user/cmd/api/model"
	db "Table/db/mongo"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceContext struct {
	Config    config.Config
	Mongo     *mongo.Database
	UserModel model.UserModel
	JwtModel  types.Claims
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn, _ := db.NewMongoConn(context.Background(), c.Mongo.Uri)
	return &ServiceContext{
		Config:    c,
		Mongo:     conn.Database(c.Mongo.Database),
		UserModel: model.NewUserModel(c.Mongo.Uri, c.Mongo.Database, c.Mongo.Collection, c.Cache),
		JwtModel:  types.NewClaim(),
	}
}
