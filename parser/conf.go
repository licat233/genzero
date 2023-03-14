package parser

import "github.com/licat233/genzero/config"

var (
	baseIgnoreTables  = []string{}
	baseIgnoreColumns = []string{
		"version",
		"create_time",
		"created_time",
		"create_at",
		"created_at",
		"update_time",
		"updated_time",
		"update_at",
		"updated_at",
		"delete_time",
		"deleted_time",
		"delete_at",
		"deleted_at",
		"del_state",
		"is_deleted",
		"is_delete",
	}
)

func InitConfig() {
	config.C.DatabaseConfig.IgnoreTables = append(config.C.DatabaseConfig.IgnoreTables, baseIgnoreTables...)
	config.C.DatabaseConfig.IgnoreColumns = append(config.C.DatabaseConfig.IgnoreColumns, baseIgnoreColumns...)
}
