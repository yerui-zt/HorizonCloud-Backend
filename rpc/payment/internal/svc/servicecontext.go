package svc

import (
	"HorizonX/model"
	"HorizonX/rpc/payment/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	SystemPaymentMethodModel model.SystemPaymentMethodModel
	OrderModel               model.OrderModel
	UserModel                model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:                   c,
		SystemPaymentMethodModel: model.NewSystemPaymentMethodModel(sqlConn),
		OrderModel:               model.NewOrderModel(sqlConn),
		UserModel:                model.NewUserModel(sqlConn),
	}
}
