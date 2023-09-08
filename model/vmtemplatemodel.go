package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ VmTemplateModel = (*customVmTemplateModel)(nil)

type (
	// VmTemplateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVmTemplateModel.
	VmTemplateModel interface {
		vmTemplateModel
	}

	customVmTemplateModel struct {
		*defaultVmTemplateModel
	}
)

// NewVmTemplateModel returns a model for the database table.
func NewVmTemplateModel(conn sqlx.SqlConn) VmTemplateModel {
	return &customVmTemplateModel{
		defaultVmTemplateModel: newVmTemplateModel(conn),
	}
}
