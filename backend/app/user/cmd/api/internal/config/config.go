package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Cache   cache.CacheConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	Mongo struct {
		Uri        string
		Database   string
		Collection string
	}
	Minio struct {
		Bucket     string
		EndPoint   string
		AccessKey  string
		SecretKey  string
		LinkExpire int64
	}
}
