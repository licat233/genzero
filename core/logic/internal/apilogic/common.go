package apilogic

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/licat233/genzero/config"
	apiconf "github.com/licat233/genzero/core/api/conf"
	"github.com/licat233/genzero/core/logic/conf"
	pbconf "github.com/licat233/genzero/core/pb/conf"
	"github.com/licat233/genzero/core/utils"
	"github.com/licat233/genzero/sql"
	"github.com/licat233/genzero/tools"
)

func (l *Logic) getLogicFilename(method string, supply string) string {
	var filename string
	method = tools.ToLowerCamel(method)
	supply = tools.ToCamel(supply)
	filename = fmt.Sprintf("%s%s%sLogic", method, l.CamelName, supply)
	switch config.C.Logic.Rpc.Style {
	case config.CamelCase:
		filename = tools.ToCamel(filename)
	case config.LowerCamelCase:
		filename = tools.ToLowerCamel(filename)
	case config.SnakeCase:
		filename = tools.ToSnake(filename)
	default:
		filename = tools.ToLowerCamel(filename)
	}
	filename = path.Join(l.Dir, l.LowerCamelName, filename)
	return filename + ".go"
}

func (l *Logic) getLogicFileContent(filename string) (string, error) {
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
	// writer := bufio.NewWriter(file)
	var writer bytes.Buffer

	//用于标记 todo是否查找到
	findedTodo := false
	findedFunc := false
	findedReturn := false
	// 读取文件中的每一行，并进行操作
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		// 更改 resp 类型
		if !findedFunc && tools.IsFunction(line) { //函数声明行
			//替换响应体
			line = tools.ReplaceFuncResp(line, "(any, error)")
			findedFunc = true
			goto WRITE
		}

		// 更改 todo
		if !findedTodo && findedFunc && strings.TrimSpace(line) == conf.TodoMark {
			line = logicContent + "\n"
			findedTodo = true
			goto WRITE
		}

		// 更改 return
		if !findedReturn && findedTodo && IsReturn(line) {
			returnContent = tools.TrimSpace(returnContent)
			if returnContent == "" {
				returnContent = "return respx.DefaultStateResp(nil)"
			}
			line = returnContent
			findedReturn = true
		}

	WRITE:
		// 写入新的内容到文件
		_, err = writer.WriteString(line)
		if err != nil {
			tools.Error("write file error: %s", err)
			panic(err)
		}
	}
	content := writer.String()
	err = tools.WriteFile(filename, content)
	return err
}

func (l *Logic) ModelFields() []sql.Field {
	res := make([]sql.Field, 0)
	for _, field := range l.Table.Fields {
		res = append(res, field)
	}
	return res
}

func (l *Logic) PbFields() []sql.Field {
	pbIgnoreCol := utils.MergeSlice(config.C.Pb.IgnoreColumns, pbconf.BaseIgnoreColumns)
	cols := utils.FilterIgnoreFields(l.Table.Fields, pbIgnoreCol)
	res := make([]sql.Field, 0)
	for _, field := range cols {
		res = append(res, field)
	}
	return res
}

func (l *Logic) ApiFields() []sql.Field {
	apiIgnoreCol := utils.MergeSlice(config.C.Api.IgnoreColumns, apiconf.BaseIgnoreColumns)
	cols := utils.FilterIgnoreFields(l.Table.Fields, apiIgnoreCol)
	res := make([]sql.Field, 0)
	for _, field := range cols {
		res = append(res, field)
	}
	return res
}

func (l *Logic) AddFields() []sql.Field {
	res := make([]sql.Field, 0)
	for _, field := range l.Table.Fields {
		if tools.HasInSlice(moreIgnoreColumns, field.Name) {
			continue
		}
		res = append(res, field)
	}
	return res
}

func (l *Logic) PutFields() []sql.Field {
	res := make([]sql.Field, 0)
	for _, field := range l.Table.Fields {
		res = append(res, field)
	}
	return res
}

// 提取proto的包名
func PickRpcGoPackageName(s string) string {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "func (") {
			list := regexp.MustCompile(`[a-zA-Z]\w+?\(in\s\*([\w\.]+?)\.`).FindStringSubmatch(line)
			if len(list) >= 2 {
				return strings.TrimSpace(list[1])
			}
		}
	}
	return ""
}

func (l *Logic) fileValidator(filename string) (bool, error) {
	body, err := l.getLogicFileContent(filename)
	if err != nil {
		return false, err
	}
	if len(body) == 0 {
		return false, nil
	}
	if l.RpcGoPkgName == "" {
		l.RpcGoPkgName = config.C.Pb.GoPackage
	}
	//判断是否存在 // todo: add your logic here and delete this line
	if !tools.ContainsString(body, conf.TodoMark) {
		return false, nil
	}

	return true, nil
}

func IsReturn(line string) bool {
	line = strings.TrimSpace(line)
	return strings.HasPrefix(line, "return ") || line == "return"
}
