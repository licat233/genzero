package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ JwtBlacklistModel = (*customJwtBlacklistModel)(nil)

type (
	// JwtBlacklistModel is an interface to be customized, add more methods here,
	// and implement the added methods in customJwtBlacklistModel.
	JwtBlacklistModel interface {
		jwtBlacklistModel
	}

	customJwtBlacklistModel struct {
		*defaultJwtBlacklistModel
	}
)

// NewJwtBlacklistModel returns a model for the database table.
func NewJwtBlacklistModel(conn sqlx.SqlConn) JwtBlacklistModel {
	return &customJwtBlacklistModel{
		defaultJwtBlacklistModel: newJwtBlacklistModel(conn),
	}
}
