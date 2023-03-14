package model

import (
	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/core/model/conf"
	"github.com/licat233/genzero/core/model/internal"
	"github.com/licat233/genzero/core/utils"
	"github.com/licat233/genzero/global"
	"github.com/licat233/genzero/tools"
)

type ModelCore struct {
	Tables internal.TableModelCollection
}

func New() *ModelCore {
	dbTables := utils.FilterTables(global.Schema.Tables, config.C.ModelConfig.Tables, utils.MergeSlice(config.C.ModelConfig.IgnoreTables, conf.BaseIgnoreTables))
	tables := make(internal.TableModelCollection, 0, len(dbTables))
	for _, t := range dbTables {
		tables = append(tables, *internal.NewTableModel(t.Copy()))
	}
	return &ModelCore{
		Tables: tables,
	}
}

func (m *ModelCore) Run() (err error) {
	for _, table := range m.Tables {
		if err := table.Run(); err != nil {
			return err
		}
	}
	err = tools.ExecGoimports(config.C.ModelConfig.Dir)
	return
}

func (m *ModelCore) Generate() error {
	if err := m.initTplContent(); err != nil {
		return err
	}
	for _, table := range m.Tables {
		if err := table.Run(); err != nil {
			return err
		}
	}

	return nil
}

func (m *ModelCore) initTplContent() (err error) {
	// 判断当前目录下是否存在./template/model.tpl文件
	protoTplPath := "./template/model.tpl"
	exist, err := tools.PathExists(protoTplPath)
	if err != nil {
		return
	}
	var tplContent string
	if exist {
		//如果存在，则读取该内容作为模板
		tplContent, err = tools.ReadFile(protoTplPath)
		if err != nil {
			return err
		}
	}
	if tplContent != "" {
		conf.TplContent = tplContent
	}
	return
}
