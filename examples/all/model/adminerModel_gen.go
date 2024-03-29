// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	adminerFieldNames          = builder.RawFieldNames(&Adminer{})
	adminerRows                = strings.Join(adminerFieldNames, ",")
	adminerRowsExpectAutoSet   = strings.Join(stringx.Remove(adminerFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	adminerRowsWithPlaceHolder = strings.Join(stringx.Remove(adminerFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	adminerModel interface {
		Insert(ctx context.Context, data *Adminer) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Adminer, error)
		Update(ctx context.Context, data *Adminer) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAdminerModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Adminer struct {
		Id           int64          `db:"id"`             // 表主键
		Uuid         string         `db:"uuid"`           // 唯一识别码
		Name         string         `db:"name"`           // 管理员名称
		Avatar       string         `db:"avatar"`         // 头像
		Passport     string         `db:"passport"`       // 账号
		Password     string         `db:"password"`       // 密码
		Email        string         `db:"email"`          // 邮箱
		Resume       sql.NullString `db:"resume"`         // 个人简介
		Status       int64          `db:"status"`         // 账号状态，是否可用
		IsSuperAdmin int64          `db:"is_super_admin"` // 是否为超级管理员
		LoginCount   int64          `db:"login_count"`    // 登录次数
		LastLogin    time.Time      `db:"last_login"`     // 最后一次登录时间
		CreateAt     time.Time      `db:"create_at"`      // 创建时间
		UpdateAt     time.Time      `db:"update_at"`      // 更新时间
		IsDeleted    int64          `db:"is_deleted"`     // 是否已被删除
		DeleteAt     time.Time      `db:"delete_at"`      // 删除时间
	}
)

func newAdminerModel(conn sqlx.SqlConn) *defaultAdminerModel {
	return &defaultAdminerModel{
		conn:  conn,
		table: "`adminer`",
	}
}

func (m *defaultAdminerModel) withSession(session sqlx.Session) *defaultAdminerModel {
	return &defaultAdminerModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`adminer`",
	}
}

func (m *defaultAdminerModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultAdminerModel) FindOne(ctx context.Context, id int64) (*Adminer, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", adminerRows, m.table)
	var resp Adminer
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

func (m *defaultAdminerModel) Insert(ctx context.Context, data *Adminer) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, adminerRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Uuid, data.Name, data.Avatar, data.Passport, data.Password, data.Email, data.Resume, data.Status, data.IsSuperAdmin, data.LoginCount, data.LastLogin, data.IsDeleted, data.DeleteAt)
	return ret, err
}

func (m *defaultAdminerModel) Update(ctx context.Context, data *Adminer) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, adminerRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Uuid, data.Name, data.Avatar, data.Passport, data.Password, data.Email, data.Resume, data.Status, data.IsSuperAdmin, data.LoginCount, data.LastLogin, data.IsDeleted, data.DeleteAt, data.Id)
	return err
}

func (m *defaultAdminerModel) tableName() string {
	return m.table
}
