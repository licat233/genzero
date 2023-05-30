package conf

import (
	"strings"
)

var (
	BaseIgnoreTables  = []string{} //当前服务必须忽略的表
	BaseIgnoreColumns = []string{} //每个结构体必须忽略的列

	MoreIgnoreTables  = []string{}             //当前服务可能忽略的表
	MoreIgnoreColumns = []string{"id", "uuid"} //某个结构可能忽略的列
)

func init() {
	TplContent = strings.ReplaceAll(TplContent, "^", "`")
	TplContent = strings.Trim(TplContent, "\n")
}
