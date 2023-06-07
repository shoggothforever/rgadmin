package svc

import (
	"Table/app/user/cmd/api/internal/config"
	"Table/app/user/cmd/api/internal/middleware"
	"Table/app/user/cmd/api/internal/types"
	"Table/app/user/cmd/api/model"
	"Table/db/minio"
	"Table/db/mongo"
	dbredis "Table/db/redis"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
	"github.com/zeromicro/go-zero/rest"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceContext struct {
	Config    config.Config
	Mongo     *mongo.Database
	UserModel model.UserModel
	JwtModel  types.Claims
	Minio     *minio.Client
	Redis     *redis.Client
	Coors     rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn, _ := dbmongo.NewMongoConn(context.Background(), c.Mongo.Uri)
	return &ServiceContext{
		Config:    c,
		Mongo:     conn.Database(c.Mongo.Database),
		UserModel: model.NewUserModel(c.Mongo.Uri, c.Mongo.Database, c.Mongo.Collection, c.Cache),
		JwtModel:  types.NewClaim(),
		Minio:     dbminio.NewMinioClient(c.Minio.EndPoint, c.Minio.AccessKey, c.Minio.SecretKey, c.Minio.Bucket, false),
		Redis:     dbredis.NewRedisClient(),
		Coors:     middleware.NewCoorsMiddleware().Handle,
	}
}
