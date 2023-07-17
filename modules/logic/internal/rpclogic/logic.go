package rpclogic

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/sql"
	"github.com/licat233/genzero/tools"
)

type Logic struct {
	CamelName      string
	LowerCamelName string
	SnakeName      string
	ModelName      string
	RpcGoPkgName   string
	PluralizedName string
	Dir            string
	Multiple       bool

	ConveFields string //注意：每个方法的数据不一样，会变，用来做临时模版渲染数据

	Table *sql.Table
}

type LogicCollection []*Logic

func NewLogic(t *sql.Table) *Logic {
	return &Logic{
		CamelName:      tools.ToCamel(t.Name),
		LowerCamelName: tools.ToLowerCamel(t.Name),
		SnakeName:      tools.ToSnake(t.Name),
		ModelName:      tools.ToCamel(t.Name) + "Model",
		RpcGoPkgName:   tools.PickGoPkgName(config.C.Pb.GoPackage),
		PluralizedName: tools.PluralizedName(tools.ToCamel(t.Name)),
		Dir:            config.C.Logic.Rpc.Dir,
		Multiple:       config.C.Logic.Rpc.Multiple,
		ConveFields:    "",
		Table:          t,
	}
}

func (l *Logic) Run() (err error) {
	if err = l.Get(); err != nil {
		return err
	}
	if err = l.Add(); err != nil {
		return err
	}
	if err = l.Put(); err != nil {
		return err
	}
	if err = l.Del(); err != nil {
		return err
	}
	if err = l.List(); err != nil {
		return err
	}
	if err = l.Enums(); err != nil {
		return err
	}
	return
}

func (l *Logic) Get() (err error) {
	filename := l.getLogicFilename("get", "")
	ok, err := l.fileValidator(filename)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	logicContentTpl := `res, err := l.svcCtx.{{.ModelName}}.FindById(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}`

	logicContent, err := tools.ParserTpl(logicContentTpl, l)
	if err != nil {
		return err
	}

	returnDataTpl := `{{.CamelName}}: dataconv.Md{{.CamelName}}ToPb{{.CamelName}}(res),`
	returnData, err := tools.ParserTpl(returnDataTpl, l)
	if err != nil {
		return err
	}

	err = modifyLogicFileContent(filename, logicContent, returnData)
	if err != nil {
		return err
	}

	return
}

func (l *Logic) Add() (err error) {
	filename := l.getLogicFilename("add", "")
	ok, err := l.fileValidator(filename)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	var conveFieldsBuf bytes.Buffer
	for _, field := range l.AddFields() {
		if field.Type == "time.Time" {
			conveFieldsBuf.WriteString(fmt.Sprintf("%s: time.Unix(in.%s, 0).Local(),\n", field.UpperCamelCaseName, field.UpperCamelCaseName))
			continue
		}
		if strings.ToLower(field.Name) == "uuid" {
			conveFieldsBuf.WriteString(fmt.Sprintf("%s: uniqueid.NewUUID(), // 这里的uniqueid包，请自己定义\n", field.UpperCamelCaseName))
			continue
		}
		conveFieldsBuf.WriteString(fmt.Sprintf("%s: in.%s,\n", field.UpperCamelCaseName, field.UpperCamelCaseName))
	}
	l.ConveFields = conveFieldsBuf.String()

	logicContentTpl := `data := &model.{{.CamelName}}{
		{{.ConveFields}}
	}
	if _, err := l.svcCtx.{{.ModelName}}.Insert(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
`

	logicContent, err := tools.ParserTpl(logicContentTpl, l)
	if err != nil {
		return err
	}

	err = modifyLogicFileContent(filename, logicContent, "")

	return err
}

func (l *Logic) Put() (err error) {
	filename := l.getLogicFilename("put", "")
	ok, err := l.fileValidator(filename)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	var conveFieldsBuf bytes.Buffer
	for _, field := range l.PutFields() {
		if field.Type == "time.Time" {
			conveFieldsBuf.WriteString(fmt.Sprintf("%s: time.Unix(in.%s, 0).Local(),\n", field.UpperCamelCaseName, field.UpperCamelCaseName))
			continue
		}
		conveFieldsBuf.WriteString(fmt.Sprintf("%s: in.%s,\n", field.UpperCamelCaseName, field.UpperCamelCaseName))
	}
	l.ConveFields = conveFieldsBuf.String()

	logicContentTpl := `data := &model.{{.CamelName}}{
		{{.ConveFields}}
	}

	if err := l.svcCtx.{{.ModelName}}.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}`

	logicContent, err := tools.ParserTpl(logicContentTpl, l)
	if err != nil {
		return err
	}

	err = modifyLogicFileContent(filename, logicContent, "")

	return err
}

func (l *Logic) Del() (err error) {
	filename := l.getLogicFilename("del", "")
	ok, err := l.fileValidator(filename)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	//分为软删除和硬删除
	var logicContentTpl string
	if l.Table.ExistIsDelField() {
		logicContentTpl = `if err := l.svcCtx.{{.ModelName}}.SoftDelete(l.ctx, in.Id); err != nil {
			l.Logger.Error(err)
			return nil, errorx.IntRpcErr(err)
		}`
	} else {
		logicContentTpl = `if err := l.svcCtx.{{.ModelName}}.Delete(l.ctx, in.Id); err != nil {
			l.Logger.Error(err)
			return nil, errorx.IntRpcErr(err)
		}`
	}

	logicContent, err := tools.ParserTpl(logicContentTpl, l)
	if err != nil {
		return err
	}

	err = modifyLogicFileContent(filename, logicContent, "")

	return
}

