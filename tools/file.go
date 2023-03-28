package tools

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

// 更新文件内容
func UpdateFileContent(filePath, oldContent, newContent string) error {
	// 读取文件内容
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// 替换旧内容为新内容
	newFileBytes := bytes.Replace(fileBytes, []byte(oldContent), []byte(newContent), -1)

	// 写入更新后的内容到原文件
	err = os.WriteFile(filePath, newFileBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

// 逐行检查文件内容，只处理一次
func CheckFileContent(filePath string, checkFunc func(string) bool) (bool, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// 逐行读取文件内容，并调用checkFunc函数
	for scanner.Scan() {
		line := scanner.Text()
		if checkFunc(line) {
			return true, nil
		}
	}

	if scanner.Err() != nil {
		return false, scanner.Err()
	}

	return false, nil
}

func UpdateFileByLine(filePath string, updateFunc func(string) string) error {
	// 打开文件
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	// 逐行读取文件内容，并调用updateFunc函数进行修改
	for scanner.Scan() {
		line := scanner.Text()
		newLine := updateFunc(line)
		lines = append(lines, newLine)
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	// 清空文件内容
	err = file.Truncate(0)
	if err != nil {
		return err
	}

	// 将更新后的内容写入文件中
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, line := range lines {
		_, err := writer.WriteString(fmt.Sprintf("%s\n", line))
		if err != nil {
			return err
		}
	}
	return nil
}
