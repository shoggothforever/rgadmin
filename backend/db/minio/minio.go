package dbminio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zeromicro/go-zero/core/logx"
	"mime/multipart"
)

func NewMinioClient(endpoint, accesskey, secretkey, bucket string, useSSL bool) *minio.Client {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accesskey, secretkey, ""),
		Secure: useSSL,
	})
	if err != nil {
		logx.Info("与minio客户端建立连接失败  ", err)
	}
	return minioClient
}
func Upload2Minio(ctx context.Context, client *minio.Client, bucketName string, file multipart.File, fileHeader *multipart.FileHeader) error {
	objectName := fileHeader.Filename
	objectSize := fileHeader.Size
	objectType := fileHeader.Header.Get("Content-Type")
	logx.Info(objectName, objectType, objectSize)
	_, err := client.PutObject(ctx, bucketName, objectName, file, objectSize, minio.PutObjectOptions{
		ContentType: objectType,
	})
	if err != nil {
		return err
	}
	return nil
}
