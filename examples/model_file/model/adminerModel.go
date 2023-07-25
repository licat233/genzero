package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AdminerModel = (*customAdminerModel)(nil)

type (
	// AdminerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminerModel.
	AdminerModel interface {
		adminerModel  // extended interface by gozero
		adminer_model // extended interface by genzero
	}

	customAdminerModel struct {
		*defaultAdminerModel
	}
)

// NewAdminerModel returns a model for the database table.
func NewAdminerModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AdminerModel {
	cacheAdminerIdPrefix = "cache:admin:adminer:id:" // modifying cache id prefix by genzero
	return &customAdminerModel{
		defaultAdminerModel: newAdminerModel(conn, c, opts...),
	}
}
