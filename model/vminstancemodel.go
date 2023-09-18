package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ VmInstanceModel = (*customVmInstanceModel)(nil)

type (
	// VmInstanceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVmInstanceModel.
	VmInstanceModel interface {
		vmInstanceModel
	}

	customVmInstanceModel struct {
		*defaultVmInstanceModel
	}
)

// NewVmInstanceModel returns a model for the database table.
func NewVmInstanceModel(conn sqlx.SqlConn) VmInstanceModel {
	return &customVmInstanceModel{
		defaultVmInstanceModel: newVmInstanceModel(conn),
	}
}
