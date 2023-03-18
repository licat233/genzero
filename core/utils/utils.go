package utils

import (
	"strings"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/tools"
)

// ConvertStringStyle 转化字符串风格，默认 snake 风格
func ConvertStringStyle(style, name string) string {
	switch style {
	case config.CamelCase:
		return tools.ToCamel(name)
	case config.LowerCamelCase:
		return tools.ToLowerCamel(name)
	case config.SnakeCase:
		return tools.ToSnake(name)
	default:
		return tools.ToSnake(name)
	}
}

func HandleOptContent(opts ...string) string {
	optstring := strings.Join(opts, ",")
	list := strings.Split(optstring, ",")
	filter := []string{}
	for _, arg := range list {
		arg = strings.TrimSpace(arg)
		if arg == "" {
			continue
		}
		filter = append(filter, arg)
	}
	return strings.Join(filter, ",")
}

func ToCamelHandler(value string) string {
	if value == "" {
		return ""
	}
	list := strings.Split(value, ",")
	for i, v := range list {
		list[i] = tools.ToCamel(v)
	}
	return strings.Join(list, ",")
}
