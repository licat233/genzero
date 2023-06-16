package sql

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
	}
	DelFieldNames = []string{
		"is_deleted",
		"is_delete",
		"is_del",
		"deleted",
		"delete_flag",
		"del_flag",
		"del_state",
		"delete_state",
		"deleted_state",
	}
	DelAtFieldNames = []string{
		"deleted_at",
		"delete_at",
		"deleted_time",
		"delete_time",
		"del_at",
		"del_time",
	}
	NameFieldNames = []string{
		"name",
		"username",
		"nickname",
	}
)

func InitConfig() {
	// baseIgnoreColumns = append(baseIgnoreColumns, DelFieldNames...)
	// baseIgnoreColumns = append(baseIgnoreColumns, DelAtFieldNames...)
	config.C.DB.IgnoreColumns = append(config.C.DB.IgnoreColumns, append(append(baseIgnoreColumns, DelAtFieldNames...), DelFieldNames...)...)

	config.C.DB.IgnoreTables = append(config.C.DB.IgnoreTables, baseIgnoreTables...)

}
