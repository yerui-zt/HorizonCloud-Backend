package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SshKeysModel = (*customSshKeysModel)(nil)

type (
	// SshKeysModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSshKeysModel.
	SshKeysModel interface {
		sshKeysModel
	}

	customSshKeysModel struct {
		*defaultSshKeysModel
	}
)

// NewSshKeysModel returns a model for the database table.
func NewSshKeysModel(conn sqlx.SqlConn) SshKeysModel {
	return &customSshKeysModel{
		defaultSshKeysModel: newSshKeysModel(conn),
	}
}
