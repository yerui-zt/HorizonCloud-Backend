package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	Mysql struct {
		DataSource string
	}

	UserRPC     zrpc.RpcClientConf
	IdentityRPC zrpc.RpcClientConf

	Jwt struct {
		Issuer       string
		AccessSecret string
		AccessExpire int64
	}
}
