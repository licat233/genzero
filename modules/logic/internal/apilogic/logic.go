package apilogic

import (
	"bytes"
	"fmt"
	"path"
	"strings"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/modules/utils"
	"github.com/licat233/genzero/sql"
	"github.com/licat233/genzero/tools"
)

type Logic struct {
	CamelName      string
	LowerCamelName string
	SnakeName      string
	PluralizedName string

	ModelName    string
	RpcSvcName   string
	RpcGoPkgName string
	Dir          string

	UseRpc bool

	ConveFields string //注意：每个方法的数据不一样，会变

	Table *sql.Table
}

type LogicCollection []*Logic

func NewLogic(t *sql.Table) *Logic {
	return &Logic{
		CamelName:      tools.ToCamel(t.Name),
		LowerCamelName: tools.ToLowerCamel(t.Name),
		SnakeName:      tools.ToSnake(t.Name),
		PluralizedName: tools.PluralizedName(tools.ToCamel(t.Name)),
		ModelName:      tools.ToCamel(t.Name) + "Model",
		RpcSvcName:     tools.ToCamel(config.C.Api.ServiceName) + "Rpc",
		RpcGoPkgName:   tools.PickGoPkgName(config.C.Pb.GoPackage),
		Dir:            config.C.Logic.Api.Dir,
		UseRpc:         config.C.Logic.Api.UseRpc,
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
	dirname := path.Join(l.Dir, utils.ConvertStringStyle(config.C.Logic.Api.FileStyle, l.CamelName))
	if err := tools.FormatGoFile(dirname); err != nil {
		tools.Error("[logic api] format go content error, in dir: %s", dirname)
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

	var logicContentTpl string

	if l.UseRpc {
		logicContentTpl = `rpcResp, err := l.svcCtx.{{.RpcSvcName}}.Get{{.CamelName}}(l.ctx, &{{.RpcGoPkgName}}.Get{{.CamelName}}Req{
			Id: req.Id,
		})
		if err != nil {
			//若rpc的错误已经包装过了，无需再处理，直接返回即可
			return nil, err
		}
		pb{{.CamelName}} := rpcResp.{{.CamelName}}
		data := dataconv.Pb{{.CamelName}}ToApi{{.CamelName}}(pb{{.CamelName}})
		`
	} else {
		logicContentTpl = `md{{.CamelName}}, err := l.svcCtx.{{.ModelName}}.FindOne(l.ctx, req.Id)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Error("failed to find {{.LowerCamelName}}, error: ", err)
			return nil, errorx.InternalError(err)
		}
		data := dataconv.Md{{.CamelName}}ToApi{{.CamelName}}(md{{.CamelName}})
		`
	}

	logicContent, err := tools.ParserTpl(logicContentTpl, l)
	if err != nil {
		return err
	}

	returnContent := `return respx.DefaultSingleResp(data, nil)`

	err = modifyLogicFileContent(filename, logicContent, returnContent)
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
		if !l.UseRpc {
			if field.Type == "time.Time" {
				conveFieldsBuf.WriteString(fmt.Sprintf("%s: time.Unix(req.%s, 0).Local(),\n", field.UpperCamelCaseName, field.UpperCamelCaseName))
				continue
			} else if strings.ToLower(field.Name) == "uuid" {
				conveFieldsBuf.WriteString(fmt.Sprintf("%s: uniqueid.NewUUID(), // 这里的uniqueid包，请自己定义\n", field.UpperCamelCaseName))
				continue
			}
		}
		conveFieldsBuf.WriteString(fmt.Sprintf("%s: req.%s,\n", field.UpperCamelCaseName, field.UpperCamelCaseName))
	}
	l.ConveFields = conveFieldsBuf.String()

	var logicContentTpl string
	if l.UseRpc {
		logicContentTpl = `in := &{{.RpcGoPkgName}}.Add{{.CamelName}}Req{
			{{.ConveFields}}
		}
		if _, err := l.svcCtx.{{.RpcSvcName}}.Add{{.CamelName}}(l.ctx, in); err != nil {
			//若rpc的错误已经包装过了，无需再处理，直接返回即可
			return nil, err
		}`
	} else {
		logicContentTpl = `in := &model.{{.CamelName}}{
			{{.ConveFields}}
		}
		if _, err := l.svcCtx.{{.ModelName}}.Insert(l.ctx, in); err != nil {
		    l.Logger.Error("failed to insert {{.LowerCamelName}}, error: ", err)
		    return nil, errorx.InternalError(err)
	    }
		`
	}
	logicContent, err := tools.ParserTpl(logicContentTpl, l)
	if err != nil {
		return err
	}

	returnContent := `return respx.DefaultStatusResp(nil)`

	err = modifyLogicFileContent(filename, logicContent, returnContent)
	if err != nil {
		return err
	}

	return
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
		if !l.UseRpc {
			if field.Type == "time.Time" {
				conveFieldsBuf.WriteString(fmt.Sprintf("%s: time.Unix(req.%s, 0).Local(),\n", field.UpperCamelCaseName, field.UpperCamelCaseName))
				continue
			}
		}
		conveFieldsBuf.WriteString(fmt.Sprintf("%s: req.%s,\n", field.UpperCamelCaseName, field.UpperCamelCaseName))
	}
	l.ConveFields = conveFieldsBuf.String()

	var logicContentTpl string
	if l.UseRpc {
		logicContentTpl = `in := &{{.RpcGoPkgName}}.Put{{.CamelName}}Req{
			{{.ConveFields}}
		}
		if _, err := l.svcCtx.{{.RpcSvcName}}.Put{{.CamelName}}(l.ctx, in); err != nil {
			//若rpc的错误已经包装过了，无需再处理，直接返回即可
			return nil, err
		}`
	} else {
		logicContentTpl = `in := &model.{{.CamelName}}{
			{{.ConveFields}}
		}
		if err := l.svcCtx.{{.ModelName}}.Update(l.ctx, in); err != nil {
		    l.Logger.Error("failed to update {{.LowerCamelName}}, error: ", err)
		    return nil, errorx.InternalError(err)
	    }
		`
	}
	logicContent, err := tools.ParserTpl(logicContentTpl, l)
	if err != nil {
		return err
	}

	returnContent := `return respx.DefaultStatusResp(nil)`

	err = modifyLogicFileContent(filename, logicContent, returnContent)
	if err != nil {
		return err
	}

	return
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

	var logicContentTpl string
	if l.UseRpc {
		logicContentTpl = `_, err := l.svcCtx.{{.RpcSvcName}}.Del{{.CamelName}}(l.ctx, &{{.RpcGoPkgName}}.Del{{.CamelName}}Req{
			Id: req.Id,
		})
		if err != nil {
			//若rpc的错误已经包装过了，无需再处理，直接返回即可
			return nil, err
		}`
	} else {
		//分为软删除和硬删除
		if l.Table.ExistIsDelField() {
			logicContentTpl = `if err := l.svcCtx.{{.ModelName}}.SoftDelete(l.ctx, req.Id); err != nil {
				l.Logger.Error("failed to soft delete {{.LowerCamelName}}, error: ", err)
				return nil, errorx.InternalError(err)
			}`
		} else {
			logicContentTpl = `if err := l.svcCtx.{{.ModelName}}.Delete(l.ctx, req.Id); err != nil {
				l.Logger.Error("failed to delete {{.LowerCamelName}}, error: ", err)
				return nil, errorx.InternalError(err)
			}`
		}
	}
	logicContent, err := tools.ParserTpl(logicContentTpl, l)
	if err != nil {
		return err
	}

	returnContent := `return respx.DefaultStatusResp(nil)`

	err = modifyLogicFileContent(filename, logicContent, returnContent)
	if err != nil {
		return err
	}

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

	var conveFieldsBuf bytes.Buffer
	for _, field := range l.ApiFields() {
		conveFieldsBuf.WriteString(fmt.Sprintf("%s: req.%s,\n", field.UpperCamelCaseName, field.UpperCamelCaseName))
	}
	l.ConveFields = conveFieldsBuf.String()

	var logicContentTpl string
	var returnContent string
	if l.UseRpc {
		logicContentTpl = `in := &{{.RpcGoPkgName}}.Get{{.CamelName}}ListReq{
			ListReq: &{{.RpcGoPkgName}}.ListReq{
				PageSize: req.PageSize,
				Page:     req.Page,
				Keyword:  req.Keyword,
			},
			{{.CamelName}}: &{{.RpcGoPkgName}}.{{.CamelName}}{
				{{.ConveFields}}
			},
		}
		rpcResp, err := l.svcCtx.{{.RpcSvcName}}.Get{{.CamelName}}List(l.ctx, in)
		if err != nil {
			//rpc的错误已经包装过了，无需再处理，直接返回即可
			return nil, err
		}
		pbList := rpcResp.{{.PluralizedName}}
		data := dataconv.Pb{{.PluralizedName}}2Api{{.PluralizedName}}(pbList)
		`
		returnContent = `return respx.DefaultListResp(data, rpcResp.Total, req.PageSize, req.Page, nil)`
	} else {
		logicContentTpl = `in := &model.{{.CamelName}}{
			{{.ConveFields}}
		}
		mdList, total, err := l.svcCtx.{{.ModelName}}.FindList(l.ctx, req.PageSize, req.Page, req.Keyword, in)
		if err != nil {
			l.Logger.Error("failed to query {{.LowerCamelName}} list, error: ", err)
			return nil, errorx.InternalError(err)
		}
		data := dataconv.Md{{.PluralizedName}}2Api{{.PluralizedName}}(mdList)
		`
		returnContent = `return respx.DefaultListResp(data, total, req.PageSize, req.Page, nil)`
	}
	logicContent, err := tools.ParserTpl(logicContentTpl, l)
	if err != nil {
		return err
	}

	err = modifyLogicFileContent(filename, logicContent, returnContent)
	if err != nil {
		return err
	}

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

	var logicContentTpl string
	if l.UseRpc {
		logicContentTpl = `rpcResp, err := l.svcCtx.{{.RpcSvcName}}.Get{{.CamelName}}Enums(l.ctx, &{{.RpcGoPkgName}}.Get{{.CamelName}}EnumsReq{
			ParentId: req.ParentId,
		})
	if err != nil {
		return nil, err
	}
	data := make([]*types.Enum, 0)
	for _, v := range rpcResp.Enums {
		data = append(data, &types.Enum{
			Label: v.Label,
			Value: v.Value,
		})
	}
		`
	} else {
		logicContentTpl = `md{{.CamelName}}List, err := l.svcCtx.{{.ModelName}}.FindAll(l.ctx)
		if err != nil {
			l.Logger.Error("failed to query {{.LowerCamelName}} list, error: ", err)
			return nil, errorx.InternalError(err)
		}
		data := make([]*types.Enum, 0)
	for _, v := range md{{.CamelName}}List {
		data = append(data, &types.Enum{
			Label: v.__NAME__,
			Value: v.Id,
		})
	}
		`
		logicContentTpl = strings.ReplaceAll(logicContentTpl, "__NAME__", nameField.UpperCamelCaseName)
	}

	logicContent, err := tools.ParserTpl(logicContentTpl, l)
	if err != nil {
		return err
	}

	returnContent := `return respx.DefaultSingleResp(data, nil)`

	err = modifyLogicFileContent(filename, logicContent, returnContent)
	if err != nil {
		return err
	}

	return
}

