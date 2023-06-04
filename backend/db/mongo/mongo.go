package dbmongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mime/multipart"
	"time"
)

func NewMongoConn(ctx context.Context, uri string) (*mongo.Client, error) {
	if len(uri) == 0 {
		logrus.Fatal("uri settings corrupt!")
		return nil, errors.New("传入无效uri")
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetTimeout(time.Second*5).SetRetryReads(true))
	if err != nil {
		logrus.Error("Build connect failed")
		return nil, err
	}
	if err = client.Ping(ctx, nil); err != nil {
		logrus.Error("Connect failed", err)
		return nil, err
	} else {
		fmt.Println("Pong!")
	}
	return client, nil
}
func UploadFile2Minio(ctx context.Context, client *minio.Client, bucketName string, file multipart.File, fileheader multipart.FileHeader) error {
	objectName := fileheader.Filename
	objectSize := fileheader.Size
	objectType := fileheader.Header.Get("Content-Type")
	_, err := client.PutObject(ctx, bucketName, objectName, file, objectSize, minio.PutObjectOptions{
		ContentType: objectType,
	})
	if err != nil {
		logx.Errorf("传输文件错误,文件名:%s,文件大小%d,错误：", objectName, objectSize, err)
		return err
	}
	return nil
}
