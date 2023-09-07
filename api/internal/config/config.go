package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	UserRPC     zrpc.RpcClientConf
	IdentityRPC zrpc.RpcClientConf

	Jwt struct {
		Issuer       string
		AccessSecret string
		AccessExpire int64
	}
}
