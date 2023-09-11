package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SystemPaymentMethodModel = (*customSystemPaymentMethodModel)(nil)

type (
	// SystemPaymentMethodModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemPaymentMethodModel.
	SystemPaymentMethodModel interface {
		systemPaymentMethodModel
	}

	customSystemPaymentMethodModel struct {
		*defaultSystemPaymentMethodModel
	}
)

// NewSystemPaymentMethodModel returns a model for the database table.
func NewSystemPaymentMethodModel(conn sqlx.SqlConn) SystemPaymentMethodModel {
	return &customSystemPaymentMethodModel{
		defaultSystemPaymentMethodModel: newSystemPaymentMethodModel(conn),
	}
}
