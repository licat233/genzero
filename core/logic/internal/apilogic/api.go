package apilogic

import (
	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/core/utils"
	"github.com/licat233/genzero/global"
	"github.com/licat233/genzero/parser"
)

type ApiLogic struct {
	Logics LogicCollection
}

var baseIgnoreTables = []string{}

func New(t *parser.Table) *ApiLogic {
	dbTables := utils.FilterTables(global.Schema.Tables, config.C.LogicConfig.Api.Tables, utils.MergeSlice(config.C.LogicConfig.Api.IgnoreTables, baseIgnoreTables))
	logics := make(LogicCollection, 0, len(dbTables))
	for _, t := range dbTables {
		logics = append(logics, NewLogic(t.Copy()))
	}
	return &ApiLogic{
		Logics: logics,
	}
}

func (l *ApiLogic) Run() error {
	return nil
}
