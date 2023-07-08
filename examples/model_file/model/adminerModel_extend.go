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

var AdminerTableName = "adminer"

type adminer_model interface {
	FindCount(ctx context.Context) (int64, error)
	FindAll(ctx context.Context) ([]*Adminer, error)
	FindList(ctx context.Context, pageSize, page int64, keyword string, adminer *Adminer) (resp []*Adminer, total int64, err error)
	TableName() string
	FindByUuid(ctx context.Context, uuid string) (*Adminer, error)
	FindByName(ctx context.Context, name string) (*Adminer, error)
	FindByAvatar(ctx context.Context, avatar string) (*Adminer, error)
	FindByPassport(ctx context.Context, passport string) (*Adminer, error)
	FindByPassword(ctx context.Context, password string) (*Adminer, error)
	FindByEmail(ctx context.Context, email string) (*Adminer, error)
	FindByStatus(ctx context.Context, status int64) (*Adminer, error)
	FindByIsSuperAdmin(ctx context.Context, isSuperAdmin int64) (*Adminer, error)
	FindByLoginCount(ctx context.Context, loginCount int64) (*Adminer, error)
	FindByLastLogin(ctx context.Context, lastLogin time.Time) (*Adminer, error)
	FindsByIds(ctx context.Context, ids []int64) ([]*Adminer, error)
	FindsByUuids(ctx context.Context, uuids []string) ([]*Adminer, error)
	FindsByNames(ctx context.Context, names []string) ([]*Adminer, error)
	FindsByAvatars(ctx context.Context, avatars []string) ([]*Adminer, error)
	FindsByPassports(ctx context.Context, passports []string) ([]*Adminer, error)
	FindsByPasswords(ctx context.Context, passwords []string) ([]*Adminer, error)
	FindsByEmails(ctx context.Context, emails []string) ([]*Adminer, error)
	FindsByStatuslist(ctx context.Context, statuslist []int64) ([]*Adminer, error)
	FindsByIsSuperAdmins(ctx context.Context, isSuperAdmins []int64) ([]*Adminer, error)
	FindsByLoginCounts(ctx context.Context, loginCounts []int64) ([]*Adminer, error)
	FindsByLastLogins(ctx context.Context, lastLogins []time.Time) ([]*Adminer, error)

	FindsById(ctx context.Context, id int64) ([]*Adminer, error)
	FindsByUuid(ctx context.Context, uuid string) ([]*Adminer, error)
	FindsByName(ctx context.Context, name string) ([]*Adminer, error)
	FindsByAvatar(ctx context.Context, avatar string) ([]*Adminer, error)
	FindsByPassport(ctx context.Context, passport string) ([]*Adminer, error)
	FindsByPassword(ctx context.Context, password string) ([]*Adminer, error)
	FindsByEmail(ctx context.Context, email string) ([]*Adminer, error)
	FindsByStatus(ctx context.Context, status int64) ([]*Adminer, error)
	FindsByIsSuperAdmin(ctx context.Context, isSuperAdmin int64) ([]*Adminer, error)
	FindsByLoginCount(ctx context.Context, loginCount int64) ([]*Adminer, error)
	FindsByLastLogin(ctx context.Context, lastLogin time.Time) ([]*Adminer, error)

	formatUuidKey(uuid string) string
	SoftDelete(ctx context.Context, id int64) error
}

func (m *defaultAdminerModel) FindCount(ctx context.Context) (int64, error) {
	var count int64
	query := fmt.Sprintf("select count(*) as count from %s where `is_deleted` = '0'", m.table)
	err := m.conn.QueryRowCtx(ctx, &count, query)
	return count, err
}

func (m *defaultAdminerModel) FindAll(ctx context.Context) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	query := fmt.Sprintf("select %s from %s where `is_deleted` = '0' limit 99999", adminerRows, m.table)
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

