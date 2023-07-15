package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	//UserRpcConf zrpc.RpcClientConf
	Mongo struct {
		Uri        string
		Database   string
		Collection string
	}
}
