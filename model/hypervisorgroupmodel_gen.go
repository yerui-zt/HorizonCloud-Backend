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
	hypervisorGroupFieldNames          = builder.RawFieldNames(&HypervisorGroup{})
	hypervisorGroupRows                = strings.Join(hypervisorGroupFieldNames, ",")
	hypervisorGroupRowsExpectAutoSet   = strings.Join(stringx.Remove(hypervisorGroupFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	hypervisorGroupRowsWithPlaceHolder = strings.Join(stringx.Remove(hypervisorGroupFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	hypervisorGroupModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *HypervisorGroup) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*HypervisorGroup, error)
		FindOneByName(ctx context.Context, name string) (*HypervisorGroup, error)
		// Update(ctx context.Context, data *HypervisorGroup) error
		Update(ctx context.Context, session sqlx.Session, data *HypervisorGroup) (sql.Result, error)
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*HypervisorGroup, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HypervisorGroup, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HypervisorGroup, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*HypervisorGroup, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*HypervisorGroup, error)

		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultHypervisorGroupModel struct {
		conn  sqlx.SqlConn
		table string
	}

	HypervisorGroup struct {
		Id      int64  `db:"id"`
		Name    string `db:"name"`
		Region  string `db:"region"`
		Enable  int64  `db:"enable"`
		Visible int64  `db:"visible"`
	}
)

func newHypervisorGroupModel(conn sqlx.SqlConn) *defaultHypervisorGroupModel {
	return &defaultHypervisorGroupModel{
		conn:  conn,
		table: "`hypervisor_group`",
	}
}

func (m *defaultHypervisorGroupModel) withSession(session sqlx.Session) *defaultHypervisorGroupModel {
	return &defaultHypervisorGroupModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`hypervisor_group`",
	}
}

func (m *defaultHypervisorGroupModel) Insert(ctx context.Context, session sqlx.Session, data *HypervisorGroup) (sql.Result, error) {

	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, hypervisorGroupRowsExpectAutoSet)
	if session != nil {
		return session.ExecCtx(ctx, query, data.Name, data.Region, data.Enable, data.Visible)
	}
	return m.conn.ExecCtx(ctx, query, data.Name, data.Region, data.Enable, data.Visible)
}

func (m *defaultHypervisorGroupModel) FindOne(ctx context.Context, id int64) (*HypervisorGroup, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", hypervisorGroupRows, m.table)
	var resp HypervisorGroup
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

func (m *defaultHypervisorGroupModel) FindOneByName(ctx context.Context, name string) (*HypervisorGroup, error) {
	var resp HypervisorGroup
	query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", hypervisorGroupRows, m.table)
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

func (m *defaultHypervisorGroupModel) Update(ctx context.Context, session sqlx.Session, newData *HypervisorGroup) (sql.Result, error) {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, hypervisorGroupRowsWithPlaceHolder)
	if session != nil {
		return session.ExecCtx(ctx, query, newData.Name, newData.Region, newData.Enable, newData.Visible, newData.Id)
	}
	return m.conn.ExecCtx(ctx, query, newData.Name, newData.Region, newData.Enable, newData.Visible, newData.Id)
}

func (m *defaultHypervisorGroupModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

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

func (m *defaultHypervisorGroupModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

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

func (m *defaultHypervisorGroupModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*HypervisorGroup, error) {

	builder = builder.Columns(hypervisorGroupRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HypervisorGroup

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHypervisorGroupModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HypervisorGroup, error) {

	builder = builder.Columns(hypervisorGroupRows)

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

	var resp []*HypervisorGroup

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHypervisorGroupModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HypervisorGroup, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(hypervisorGroupRows)

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

	var resp []*HypervisorGroup

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultHypervisorGroupModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*HypervisorGroup, error) {

	builder = builder.Columns(hypervisorGroupRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HypervisorGroup

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHypervisorGroupModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*HypervisorGroup, error) {

	builder = builder.Columns(hypervisorGroupRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HypervisorGroup

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHypervisorGroupModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

func (m *defaultHypervisorGroupModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}
func (m *defaultHypervisorGroupModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	if session != nil {
		_, err := session.ExecCtx(ctx, query, id)
		return err
	}
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultHypervisorGroupModel) tableName() string {
	return m.table
}