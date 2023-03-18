package apilogic

import (
	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/parser"
	"github.com/licat233/genzero/tools"
)

type Logic struct {
	CamelName      string
	LowerCamelName string
	SnakeName      string
	ModelName      string
	Dir            string
	Table          *parser.Table
}

type LogicCollection []*Logic

func NewLogic(t *parser.Table) *Logic {
	return &Logic{
		CamelName:      tools.ToCamel(t.Name),
		LowerCamelName: tools.ToLowerCamel(t.Name),
		SnakeName:      tools.ToSnake(t.Name),
		ModelName:      tools.ToCamel(t.Name) + "Model",
		Dir:            config.C.LogicConfig.Api.Dir,
		Table:          t,
	}
}
