// Code generated by genzero. DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

var JwtBlacklistTableName = "jwt_blacklist"

type jwtBlacklist_model interface {
	FindAll(ctx context.Context) ([]*JwtBlacklist, error)
	FindList(ctx context.Context, pageSize, page int64, keyword string, jwtBlacklist *JwtBlacklist) (resp []*JwtBlacklist, total int64, err error)
	FindsByIds(ctx context.Context, ids []int64) ([]*JwtBlacklist, error)
	TableName() string
	FindByAdminerId(ctx context.Context, adminerId int64) (*JwtBlacklist, error)
	FindByUuid(ctx context.Context, uuid string) (*JwtBlacklist, error)
	FindByToken(ctx context.Context, token string) (*JwtBlacklist, error)
	FindByPlatform(ctx context.Context, platform string) (*JwtBlacklist, error)
	FindByIp(ctx context.Context, ip string) (*JwtBlacklist, error)
	FindByExpireAt(ctx context.Context, expireAt time.Time) (*JwtBlacklist, error)
	formatUuidKey(uuid string) string
}

func (m *defaultJwtBlacklistModel) FindAll(ctx context.Context) ([]*JwtBlacklist, error) {
	var resp = make([]*JwtBlacklist, 0)
	query := fmt.Sprintf("select %s from %s limit 99999", jwtBlacklistRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultJwtBlacklistModel) FindList(ctx context.Context, pageSize, page int64, keyword string, jwtBlacklist *JwtBlacklist) (resp []*JwtBlacklist, total int64, err error) {
	sq := squirrel.Select(jwtBlacklistRows).From(m.table)
	if jwtBlacklist != nil {
		if jwtBlacklist.Id >= 0 {
			sq = sq.Where("id = ?", jwtBlacklist.Id)
		}
		if jwtBlacklist.AdminerId > 0 {
			sq = sq.Where("adminer_id = ?", jwtBlacklist.AdminerId)
		}
		if jwtBlacklist.Uuid != "" {
			sq = sq.Where("uuid = ?", jwtBlacklist.Uuid)
		}
		if jwtBlacklist.Token != "" {
			sq = sq.Where("token = ?", jwtBlacklist.Token)
		}
		if jwtBlacklist.Platform != "" {
			sq = sq.Where("platform = ?", jwtBlacklist.Platform)
		}
		if jwtBlacklist.Ip != "" {
			sq = sq.Where("ip = ?", jwtBlacklist.Ip)
		}
		if jwtBlacklist.ExpireAt.IsZero() {
			sq = sq.Where("expire_at = ?", jwtBlacklist.ExpireAt.Format("2006-01-02 15:04:05"))
		}
	}
	if pageSize > 0 && page > 0 {
		sqCount := sq.RemoveLimit().RemoveOffset()
		sq = sq.Limit(uint64(pageSize)).Offset(uint64((page - 1) * pageSize))
		queryCount, agrsCount, e := sqCount.ToSql()
		if e != nil {
			err = e
			return
		}
		queryCount = strings.ReplaceAll(queryCount, jwtBlacklistRows, "COUNT(*)")
		if err = m.conn.QueryRowCtx(ctx, &total, queryCount, agrsCount...); err != nil {
			return
		}
	}
	query, agrs, err := sq.ToSql()
	if err != nil {
		return
	}
	resp = make([]*JwtBlacklist, 0)
	if err = m.conn.QueryRowsCtx(ctx, &resp, query, agrs...); err != nil {
		return
	}
	return
}

func (m *defaultJwtBlacklistModel) FindsByIds(ctx context.Context, ids []int64) ([]*JwtBlacklist, error) {
	var resp = make([]*JwtBlacklist, 0)
	if len(ids) == 0 {
		return resp, nil
	}
	query := fmt.Sprintf("select %s from %s where `id` in(?) ", jwtBlacklistRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, ids)
	return resp, err
}

func (m *defaultJwtBlacklistModel) TableName() string {
	return m.table
}

func (m *defaultJwtBlacklistModel) FindByAdminerId(ctx context.Context, adminerId int64) (*JwtBlacklist, error) {
	var resp JwtBlacklist
	query := fmt.Sprintf("select %s from %s where `adminer_id` = ? limit 1", jwtBlacklistRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, adminerId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultJwtBlacklistModel) FindByUuid(ctx context.Context, uuid string) (*JwtBlacklist, error) {
	var resp JwtBlacklist
	query := fmt.Sprintf("select %s from %s where `uuid` = ? limit 1", jwtBlacklistRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, uuid)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultJwtBlacklistModel) FindByToken(ctx context.Context, token string) (*JwtBlacklist, error) {
	var resp JwtBlacklist
	query := fmt.Sprintf("select %s from %s where `token` = ? limit 1", jwtBlacklistRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, token)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultJwtBlacklistModel) FindByPlatform(ctx context.Context, platform string) (*JwtBlacklist, error) {
	var resp JwtBlacklist
	query := fmt.Sprintf("select %s from %s where `platform` = ? limit 1", jwtBlacklistRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, platform)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultJwtBlacklistModel) FindByIp(ctx context.Context, ip string) (*JwtBlacklist, error) {
	var resp JwtBlacklist
	query := fmt.Sprintf("select %s from %s where `ip` = ? limit 1", jwtBlacklistRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, ip)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultJwtBlacklistModel) FindByExpireAt(ctx context.Context, expireAt time.Time) (*JwtBlacklist, error) {
	var resp JwtBlacklist
	query := fmt.Sprintf("select %s from %s where `expire_at` = ? limit 1", jwtBlacklistRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, expireAt.Format("2006-01-02 15:04:05"))
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultJwtBlacklistModel) formatUuidKey(uuid string) string {
	return fmt.Sprintf("cache:jwtBlacklist:uuid:%v", uuid)
}
