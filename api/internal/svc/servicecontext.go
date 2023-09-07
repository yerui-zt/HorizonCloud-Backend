package svc

import (
	"HorizonX/api/internal/config"
	"HorizonX/api/internal/middleware"
	"HorizonX/rpc/identity/identityservice"
	"HorizonX/rpc/user/userservice"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	UserRPC        userservice.UserService
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRPC: userservice.NewUserService(zrpc.MustNewClient(c.UserRPC)),
		AuthMiddleware: middleware.NewAuthMiddleware(
			c,
			identityservice.NewIdentityService(zrpc.MustNewClient(c.IdentityRPC)),
		).Handle,
	}
}
