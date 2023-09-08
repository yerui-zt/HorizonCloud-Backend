package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HypervisorGroupModel = (*customHypervisorGroupModel)(nil)

type (
	// HypervisorGroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHypervisorGroupModel.
	HypervisorGroupModel interface {
		hypervisorGroupModel
	}

	customHypervisorGroupModel struct {
		*defaultHypervisorGroupModel
	}
)

// NewHypervisorGroupModel returns a model for the database table.
func NewHypervisorGroupModel(conn sqlx.SqlConn) HypervisorGroupModel {
	return &customHypervisorGroupModel{
		defaultHypervisorGroupModel: newHypervisorGroupModel(conn),
	}
}
