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
	orderItemFieldNames          = builder.RawFieldNames(&OrderItem{})
	orderItemRows                = strings.Join(orderItemFieldNames, ",")
	orderItemRowsExpectAutoSet   = strings.Join(stringx.Remove(orderItemFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	orderItemRowsWithPlaceHolder = strings.Join(stringx.Remove(orderItemFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	orderItemModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *OrderItem) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*OrderItem, error)
		// Update(ctx context.Context, data *OrderItem) error
		Update(ctx context.Context, session sqlx.Session, data *OrderItem) (sql.Result, error)
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*OrderItem, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*OrderItem, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*OrderItem, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*OrderItem, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*OrderItem, error)

		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultOrderItemModel struct {
		conn  sqlx.SqlConn
		table string
	}

	OrderItem struct {
		Id         int64  `db:"id"`
		OrderId    int64  `db:"order_id"`
		ActionType string `db:"action_type"` // 订单类型: VM开通、充值等：'vm_instance_create', 'addFunds'
		Action     string `db:"action"`
		Name       string `db:"name"`
		Content    string `db:"content"`
		Quantity   int64  `db:"quantity"`
		Amount     int64  `db:"amount"`
	}
)

func newOrderItemModel(conn sqlx.SqlConn) *defaultOrderItemModel {
	return &defaultOrderItemModel{
		conn:  conn,
		table: "`order_item`",
	}
}

func (m *defaultOrderItemModel) withSession(session sqlx.Session) *defaultOrderItemModel {
	return &defaultOrderItemModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`order_item`",
	}
}

func (m *defaultOrderItemModel) Insert(ctx context.Context, session sqlx.Session, data *OrderItem) (sql.Result, error) {

	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, orderItemRowsExpectAutoSet)
	if session != nil {
		return session.ExecCtx(ctx, query, data.OrderId, data.ActionType, data.Action, data.Name, data.Content, data.Quantity, data.Amount)
	}
	return m.conn.ExecCtx(ctx, query, data.OrderId, data.ActionType, data.Action, data.Name, data.Content, data.Quantity, data.Amount)
}

func (m *defaultOrderItemModel) FindOne(ctx context.Context, id int64) (*OrderItem, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderItemRows, m.table)
	var resp OrderItem
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

func (m *defaultOrderItemModel) Update(ctx context.Context, session sqlx.Session, data *OrderItem) (sql.Result, error) {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderItemRowsWithPlaceHolder)
	if session != nil {
		return session.ExecCtx(ctx, query, data.OrderId, data.ActionType, data.Action, data.Name, data.Content, data.Quantity, data.Amount, data.Id)
	}
	return m.conn.ExecCtx(ctx, query, data.OrderId, data.ActionType, data.Action, data.Name, data.Content, data.Quantity, data.Amount, data.Id)
}

func (m *defaultOrderItemModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

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

func (m *defaultOrderItemModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

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

func (m *defaultOrderItemModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*OrderItem, error) {

	builder = builder.Columns(orderItemRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*OrderItem

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultOrderItemModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*OrderItem, error) {

	builder = builder.Columns(orderItemRows)

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

	var resp []*OrderItem

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultOrderItemModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*OrderItem, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(orderItemRows)

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

	var resp []*OrderItem

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultOrderItemModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*OrderItem, error) {

	builder = builder.Columns(orderItemRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*OrderItem

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultOrderItemModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*OrderItem, error) {

	builder = builder.Columns(orderItemRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*OrderItem

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultOrderItemModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

func (m *defaultOrderItemModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}
func (m *defaultOrderItemModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	if session != nil {
		_, err := session.ExecCtx(ctx, query, id)
		return err
	}
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultOrderItemModel) tableName() string {
	return m.table
}
