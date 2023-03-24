package tools

import (
	"regexp"
	"strings"
)

func InsertString(str, startTag, endTag, insertStr string) string {
	start := strings.Index(str, startTag) + len(startTag)
	end := strings.Index(str, endTag)
	if start == -1 || end == -1 {
		return str
	}
	return str[:start] + insertStr + str[end:]
}

func ContainsString(str, substr string) bool {
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == substr {
			return true
		}
	}
	return false
}

func IsFunction(str string) bool {
	str = strings.TrimSpace(str)
	reg := `^func\s+\(.+\)\s+(\(?[^()]*?\)?)\s+{(\s*\\\\.*)?`
	return regexp.MustCompile(reg).MatchString(str)
}

func FindFuncResp(str string) string {
	str = strings.TrimSpace(str)
	reg := `^func\s+\(.+\)\s+(\(?[^()]*?\)?)\s+{(\s*\\\\.*)?`
	rgp := regexp.MustCompile(reg)
	res := rgp.FindStringSubmatch(str)
	if res == nil {
		return ""
	}
	return res[1]
}

func ReplaceFuncResp(str string, repl string) string {
	resp := FindFuncResp(str)
	if resp == "" {
		return str
	}
	res := strings.Replace(str, resp, repl, -1)
	return res
}
