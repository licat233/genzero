package internal

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/core/model/conf"
	"github.com/licat233/genzero/core/model/internal/funcs"
	"github.com/licat233/genzero/sql"
	"github.com/licat233/genzero/tools"
)

type TableModel struct {
	ProjectName   string
	TableName     string
	InterfaceName string
	OutFileName   string
	TplContent    string
	IsCacheMode   bool
	table         *sql.Table
	Funcs         []funcs.ModelFunc
}

type TableModelCollection []*TableModel

func NewTableModel(t *sql.Table) *TableModel {
	lowerName := tools.ToLowerCamel(t.Name)
	return &TableModel{
		ProjectName:   config.ProjectName,
		TableName:     t.Name,
		InterfaceName: fmt.Sprintf("%s_model", lowerName),
		OutFileName:   path.Join(config.C.Model.Dir, fmt.Sprintf("%sModel_extend.go", lowerName)),
		TplContent:    conf.TplContent,
		IsCacheMode:   false,
		table:         t,
		Funcs: []funcs.ModelFunc{
			funcs.NewFindAll(t),
			funcs.NewFindList(t),
			funcs.NewTableName(t),
			funcs.NewFindByUuid(t),
			funcs.NewSoftDelete(t),
			funcs.NewFormatUuidKey(t),
		},
	}
}

func (t *TableModel) Run() error {
	return t.Generate()
}

func (t *TableModel) Generate() error {
	err := t.Init()
	if err != nil {
		return err
	}

	content, err := t.Render()
	if err != nil {
		return err
	}

	if tools.TrimSpace(content) == "" {
		return errors.New("content is empty")
	}

	err = t.WriteFile(content)
	if err != nil {
		return err
	}
	err = t.ExtendOriginalInterface()
	if err != nil {
		return err
	}
	err = tools.FormatGoFile(t.OutFileName)
	if err != nil {
		tools.Error("[model core] format go content error, in file: %s", t.OutFileName)
	}
	return err
}

func (t *TableModel) WriteFile(content string) error {
	return tools.WriteFile(t.OutFileName, content)
}

func (t *TableModel) Render() (string, error) {
	tmpl, err := tools.Template("model").Parse(t.TplContent)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, *t)
	if err != nil {
		return "", err
	}
	content := buf.String()

	return content, err
}

func (t *TableModel) Init() (err error) {
	isCache, err := t.isCacheMode()
	if err != nil {
		return err
	}
	t.IsCacheMode = isCache
	// conf.ChangeQueryString(isCache)
	return nil
}

func (t *TableModel) isCacheMode() (bool, error) {
	genFilename := fmt.Sprintf("%sModel_gen.go", tools.ToLowerCamel(t.TableName))
	filePath := path.Join(config.C.Model.Dir, genFilename)
	exists, err := tools.PathExists(filePath)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, fmt.Errorf("[model core] Initialization failed: the (%s) file does not exist. Please use the goctl tool to create it first.\n - goctl: https://go-zero.dev/cn/docs/goctl/goctl/", filePath)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "sqlc.CachedConn" {
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}
	return false, nil
}

// 拓展原始接口，添加当前定义的接口
func (t *TableModel) ExtendOriginalInterface() error {
	genFilename := fmt.Sprintf("%sModel_gen.go", tools.ToLowerCamel(t.TableName))
	filePath, err := tools.FindFile(config.C.Model.Dir, genFilename)
	if err != nil {
		return err
	}
	has, err := tools.PathExists(filePath)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("[model core] Initialization failed: the (%s) file does not exist. Please use the goctl tool to create it first.\n - goctl: https://go-zero.dev/cn/docs/goctl/goctl/", filePath)
	}

	// 打开文件以供读取
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	table := tools.ToLowerCamel(t.TableName)
	target := fmt.Sprintf("%sModel interface {", table)

	// 使用bufio.Scanner获取文件中每一行的内容
	scanner := bufio.NewScanner(file)

	// 读取每行内容并进行修改
	var exist = false
	var newContent = new(bytes.Buffer)
	for scanner.Scan() {
		line := scanner.Text()
		if exist && strings.HasSuffix(line, t.InterfaceName) {
			//已经存在，不需要修改
			return nil
		}
		if strings.HasSuffix(strings.TrimSpace(line), target) {
			exist = true
			line = fmt.Sprintf("%s // extends interface\n\t\t%s", line, t.InterfaceName)
		}
		newContent.WriteString(line + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return err
	}

	// 将修改后的内容写回到文件中
	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = file.WriteString(newContent.String())
	if err != nil {
		tools.Error("文件写入失败，请检查文件路径是否正确")
		return err
	}

	return nil
}