func (l *Logic) List() (err error) {
	filename := l.getLogicFilename("get", "list")
	ok, err := l.fileValidator(filename)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	logicContentTpl := `pageSize, page, keyword := dataconv.ListReqParams(in.ListReq)
	list, total, err := l.svcCtx.{{.ModelName}}.FindList(l.ctx, pageSize, page, keyword, dataconv.Pb{{.CamelName}}ToMd{{.CamelName}}(in.{{.CamelName}}))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}`

	logicContent, err := tools.ParserTpl(logicContentTpl, l)
	if err != nil {
		return err
	}
	returnData := fmt.Sprintf("%s: dataconv.Md%s2Pb%s(list),\nTotal:      total,", l.PluralizedName, l.PluralizedName, l.PluralizedName)

	err = modifyLogicFileContent(filename, logicContent, returnData)

	return
}

func (l *Logic) Enums() (err error) {
	//前提是要存在name字段，或者username，亦或是 nickname
	nameField := l.Table.GetNameField()
	if nameField == nil {
		return nil
	}

	filename := l.getLogicFilename("get", "enums")
	ok, err := l.fileValidator(filename)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	logicContentTpl := `list, _, err := l.svcCtx.{{.ModelName}}.FindList(l.ctx, -1, -1, "", nil)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	enums := []*{{.RpcGoPkgName}}.Enum{}
	for _, item := range list {
		enums = append(enums, &{{.RpcGoPkgName}}.Enum{
			Label: item.__NAME__,
			Value: item.Id,
		})
	}`

	logicContentTpl = strings.ReplaceAll(logicContentTpl, "__NAME__", nameField.UpperCamelCaseName)

	logicContent, err := tools.ParserTpl(logicContentTpl, l)
	if err != nil {
		return err
	}

	err = modifyLogicFileContent(filename, logicContent, "Enums: enums,")

	return
}

// 生成pb结构体转md结构体的方法
func (l *Logic) PbToMd() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("\nfunc Pb%sToMd%s(pb *%s.%s) *model.%s {\n", l.CamelName, l.CamelName, l.RpcGoPkgName, l.CamelName, l.CamelName))
	buf.WriteString(`if pb == nil {
		return nil
	}
	`)
	buf.WriteString(fmt.Sprintf("return &model.%s{\n", l.CamelName))
	for _, field := range l.PbFields() {
		if field.Type == "time.Time" {
			buf.WriteString(fmt.Sprintf("%s: time.Unix(pb.%s, 0).Local(),\n", field.UpperCamelCaseName, field.UpperCamelCaseName))
			continue
		}
		buf.WriteString(fmt.Sprintf("%s: pb.%s,\n", field.UpperCamelCaseName, field.UpperCamelCaseName))
	}
	buf.WriteString("}\n")
	buf.WriteString("}\n")
	return buf.String()
}

// 生成md结构体转pb结构体的方法
func (l *Logic) MdToPb() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("\nfunc Md%sToPb%s(md *model.%s) *%s.%s {\n", l.CamelName, l.CamelName, l.CamelName, l.RpcGoPkgName, l.CamelName))
	buf.WriteString(`if md == nil {
		return nil
	}
	`)
	buf.WriteString(fmt.Sprintf("return &%s.%s{\n", l.RpcGoPkgName, l.CamelName))
	for _, field := range l.PbFields() {
		if field.Type == "time.Time" {
			buf.WriteString(fmt.Sprintf("%s: md.%s.Unix(),\n", field.UpperCamelCaseName, field.UpperCamelCaseName))
			continue
		}
		buf.WriteString(fmt.Sprintf("%s: md.%s,\n", field.UpperCamelCaseName, field.UpperCamelCaseName))
	}
	buf.WriteString("}\n")
	buf.WriteString("}\n")
	return buf.String()
}

func (l *Logic) PbList2MdList() (string, error) {
	tpl := `
	func Pb{{.PluralizedName}}2Md{{.PluralizedName}}(pbList []*{{.RpcGoPkgName}}.{{.CamelName}}) []*model.{{.CamelName}} {
		if pbList == nil {
			return nil
		}
		data := make([]*model.{{.CamelName}}, 0)
		for _, v := range pbList {
			data = append(data, Pb{{.CamelName}}ToMd{{.CamelName}}(v))
		}
		return data
	}
	`

	return tools.ParserTpl(tpl, l)
}

func (l *Logic) MdList2PbList() (string, error) {
	tpl := `
	func Md{{.PluralizedName}}2Pb{{.PluralizedName}}(mdList []*model.{{.CamelName}}) []*{{.RpcGoPkgName}}.{{.CamelName}} {
		if mdList == nil {
			return nil
		}
		data := make([]*{{.RpcGoPkgName}}.{{.CamelName}}, 0)
		for _, v := range mdList {
			data = append(data, Md{{.CamelName}}ToPb{{.CamelName}}(v))
		}
		return data
	}
	`

	return tools.ParserTpl(tpl, l)
}

func ListReqParams(rpcGoPkgName string) string {
	tpl := `
	func ListReqParams(req *{{.RpcGoPkgName}}.ListReq) (pageSize int64, page int64, keyword string) {
		if req != nil {
			pageSize = req.PageSize
			page = req.Page
			keyword = req.Keyword
		}
		return
	}
`
	return strings.ReplaceAll(tpl, "{{.RpcGoPkgName}}", rpcGoPkgName)
}
