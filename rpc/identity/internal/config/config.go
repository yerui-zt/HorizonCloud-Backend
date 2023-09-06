package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Jwt struct {
		Issuer       string
		AccessSecret string
		AccessExpire int64
	}
}
