package global

import (
	"errors"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/parser"
)

var Schema *parser.Schema

func InitSchema() (err error) {
	parser.InitConfig()
	//优先使用ddl模式
	if config.C.DatabaseConfig.Src != "" {
		Schema, err = parser.ParseSqlFile(config.C.DatabaseConfig.Src)
		return
	}
	if config.C.DatabaseConfig.DSN != "" {
		Schema, err = parser.ParseSqlDsn(config.C.DatabaseConfig.DSN)
		return
	}

	return errors.New("failed to connect to the database. must set dsn or src")
}
