package svc

import (
	"HorizonX/model"
	"HorizonX/rpc/vm/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserModel            model.UserModel
	HypervisorGroupModel model.HypervisorGroupModel
	HypervisorNodeModel  model.HypervisorNodeModel
	VmPlanModel          model.VmPlanModel
	VmInstanceModel      model.VmInstanceModel
	IpGroupModel         model.IpGroupModel
	IpAddressModel       model.IpAddressModel
	SshKeysModel         model.SshKeysModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:               c,
		UserModel:            model.NewUserModel(sqlConn),
		HypervisorGroupModel: model.NewHypervisorGroupModel(sqlConn),
		HypervisorNodeModel:  model.NewHypervisorNodeModel(sqlConn),
		VmPlanModel:          model.NewVmPlanModel(sqlConn),
		VmInstanceModel:      model.NewVmInstanceModel(sqlConn),
		IpGroupModel:         model.NewIpGroupModel(sqlConn),
		IpAddressModel:       model.NewIpAddressModel(sqlConn),
		SshKeysModel:         model.NewSshKeysModel(sqlConn),
	}
}
