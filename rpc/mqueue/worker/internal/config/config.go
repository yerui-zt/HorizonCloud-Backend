package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf
	Redis redis.RedisConf

	Mysql struct {
		DataSource string
	}

	VmRPC zrpc.RpcClientConf
}
