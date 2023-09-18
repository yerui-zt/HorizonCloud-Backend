package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ IpGroupModel = (*customIpGroupModel)(nil)

type (
	// IpGroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customIpGroupModel.
	IpGroupModel interface {
		ipGroupModel
	}

	customIpGroupModel struct {
		*defaultIpGroupModel
	}
)

// NewIpGroupModel returns a model for the database table.
func NewIpGroupModel(conn sqlx.SqlConn) IpGroupModel {
	return &customIpGroupModel{
		defaultIpGroupModel: newIpGroupModel(conn),
	}
}
