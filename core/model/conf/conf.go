package conf

var (
	BaseIgnoreTables  = []string{} //当前服务必须忽略的表
	BaseIgnoreColumns = []string{} //每个结构体必须忽略的列

	MoreIgnoreTables  = []string{}             //当前服务可能忽略的表
	MoreIgnoreColumns = []string{"id", "uuid"} //某个结构可能忽略的列
	IsCacheMode       bool
)

// func ChangeQueryString(isCache bool) {
// 	IsCacheMode = isCache
// 	if isCache {
// 		QueryRow = "m.QueryRowNoCacheCtx"
// 		QueryRows = "m.QueryRowsNoCacheCtx"
// 		Exec = "m.ExecNoCacheCtx"
// 	} else {
// 		QueryRow = "m.conn.QueryRowCtx"
// 		QueryRows = "m.conn.QueryRowsCtx"
// 		Exec = "m.conn.ExecCtx"
// 	}
// }
