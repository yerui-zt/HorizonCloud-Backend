package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ VmPlanModel = (*customVmPlanModel)(nil)

type (
	// VmPlanModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVmPlanModel.
	VmPlanModel interface {
		vmPlanModel
	}

	customVmPlanModel struct {
		*defaultVmPlanModel
	}
)

// NewVmPlanModel returns a model for the database table.
func NewVmPlanModel(conn sqlx.SqlConn) VmPlanModel {
	return &customVmPlanModel{
		defaultVmPlanModel: newVmPlanModel(conn),
	}
}
