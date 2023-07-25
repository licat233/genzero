package internal

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/modules/model/conf"
	"github.com/licat233/genzero/modules/model/internal/funcs"
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
	FuncNameList  []string
}

type TableModelCollection []*TableModel

func NewTableModel(t *sql.Table) *TableModel {
	lowerName := tools.ToLowerCamel(t.Name)
	return &TableModel{
		ProjectName:   config.ProjectName,
		TableName:     t.Name,
		InterfaceName: fmt.Sprintf("%s_model", lowerName),
		OutFileName:   filepath.Join(config.C.Model.Dir, fmt.Sprintf("%sModel_extend.go", lowerName)),
		TplContent:    conf.TplContent,
		IsCacheMode:   false,
		table:         t,
		Funcs:         []funcs.ModelFunc{},
		FuncNameList:  []string{},
	}
}

func (t *TableModel) Init() (err error) {
	isCache, err := IsCacheMode(t.TableName)
	if err != nil {
		return err
	}
	t.IsCacheMode = isCache
	t.Funcs = []funcs.ModelFunc{
		funcs.NewFindCount(t.table, isCache),
		funcs.NewFindAll(t.table, isCache),
		funcs.NewFindList(t.table, isCache),
		funcs.NewTableName(t.table, isCache),
		funcs.NewFindByAnyCollection(t.table, isCache),
		funcs.NewFindsByFieldsCollection(t.table, isCache),
		funcs.NewFindsByFieldCollection(t.table, isCache),
	}
	if t.table.ExistUuidField() {
		t.Funcs = append(t.Funcs, funcs.NewFormatUuidKey(t.table, isCache))
	}
	if t.table.ExistIsDelField() {
		t.Funcs = append(t.Funcs, funcs.NewSoftDelete(t.table, isCache))
	}
	for _, f := range t.Funcs {
		t.FuncNameList = append(t.FuncNameList, f.FullName())
	}
	return nil
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

// 拓展原始接口，添加当前定义的接口
func (t *TableModel) ExtendOriginalInterface() error {
	genFilename := fmt.Sprintf("%sModel.go", tools.ToLowerCamel(t.TableName))
	filePath, err := tools.FindFile(config.C.Model.Dir, genFilename)
	if err != nil {
		return err
	}
	has, err := tools.PathExists(filePath)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("[model] Initialization failed: the (%s) file does not exist. Please use the goctl tool to create it first.\n - goctl: https://go-zero.dev/cn/docs/goctl/goctl/", filePath)
	}

	// 打开文件以供读取
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	table := tools.ToLowerCamel(t.TableName)
	interfaceMark := fmt.Sprintf("%sModel", table)
	interfaceInsertLine := fmt.Sprintf("%s // extended interface by genzero", t.InterfaceName)
	cachePrefixMark1 := fmt.Sprintf("cache%sIdPrefix", tools.ToCamel(table))
	cachePrefixMark2 := fmt.Sprintf("&custom%sModel{", tools.ToCamel(table))
	cachePrefixInsertLine := fmt.Sprintf("%s = \"cache:%s:%s:id:\"  // modifying cache id prefix by genzero", cachePrefixMark1, tools.ToLowerCamel(config.DatabaseName), table)

	// 使用bufio.Scanner获取文件中每一行的内容
	scanner := bufio.NewScanner(file)

	//修改接口
	var interfaceModified = false
	//修改cache前缀
	var prefixModified = false
	var newContent = new(bytes.Buffer)
	// 读取每行内容并进行修改
	for scanner.Scan() {
		line := scanner.Text()
		lineTemp := strings.TrimSpace(line)
		if !interfaceModified && lineTemp == interfaceMark {
			interfaceModified = true
			line = fmt.Sprintf("%s // extended interface by gozero\n\t\t%s", line, interfaceInsertLine)
		} else if t.IsCacheMode && !prefixModified {
			if strings.HasPrefix(lineTemp, cachePrefixMark1) {
				//表示已经存在，修改过
				prefixModified = true
			} else if strings.HasSuffix(lineTemp, cachePrefixMark2) {
				prefixModified = true
				line = fmt.Sprintf("\t%s\n%s", cachePrefixInsertLine, line)
			}
		}
		newContent.WriteString(line + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return err
	}

	if !interfaceModified && !prefixModified {
		return nil
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

	_, err = file.WriteString(newContent.String())
	if err != nil {
		tools.Error("文件写入失败，请检查文件路径是否正确")
		return err
	}

	return nil
}

func IsCacheMode(tableName string) (bool, error) {
	genFilename := fmt.Sprintf("%sModel_gen.go", tools.ToLowerCamel(tableName))
	filePath := filepath.Join(config.C.Model.Dir, genFilename)
	exists, err := tools.PathExists(filePath)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, fmt.Errorf("[model] Initialization failed: the (%s) file does not exist. Please use the goctl tool to create it first.\n - goctl: https://go-zero.dev/cn/docs/goctl/goctl/", filePath)
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
		if strings.HasSuffix(line, "sqlc.CachedConn") {
			return true, nil
		} else if strings.HasSuffix(line, "sqlc.Conn") {
			return false, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}
	return false, nil
}