func (m *defaultAdminerModel) FindList(ctx context.Context, pageSize, page int64, keyword string, adminer *Adminer) (resp []*Adminer, total int64, err error) {
	hasName := false
	sq := squirrel.Select(adminerRows).From(m.table).Where("`is_deleted`= '0'")
	if adminer != nil {
		if adminer.Id >= 0 {
			sq = sq.Where("`id` = ?", adminer.Id)
		}
		if adminer.Uuid != "" {
			sq = sq.Where("`uuid` = ?", adminer.Uuid)
		}
		if adminer.Name != "" {
			sq = sq.Where("`name` = ?", adminer.Name)
			hasName = true
		}
		if adminer.Avatar != "" {
			sq = sq.Where("`avatar` = ?", adminer.Avatar)
		}
		if adminer.Passport != "" {
			sq = sq.Where("`passport` = ?", adminer.Passport)
		}
		if adminer.Password != "" {
			sq = sq.Where("`password` = ?", adminer.Password)
		}
		if adminer.Email != "" {
			sq = sq.Where("`email` = ?", adminer.Email)
		}
		if adminer.Status >= 0 {
			sq = sq.Where("`status` = ?", adminer.Status)
		}
		if adminer.IsSuperAdmin >= 0 {
			sq = sq.Where("`is_super_admin` = ?", adminer.IsSuperAdmin)
		}
		if adminer.LoginCount >= 0 {
			sq = sq.Where("`login_count` = ?", adminer.LoginCount)
		}
		if !adminer.LastLogin.IsZero() {
			sq = sq.Where("`last_login` = ?", adminer.LastLogin.Format("2006-01-02 15:04:05"))
		}
	}
	if keyword != "" && hasName {
		sq = sq.Where("`name` LIKE ?", fmt.Sprintf("%%%s%%", keyword))
	}
	if pageSize > 0 && page > 0 {
		sqCount := sq.RemoveLimit().RemoveOffset()
		sq = sq.Limit(uint64(pageSize)).Offset(uint64((page - 1) * pageSize))
		queryCount, agrsCount, e := sqCount.ToSql()
		if e != nil {
			err = e
			return
		}
		queryCount = strings.ReplaceAll(queryCount, adminerRows, "COUNT(*)")
		if err = m.conn.QueryRowCtx(ctx, &total, queryCount, agrsCount...); err != nil {
			return
		}
	}
	query, agrs, err := sq.ToSql()
	if err != nil {
		return
	}
	resp = make([]*Adminer, 0)
	if err = m.conn.QueryRowsCtx(ctx, &resp, query, agrs...); err != nil {
		return
	}
	return
}

func (m *defaultAdminerModel) TableName() string {
	return m.table
}

