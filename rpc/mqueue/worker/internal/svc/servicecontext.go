package svc

import (
	"HorizonX/model"
	"HorizonX/rpc/mqueue/worker/internal/config"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	AsynqClient *asynq.Client
	AsynqServer *asynq.Server

	OrderModel     model.OrderModel
	OrderItemModel model.OrderItemModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		AsynqClient: asynq.NewClient(&asynq.RedisClientOpt{Addr: c.Redis.Host}),
		AsynqServer: newAsynqServer(c),

		OrderModel:     model.NewOrderModel(sqlConn),
		OrderItemModel: model.NewOrderItemModel(sqlConn),
	}
}
