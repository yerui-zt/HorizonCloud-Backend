package svc

import (
	"HorizonX/api/internal/config"
	"HorizonX/api/internal/middleware"
	"HorizonX/model"
	"HorizonX/rpc/identity/identityservice"
	"HorizonX/rpc/order/orderservice"
	"HorizonX/rpc/user/userservice"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	UserRPC        userservice.UserService
	OrderRPC       orderservice.OrderService
	AuthMiddleware rest.Middleware

	HypervisorGroupModel model.HypervisorGroupModel
	VmPlanModel          model.VmPlanModel
	VmTemplateModel      model.VmTemplateModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		AuthMiddleware: middleware.NewAuthMiddleware(
			c,
			identityservice.NewIdentityService(zrpc.MustNewClient(c.IdentityRPC)),
		).Handle,

		HypervisorGroupModel: model.NewHypervisorGroupModel(sqlx.NewMysql(c.Mysql.DataSource)),
		VmPlanModel:          model.NewVmPlanModel(sqlx.NewMysql(c.Mysql.DataSource)),
		VmTemplateModel:      model.NewVmTemplateModel(sqlx.NewMysql(c.Mysql.DataSource)),

		UserRPC:  userservice.NewUserService(zrpc.MustNewClient(c.UserRPC)),
		OrderRPC: orderservice.NewOrderService(zrpc.MustNewClient(c.OrderRPC)),
	}
}
