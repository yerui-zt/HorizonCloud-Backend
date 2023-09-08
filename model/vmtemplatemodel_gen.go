// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	vmTemplateFieldNames          = builder.RawFieldNames(&VmTemplate{})
	vmTemplateRows                = strings.Join(vmTemplateFieldNames, ",")
	vmTemplateRowsExpectAutoSet   = strings.Join(stringx.Remove(vmTemplateFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	vmTemplateRowsWithPlaceHolder = strings.Join(stringx.Remove(vmTemplateFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	vmTemplateModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *VmTemplate) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*VmTemplate, error)
		FindOneByName(ctx context.Context, name string) (*VmTemplate, error)
		// Update(ctx context.Context, data *VmTemplate) error
		Update(ctx context.Context, session sqlx.Session, data *VmTemplate) (sql.Result, error)
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*VmTemplate, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*VmTemplate, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*VmTemplate, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*VmTemplate, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*VmTemplate, error)

		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultVmTemplateModel struct {
		conn  sqlx.SqlConn
		table string
	}

	VmTemplate struct {
		Id      int64  `db:"id"`
		Name    string `db:"name"`
		Release string `db:"release"`
		Enable  int64  `db:"enable"`
		Visible int64  `db:"visible"`
	}
)

func newVmTemplateModel(conn sqlx.SqlConn) *defaultVmTemplateModel {
	return &defaultVmTemplateModel{
		conn:  conn,
		table: "`vm_template`",
	}
}

func (m *defaultVmTemplateModel) withSession(session sqlx.Session) *defaultVmTemplateModel {
	return &defaultVmTemplateModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`vm_template`",
	}
}

func (m *defaultVmTemplateModel) Insert(ctx context.Context, session sqlx.Session, data *VmTemplate) (sql.Result, error) {

	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, vmTemplateRowsExpectAutoSet)
	if session != nil {
		return session.ExecCtx(ctx, query, data.Name, data.Release, data.Enable, data.Visible)
	}
	return m.conn.ExecCtx(ctx, query, data.Name, data.Release, data.Enable, data.Visible)
}

func (m *defaultVmTemplateModel) FindOne(ctx context.Context, id int64) (*VmTemplate, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", vmTemplateRows, m.table)
	var resp VmTemplate
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVmTemplateModel) FindOneByName(ctx context.Context, name string) (*VmTemplate, error) {
	var resp VmTemplate
	query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", vmTemplateRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVmTemplateModel) Update(ctx context.Context, session sqlx.Session, newData *VmTemplate) (sql.Result, error) {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, vmTemplateRowsWithPlaceHolder)
	if session != nil {
		return session.ExecCtx(ctx, query, newData.Name, newData.Release, newData.Enable, newData.Visible, newData.Id)
	}
	return m.conn.ExecCtx(ctx, query, newData.Name, newData.Release, newData.Enable, newData.Visible, newData.Id)
}

func (m *defaultVmTemplateModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindSum Least One Field"), "FindSum Least One Field")
	}

	builder = builder.Columns("IFNULL(SUM(" + field + "),0)")

	query, values, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64

	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultVmTemplateModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindCount Least One Field"), "FindCount Least One Field")
	}

	builder = builder.Columns("COUNT(" + field + ")")

	query, values, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64

	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultVmTemplateModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*VmTemplate, error) {

	builder = builder.Columns(vmTemplateRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*VmTemplate

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultVmTemplateModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*VmTemplate, error) {

	builder = builder.Columns(vmTemplateRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*VmTemplate

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultVmTemplateModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*VmTemplate, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(vmTemplateRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, total, err
	}

	var resp []*VmTemplate

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultVmTemplateModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*VmTemplate, error) {

	builder = builder.Columns(vmTemplateRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*VmTemplate

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultVmTemplateModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*VmTemplate, error) {

	builder = builder.Columns(vmTemplateRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*VmTemplate

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultVmTemplateModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

func (m *defaultVmTemplateModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}
func (m *defaultVmTemplateModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	if session != nil {
		_, err := session.ExecCtx(ctx, query, id)
		return err
	}
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultVmTemplateModel) tableName() string {
	return m.table
}