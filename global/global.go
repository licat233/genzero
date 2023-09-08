package global

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/sql"
)

var Schema *sql.Schema

func InitSchema() (err error) {
	sql.InitConfig()
	//优先使用ddl模式
	if config.C.DB.Src != "" {
		absPath, e := filepath.Abs(config.C.DB.Src)
		if e != nil {
			err = fmt.Errorf("无法获取绝对路径:%s", e)
			return
		}
		Schema, err = sql.ParseSqlFile(absPath)
		return
	} else if config.C.DB.DSN != "" {
		Schema, err = sql.ParseSqlDsn(config.C.DB.DSN)
		return
	} else {
		return errors.New("failed to connect to the database. must set dsn or src")
	}
}
