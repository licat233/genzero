package rpclogic

import (
	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/core/utils"
	"github.com/licat233/genzero/global"
	"github.com/licat233/genzero/parser"
)

type RpcLogic struct {
	Logics LogicCollection
}

var baseIgnoreTables = []string{}

func New(t *parser.Table) *RpcLogic {
	dbTables := utils.FilterTables(global.Schema.Tables, config.C.LogicConfig.Rpc.Tables, utils.MergeSlice(config.C.LogicConfig.Rpc.IgnoreTables, baseIgnoreTables))
	logics := make(LogicCollection, 0, len(dbTables))
	for _, t := range dbTables {
		logics = append(logics, NewLogic(t.Copy()))
	}
	return &RpcLogic{
		Logics: logics,
	}
}

func (l *RpcLogic) Run() error {
	return nil
}
