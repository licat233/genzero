package rpclogic

import (
	"strings"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/core/logic/conf"
	"github.com/licat233/genzero/parser"
	"github.com/licat233/genzero/tools"
)

type Logic struct {
	CamelName      string
	LowerCamelName string
	SnakeName      string
	ModelName      string
	Dir            string
	Multiple       bool
	Table          *parser.Table
}

type LogicCollection []*Logic

func NewLogic(t *parser.Table) *Logic {
	return &Logic{
		CamelName:      tools.ToCamel(t.Name),
		LowerCamelName: tools.ToLowerCamel(t.Name),
		SnakeName:      tools.ToSnake(t.Name),
		ModelName:      tools.ToCamel(t.Name) + "Model",
		Dir:            config.C.LogicConfig.Rpc.Dir,
		Multiple:       config.C.LogicConfig.Rpc.Multiple,
		Table:          t,
	}
}

func (l *Logic) Get() (err error) {
	body, err := l.getLogicFileContent("get", "")
	if err != nil {
		return err
	}
	if len(body) == 0 {
		return nil
	}
	//判断是否存在 // todo: add your logic here and delete this line
	if !strings.Contains(body, conf.TodoMark) {
		return nil
	}

	logicContentTpl := `res, err := l.svcCtx.{{.ModelName}}.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, nil
	}`

	logicContent, err := tools.ParserTpl(logicContentTpl, l)
	if err != nil {
		return err
	}

	returnDataTpl := `{{.CamelName}}: conv.Md{{.CamelName}}ToPb{{.CamelName}}(res),`
	returnData, err := tools.ParserTpl(returnDataTpl, l)
	if err != nil {
		return err
	}

	filename := l.getLogicFilename("get", "")
	err = modifyLogicFileContent(filename, logicContent, returnData)
	if err != nil {
		return err
	}

	return
}

func (l *Logic) Add() (err error) {
	body, err := l.getLogicFileContent("add", "")
	if err != nil {
		return err
	}
	if len(body) == 0 {
		return nil
	}
	//判断是否存在 // todo: add your logic here and delete this line
	if !strings.Contains(body, conf.TodoMark) {
		return nil
	}

	logicContentTpl := `data := &model.{{.CamelName}}{
		{{.ConverList}}
	}
	result, err := l.svcCtx.{{.ModelName}}.Insert(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	res, err := l.svcCtx.{{.CamelName}}Model.FindOne(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}`

	logicContent, err := tools.ParserTpl(logicContentTpl, l)
	if err != nil {
		return err
	}

	returnDataTpl := `{{.CamelName}}: conv.Md{{.CamelName}}ToPb{{.CamelName}}(res),`
	returnData, err := tools.ParserTpl(returnDataTpl, l)
	if err != nil {
		return err
	}

	filename := l.getLogicFilename("add", "")
	err = modifyLogicFileContent(filename, logicContent, returnData)
	if err != nil {
		return err
	}

	return err
}
