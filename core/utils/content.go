package utils

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

var Indent = "  " //two space

func HasMark(startMark, endMark, fileContent string) (bool, error) {
	if strings.TrimSpace(fileContent) == "" {
		return false, nil
	}
	startMark = regexp.QuoteMeta(startMark)
	endMark = regexp.QuoteMeta(endMark)

	expr := fmt.Sprintf("%s[\n]*((?s).*?)[\n]*%s", startMark, endMark)

	reg, err := regexp.Compile(expr)
	if err != nil {
		return false, err
	}
	ok := reg.MatchString(fileContent)
	return ok, nil
}

func UpdateMarkContent(startMark, endMark, fileContent string, content string) string {
	start := strings.Index(fileContent, startMark)
	end := strings.Index(fileContent, endMark)
	if start == -1 || end == -1 {
		res := fmt.Sprintf("%s\n %s", fileContent, GenBaseBlock(startMark, endMark, content, ""))
		return res
	}
	content = strings.Trim(content, "\n")
	res := fmt.Sprintf("%s\n\n%s\n\n%s", fileContent[:start+len(startMark)], content, fileContent[end:])
	return res
}

func PickMarkContents(startMark, endMark, oldContent string) ([]string, error) {
	content := strings.TrimSpace(oldContent)
	if content == "" {
		return []string{}, nil
	}

	startMark = regexp.QuoteMeta(startMark)
	endMark = regexp.QuoteMeta(endMark)

	expr := fmt.Sprintf("%s[\n]*((?s).*?)[\n]*%s", startMark, endMark)

	reg, err := regexp.Compile(expr)
	if err != nil {
		return nil, err
	}
	listArr := reg.FindAllStringSubmatch(content, -1)
	strArr := []string{}
	for _, list := range listArr {
		if len(list) != 2 {
			continue
		}
		target := strings.TrimSpace(list[len(list)-1])
		if target == "" {
			continue
		}
		target = strings.Trim(target, "")
		target = strings.Trim(target, "\n")
		target = fmt.Sprintf("\n%s\n", target)
		strArr = append(strArr, target)
	}
	return strArr, nil
}

func GetMarkContent(startMark, endMark, oldContent string) (string, error) {
	if oldContent == "" {
		return "", nil
	}
	list, err := PickMarkContents(startMark, endMark, oldContent)
	if err != nil {
		return "", err
	}
	customContent := strings.Join(list, "\n")
	customContent = strings.Trim(customContent, "\n")
	return customContent, nil
}

func GenCustomBlock(startMark, endMark, content, indent string) string {
	var buf = new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("\n\n%s// The content in this block will not be updated\n%s// 此区块内的内容不会被更新", indent, indent))
	buf.WriteString(fmt.Sprintf("\n%s%s\n", indent, startMark))
	buf.WriteString(fmt.Sprintf("\n%s\n", content))
	buf.WriteString(fmt.Sprintf("\n%s%s\n", indent, endMark))
	return buf.String()
}

func GenBaseBlock(startMark, endMark, content, indent string) string {
	var buf = new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("\n%s%s\n", indent, startMark))
	buf.WriteString(fmt.Sprintf("\n%s\n", content))
	buf.WriteString(fmt.Sprintf("\n%s%s\n", indent, endMark))
	return buf.String()
}

func PickInfoContent(content string) string {
	re := regexp.MustCompile(`(?s)info\s*\((.*?)\n\s*\)`)
	match := re.FindStringSubmatch(content)

	if len(match) == 2 {
		res := match[1]
		str := strings.TrimSpace(res)
		str = strings.Trim(str, "\n")
		if str == "" {
			return ""
		}
		return res
	} else {
		return ""
	}
}

// FormatContent 格式化内容
func FormatContent(str string) string {
	// 拆分成行,
	lines := strings.Split(str, "\n")
	isFront := true
	newLines := []string{}
	// 去除多余空格
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			if isFront {
				continue
			}
		}
		isFront = false

		newLines = append(newLines, line)
	}

	var IndentMul = func(num int) string {
		res := ""
		for i := 0; i < num; i++ {
			res += Indent
		}
		return res
	}

	var needIndentNum int
	// 处理{}内的缩进
	for index, line := range newLines {
		if len(line) == 0 {
			continue
		}
		if line[len(line)-1] == '{' {
			needIndentNum++
			continue
		}
		if strings.TrimSpace(line)[0] == '}' {
			if needIndentNum > 0 {
				needIndentNum--
			}
			continue
		}
		if needIndentNum > 0 {
			line = IndentMul(needIndentNum) + line
		}
		newLines[index] = line
	}

	needIndentNum = 0
	// 处理()内的缩进
	for index, line := range newLines {
		if len(line) == 0 {
			continue
		}
		if line[len(line)-1] == '(' {
			needIndentNum++
			continue
		}
		if strings.TrimSpace(line)[0] == ')' {
			if needIndentNum > 0 {
				needIndentNum--
			}
			continue
		}
		if needIndentNum > 0 {
			line = IndentMul(needIndentNum) + line
		}

		newLines[index] = line
	}

	newContent := strings.Join(newLines, "\n")

	re := regexp.MustCompile(`\n{2,}`)
	newContent = re.ReplaceAllString(newContent, "\n\n")
	return newContent
}
