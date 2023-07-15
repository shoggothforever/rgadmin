package svc

import (
	"Table/app/user/cmd/rpc/internal/config"
	dbmongo "Table/db/mongo"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceContext struct {
	Config config.Config
	//UserRpc user.User
	Mongo *mongo.Database
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn, _ := dbmongo.NewMongoConn(context.Background(), c.Mongo.Uri)
	return &ServiceContext{
		Config: c,
		//UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		Mongo: conn.Database(c.Mongo.Database),
	}
}
