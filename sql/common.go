package sql

import (
	"github.com/licat233/genzero/config"
)

func TableFilterCondition() (mustFilter bool, checkTables []string) {
	if len(config.C.DB.Tables) == 0 {
		return false, nil
	}
	tables := []string{}
	for _, t := range config.C.DB.Tables {
		if t == "" {
			continue
		}
		tables = append(tables, t)
	}
	if len(tables) == 1 {
		if tables[0] == "*" {
			return false, nil
		}
	}
	return len(tables) != 0, tables
}

func goType(sqlType string) string {
	v, ok := TypeForMysqlToGo[sqlType]
	if ok {
		return v
	}
	return "any"
}
