package tools

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
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

// 备份文件
func BackupFile(filename string, backupDir string) error {
	timestamp := time.Now().Format("20060102150405")
	backupFilename := fmt.Sprintf("%s_%s", timestamp, filename)
	backupPath := filepath.Join(backupDir, backupFilename)

	err := os.MkdirAll(backupDir, 0755)
	if err != nil {
		return err
	}

	return os.Rename(filename, backupPath)
}

func RemoveComment(line string) string {
	index := strings.Index(line, "//")
	if index != -1 {
		line = line[:index]
	}
	return strings.TrimSpace(line)
}

func InsertFieldsToSvcFile(filePath string, targets []string) error {
	has, err := PathExists(filePath)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("[logic] Initialization failed: the (%s) file does not exist. Please use the goctl tool to create it first.\n - goctl: https://go-zero.dev/cn/docs/goctl/goctl/", filePath)
	}
	// 打开文件以供读取
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	patternStart := regexp.MustCompile("^type ServiceContext struct {$")
	patternEnd := regexp.MustCompile("^}$")
	isStart := false
	isEnd := false
	isModify := false
	targetsMap := make(map[string]bool)
	for _, target := range targets {
		targetsMap[target] = false
	}

	var findTarget = func(line string) {
		line = RemoveComment(line)
		for target, exist := range targetsMap {
			if exist {
				continue
			}
			if strings.HasSuffix(line, target) {
				//标记为已找到
				targetsMap[target] = true
				break
			}
		}
	}

	var newContent = new(bytes.Buffer)
	// 使用bufio.Scanner获取文件中每一行的内容
	scanner := bufio.NewScanner(file)
	// 逐行扫描和打印
	for scanner.Scan() {
		line := scanner.Text()
		if isEnd {
			newContent.WriteString(line + "\n")
			continue
		}
		lineStr := strings.TrimSpace(line)
		if isStart {
			if patternEnd.MatchString(lineStr) {
				isEnd = true
				//结束
				//开始写入不存在的结构体字段
				var insertStr string
				var isInsert bool
				for target, exist := range targetsMap {
					if !exist {
						isInsert = true
						isModify = true
						insertStr = fmt.Sprintf("%s\t%s\n", insertStr, target)
					}
				}
				//存在需要写入的字段
				if isInsert {
					line = fmt.Sprintf("%s%s\n", insertStr, line)
				}
				newContent.WriteString(line + "\n")
				continue
			}
			findTarget(lineStr)
		} else if patternStart.MatchString(lineStr) {
			isStart = true
		}
		newContent.WriteString(line + "\n")
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	//如果存在更改字段，则更新文件内容
	if isModify {
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

		_, err = file.WriteString(newContent.String())
		if err != nil {
			Error("文件写入失败，请检查文件路径是否正确")
			return err
		}
	}

	return nil
}
