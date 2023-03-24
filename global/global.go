package global

import (
	"errors"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/sql"
)

var Schema *sql.Schema

func InitSchema() (err error) {
	sql.InitConfig()
	//优先使用ddl模式
	if config.C.DB.Src != "" {
		Schema, err = sql.ParseSqlFile(config.C.DB.Src)
		return
	}
	if config.C.DB.DSN != "" {
		Schema, err = sql.ParseSqlDsn(config.C.DB.DSN)
		return
	}

	return errors.New("failed to connect to the database. must set dsn or src")
}