func (l *Logic) PbToApi() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("\nfunc Pb%sToApi%s(pb *%s.%s) *types.%s {\n", l.CamelName, l.CamelName, l.RpcGoPkgName, l.CamelName, l.CamelName))
	buf.WriteString(`if pb == nil {
		return nil
	}
	`)
	buf.WriteString(fmt.Sprintf("return &types.%s{\n", l.CamelName))
	for _, field := range l.PbFields() {
		buf.WriteString(fmt.Sprintf("%s: pb.%s,\n", field.UpperCamelCaseName, field.UpperCamelCaseName))
	}
	buf.WriteString("}\n")
	buf.WriteString("}\n")
	return buf.String()
}

func (l *Logic) MdToApi() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("\nfunc Md%sToApi%s(md *%s.%s) *types.%s {\n", l.CamelName, l.CamelName, "model", l.CamelName, l.CamelName))
	buf.WriteString(`if md == nil {
		return nil
	}
	`)
	buf.WriteString(fmt.Sprintf("return &types.%s{\n", l.CamelName))
	for _, field := range l.ModelFields() {
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

func (l *Logic) PbList2ApiList() (string, error) {
	tpl := `
	func Pb{{.PluralizedName}}2Api{{.PluralizedName}}(pbList []*{{.RpcGoPkgName}}.{{.CamelName}}) []*types.{{.CamelName}} {
		if pbList == nil {
			return nil
		}
		data := make([]*types.{{.CamelName}}, 0)
		for _, v := range pbList {
			data = append(data, Pb{{.CamelName}}ToApi{{.CamelName}}(v))
		}
		return data
	}
	`

	return tools.ParserTpl(tpl, l)
}

func (l *Logic) MdList2ApiList() (string, error) {
	tpl := `
	func Md{{.PluralizedName}}2Api{{.PluralizedName}}(mdList []*model.{{.CamelName}}) []*types.{{.CamelName}} {
		if mdList == nil {
			return nil
		}
		data := make([]*types.{{.CamelName}}, 0)
		for _, v := range mdList {
			data = append(data, Md{{.CamelName}}ToApi{{.CamelName}}(v))
		}
		return data
	}
	`

	return tools.ParserTpl(tpl, l)
}
