package svc

import (
	"HorizonX/api/internal/config"
	"HorizonX/rpc/user/userservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRPC userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRPC: userservice.NewUserService(zrpc.MustNewClient(c.UserRPC)),
	}
}
