package cmd

import "fmt"

var IsDev bool

var colorId int = 30

func getCmdColor() int {
	colorId += 1
	if colorId < 30 || (colorId > 37 && colorId < 90) || colorId > 97 {
		return 30
	}
	return colorId
}

func setColorizeHelp(template string) string {
	// 在这里使用ANSI转义码添加颜色
	// 参考ANSI转义码文档：https://en.wikipedia.org/wiki/ANSI_escape_code
	return fmt.Sprintf("\033[%dm%s\033[0m", getCmdColor(), template)
}

func setColorPart(part string) string {
	return fmt.Sprintf("\033[%dm%s", getCmdColor(), part)
}
