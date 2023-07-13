package utils

import (
	"os"
	"strings"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/tools"
)

// 备份用户自己的文件
// 如果文件存在，且不是genzero生成的文件，则备份起来，防止被覆盖，丢失数据
func BackupUserFile(filename string) error {
	//判断文件是否存在
	exist, err := tools.FileExists(filename)
	if err != nil {
		return err
	}
	if !exist {
		return nil
	}
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	keyword := config.ProjectURL
	exist = strings.Contains(string(content), keyword)
	if exist {
		return nil
	}
	//如果不存在关键词，表示是用户自定义的文件，需要备份保护
	return tools.BackupFile(filename, "backup")
}
