package rpclogic

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/licat233/genzero/core/logic/conf"
	"github.com/licat233/genzero/tools"
)

func (l *Logic) getLogicFilename(method string, supply string) string {
	var filename string
	if l.Multiple {
		filename = path.Join(l.Dir, "base", fmt.Sprintf("base%s%s%slogic.go", method, strings.ToLower(l.CamelName), supply))
	} else {
		filename = path.Join(l.Dir, fmt.Sprintf("base%s%s%slogic.go", method, strings.ToLower(l.CamelName), supply))
	}
	return filename
}

func (l *Logic) getLogicFileContent(method string, supply string) (string, error) {
	filename := l.getLogicFilename(method, supply)
	exists, err := tools.PathExists(filename)
	if err != nil {
		return "", err
	}
	if !exists {
		tools.Warning("logic file not exists: %s", filename)
		return "", nil
	}
	body, err := tools.ReadFile(filename)
	return body, err
}

func modifyLogicFileContent(filename string, logicContent string, returnContent string) error {
	// 打开文件
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 创建带缓冲的读取器
	reader := bufio.NewReader(file)

	// 创建带缓冲的写入器
	writer := bufio.NewWriter(file)

	//用于标记 todo是否查找到
	findedTodo := false
	// 读取文件中的每一行，并进行操作
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		// 如果是需要替换的行，则替换为"hello"
		if strings.TrimSpace(line) == conf.TodoMark {
			line = logicContent
			findedTodo = true
		}

		if findedTodo && tools.IsReturn(line) {
			returnContent = tools.TrimSpace(returnContent)
			if returnContent != "" {
				if !strings.HasSuffix(returnContent, ",") {
					returnContent = returnContent + ","
				}
				if !strings.HasPrefix(returnContent, "\n") {
					returnContent = "\n" + returnContent
				}
				if !strings.HasSuffix(returnContent, "\n") {
					returnContent = returnContent + "\n"
				}
				line = tools.InsertString(line, "{", "}", returnContent)
			}
		}

		// 写入新的内容到文件
		_, err = writer.WriteString(line)
		if err != nil {
			panic(err)
		}
	}

	// 将缓冲区中的内容刷入文件
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}
