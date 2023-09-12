package svc

import (
	"HorizonX/model"
	"HorizonX/rpc/order/internal/config"
	"HorizonX/rpc/payment/paymentservice"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	PaymentRPC paymentservice.PaymentService

	OrderModel               model.OrderModel
	OrderItemModel           model.OrderItemModel
	UserModel                model.UserModel
	VmPlanModel              model.VmPlanModel
	HypervisorGroupModel     model.HypervisorGroupModel
	VmTemplateModel          model.VmTemplateModel
	SystemPaymentMethodModel model.SystemPaymentMethodModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,

		PaymentRPC: paymentservice.NewPaymentService(zrpc.MustNewClient(c.PaymentRPC)),

		OrderModel:               model.NewOrderModel(sqlConn),
		OrderItemModel:           model.NewOrderItemModel(sqlConn),
		UserModel:                model.NewUserModel(sqlConn),
		VmPlanModel:              model.NewVmPlanModel(sqlConn),
		HypervisorGroupModel:     model.NewHypervisorGroupModel(sqlConn),
		VmTemplateModel:          model.NewVmTemplateModel(sqlConn),
		SystemPaymentMethodModel: model.NewSystemPaymentMethodModel(sqlConn),
	}
}
