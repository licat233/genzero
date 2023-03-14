package conf

var (
	FileContent string //文件内容

	BaseIgnoreTables  = []string{} //当前服务必须忽略的表
	BaseIgnoreColumns = []string{} //每个结构体必须忽略的列

	MoreIgnoreTables  = []string{}             //当前服务可能忽略的表
	MoreIgnoreColumns = []string{"id", "uuid"} //某个结构可能忽略的列

	QueryRow  = "m.conn.QueryRowCtx"
	QueryRows = "m.conn.QueryRowsCtx"
)

func ChangeQueryString(isCache bool) {
	if isCache {
		QueryRow = "m.QueryRowNoCacheCtx"
		QueryRows = "m.QueryRowsNoCacheCtx"
	} else {
		QueryRow = "m.conn.QueryRowCtx"
		QueryRows = "m.conn.QueryRowsCtx"
	}
}
