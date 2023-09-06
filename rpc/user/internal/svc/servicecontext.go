package svc

import (
	"HorizonX/model"
	"HorizonX/rpc/identity/identityservice"
	"HorizonX/rpc/user/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	Identity identityservice.IdentityService

	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Identity:  identityservice.NewIdentityService(zrpc.MustNewClient(c.IdentityRPC)),
		UserModel: model.NewUserModel(sqlx.NewMysql(c.Mysql.DataSource)),
	}
}
