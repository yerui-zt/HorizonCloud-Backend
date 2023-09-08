package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ HypervisorNodeModel = (*customHypervisorNodeModel)(nil)

type (
	// HypervisorNodeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHypervisorNodeModel.
	HypervisorNodeModel interface {
		hypervisorNodeModel
	}

	customHypervisorNodeModel struct {
		*defaultHypervisorNodeModel
	}
)

// NewHypervisorNodeModel returns a model for the database table.
func NewHypervisorNodeModel(conn sqlx.SqlConn) HypervisorNodeModel {
	return &customHypervisorNodeModel{
		defaultHypervisorNodeModel: newHypervisorNodeModel(conn),
	}
}