func (m *defaultAdminerModel) FindByUuid(ctx context.Context, uuid string) (*Adminer, error) {
	var resp Adminer
	query := fmt.Sprintf("select %s from %s where `uuid` = ? and `is_deleted` = '0' limit 1", adminerRows, m.table)
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

func (m *defaultAdminerModel) FindByName(ctx context.Context, name string) (*Adminer, error) {
	var resp Adminer
	query := fmt.Sprintf("select %s from %s where `name` = ? and `is_deleted` = '0' limit 1", adminerRows, m.table)
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

func (m *defaultAdminerModel) FindByAvatar(ctx context.Context, avatar string) (*Adminer, error) {
	var resp Adminer
	query := fmt.Sprintf("select %s from %s where `avatar` = ? and `is_deleted` = '0' limit 1", adminerRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, avatar)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAdminerModel) FindByPassport(ctx context.Context, passport string) (*Adminer, error) {
	var resp Adminer
	query := fmt.Sprintf("select %s from %s where `passport` = ? and `is_deleted` = '0' limit 1", adminerRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, passport)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAdminerModel) FindByPassword(ctx context.Context, password string) (*Adminer, error) {
	var resp Adminer
	query := fmt.Sprintf("select %s from %s where `password` = ? and `is_deleted` = '0' limit 1", adminerRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, password)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAdminerModel) FindByEmail(ctx context.Context, email string) (*Adminer, error) {
	var resp Adminer
	query := fmt.Sprintf("select %s from %s where `email` = ? and `is_deleted` = '0' limit 1", adminerRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, email)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAdminerModel) FindByStatus(ctx context.Context, status int64) (*Adminer, error) {
	var resp Adminer
	query := fmt.Sprintf("select %s from %s where `status` = ? and `is_deleted` = '0' limit 1", adminerRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, status)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAdminerModel) FindByIsSuperAdmin(ctx context.Context, isSuperAdmin int64) (*Adminer, error) {
	var resp Adminer
	query := fmt.Sprintf("select %s from %s where `is_super_admin` = ? and `is_deleted` = '0' limit 1", adminerRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, isSuperAdmin)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAdminerModel) FindByLoginCount(ctx context.Context, loginCount int64) (*Adminer, error) {
	var resp Adminer
	query := fmt.Sprintf("select %s from %s where `login_count` = ? and `is_deleted` = '0' limit 1", adminerRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, loginCount)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAdminerModel) FindByLastLogin(ctx context.Context, lastLogin time.Time) (*Adminer, error) {
	var resp Adminer
	query := fmt.Sprintf("select %s from %s where `last_login` = ? and `is_deleted` = '0' limit 1", adminerRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, lastLogin.Format("2006-01-02 15:04:05"))
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAdminerModel) FindsByIds(ctx context.Context, ids []int64) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	if len(ids) == 0 {
		return resp, nil
	}
	query := fmt.Sprintf("select %s from %s where `id` in (?) and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, ids)
	return resp, err
}

func (m *defaultAdminerModel) FindsByUuids(ctx context.Context, uuids []string) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	if len(uuids) == 0 {
		return resp, nil
	}
	query := fmt.Sprintf("select %s from %s where `uuid` in (?) and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, uuids)
	return resp, err
}

func (m *defaultAdminerModel) FindsByNames(ctx context.Context, names []string) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	if len(names) == 0 {
		return resp, nil
	}
	query := fmt.Sprintf("select %s from %s where `name` in (?) and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, names)
	return resp, err
}

func (m *defaultAdminerModel) FindsByAvatars(ctx context.Context, avatars []string) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	if len(avatars) == 0 {
		return resp, nil
	}
	query := fmt.Sprintf("select %s from %s where `avatar` in (?) and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, avatars)
	return resp, err
}

func (m *defaultAdminerModel) FindsByPassports(ctx context.Context, passports []string) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	if len(passports) == 0 {
		return resp, nil
	}
	query := fmt.Sprintf("select %s from %s where `passport` in (?) and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, passports)
	return resp, err
}

func (m *defaultAdminerModel) FindsByPasswords(ctx context.Context, passwords []string) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	if len(passwords) == 0 {
		return resp, nil
	}
	query := fmt.Sprintf("select %s from %s where `password` in (?) and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, passwords)
	return resp, err
}

func (m *defaultAdminerModel) FindsByEmails(ctx context.Context, emails []string) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	if len(emails) == 0 {
		return resp, nil
	}
	query := fmt.Sprintf("select %s from %s where `email` in (?) and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, emails)
	return resp, err
}

func (m *defaultAdminerModel) FindsByStatuslist(ctx context.Context, statuslist []int64) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	if len(statuslist) == 0 {
		return resp, nil
	}
	query := fmt.Sprintf("select %s from %s where `status` in (?) and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, statuslist)
	return resp, err
}

func (m *defaultAdminerModel) FindsByIsSuperAdmins(ctx context.Context, isSuperAdmins []int64) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	if len(isSuperAdmins) == 0 {
		return resp, nil
	}
	query := fmt.Sprintf("select %s from %s where `is_super_admin` in (?) and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, isSuperAdmins)
	return resp, err
}

func (m *defaultAdminerModel) FindsByLoginCounts(ctx context.Context, loginCounts []int64) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	if len(loginCounts) == 0 {
		return resp, nil
	}
	query := fmt.Sprintf("select %s from %s where `login_count` in (?) and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, loginCounts)
	return resp, err
}

func (m *defaultAdminerModel) FindsByLastLogins(ctx context.Context, lastLogins []time.Time) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	if len(lastLogins) == 0 {
		return resp, nil
	}
	query := fmt.Sprintf("select %s from %s where `last_login` in (?) and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, lastLogins)
	return resp, err
}

func (m *defaultAdminerModel) FindsById(ctx context.Context, id int64) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	query := fmt.Sprintf("select %s from %s where `id` = ? and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, id)
	return resp, err
}

func (m *defaultAdminerModel) FindsByUuid(ctx context.Context, uuid string) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	query := fmt.Sprintf("select %s from %s where `uuid` = ? and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, uuid)
	return resp, err
}

func (m *defaultAdminerModel) FindsByName(ctx context.Context, name string) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	query := fmt.Sprintf("select %s from %s where `name` = ? and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, name)
	return resp, err
}

func (m *defaultAdminerModel) FindsByAvatar(ctx context.Context, avatar string) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	query := fmt.Sprintf("select %s from %s where `avatar` = ? and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, avatar)
	return resp, err
}

func (m *defaultAdminerModel) FindsByPassport(ctx context.Context, passport string) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	query := fmt.Sprintf("select %s from %s where `passport` = ? and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, passport)
	return resp, err
}

func (m *defaultAdminerModel) FindsByPassword(ctx context.Context, password string) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	query := fmt.Sprintf("select %s from %s where `password` = ? and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, password)
	return resp, err
}

func (m *defaultAdminerModel) FindsByEmail(ctx context.Context, email string) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	query := fmt.Sprintf("select %s from %s where `email` = ? and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, email)
	return resp, err
}

func (m *defaultAdminerModel) FindsByStatus(ctx context.Context, status int64) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	query := fmt.Sprintf("select %s from %s where `status` = ? and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, status)
	return resp, err
}

func (m *defaultAdminerModel) FindsByIsSuperAdmin(ctx context.Context, isSuperAdmin int64) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	query := fmt.Sprintf("select %s from %s where `is_super_admin` = ? and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, isSuperAdmin)
	return resp, err
}

func (m *defaultAdminerModel) FindsByLoginCount(ctx context.Context, loginCount int64) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	query := fmt.Sprintf("select %s from %s where `login_count` = ? and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, loginCount)
	return resp, err
}

func (m *defaultAdminerModel) FindsByLastLogin(ctx context.Context, lastLogin time.Time) ([]*Adminer, error) {
	var resp = make([]*Adminer, 0)
	query := fmt.Sprintf("select %s from %s where `last_login` = ? and `is_deleted` = '0' ", adminerRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, lastLogin)
	return resp, err
}

func (m *defaultAdminerModel) formatUuidKey(uuid string) string {
	return fmt.Sprintf("cache:adminer:uuid:%v", uuid)
}

func (m *defaultAdminerModel) SoftDelete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("update %s set `is_deleted` = '1', `delete_at`= now() where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}
