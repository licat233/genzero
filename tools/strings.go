package tools

import "strings"

func InsertString(str, startTag, endTag, insertStr string) string {
	start := strings.Index(str, startTag) + len(startTag)
	end := strings.Index(str, endTag)
	if start == -1 || end == -1 {
		return str
	}
	return str[:start] + insertStr + str[end:]
}
