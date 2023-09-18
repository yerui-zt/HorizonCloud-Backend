package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ IpAddressModel = (*customIpAddressModel)(nil)

type (
	// IpAddressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customIpAddressModel.
	IpAddressModel interface {
		ipAddressModel
	}

	customIpAddressModel struct {
		*defaultIpAddressModel
	}
)

// NewIpAddressModel returns a model for the database table.
func NewIpAddressModel(conn sqlx.SqlConn) IpAddressModel {
	return &customIpAddressModel{
		defaultIpAddressModel: newIpAddressModel(conn),
	}
}
