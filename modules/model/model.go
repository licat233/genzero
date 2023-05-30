package model

import (
	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/global"
	"github.com/licat233/genzero/modules/model/conf"
	"github.com/licat233/genzero/modules/model/internal"
	"github.com/licat233/genzero/modules/utils"
	"github.com/licat233/genzero/tools"
)

type ModelModule struct {
	Tables internal.TableModelCollection
}

func New() *ModelModule {
	dbTables := utils.FilterTables(global.Schema.Tables, config.C.Model.Tables, utils.MergeSlice(config.C.Model.IgnoreTables, conf.BaseIgnoreTables))
	tables := make(internal.TableModelCollection, 0, len(dbTables))
	for _, t := range dbTables {
		tables = append(tables, internal.NewTableModel(t.Copy()))
	}
	return &ModelModule{
		Tables: tables,
	}
}

func (m *ModelModule) Run() (err error) {
	err = m.Generate()
	if err != nil {
		tools.Error("generate model extend file faild.")
		return err
	}
	tools.Success("generate model file success.")
	return
}

func (m *ModelModule) Generate() (err error) {
	if err = m.initTplContent(); err != nil {
		return err
	}
	if len(m.Tables) == 0 {
		return nil
	}
	tasks := make([]tools.TaskFunc, 0, len(m.Tables))
	for _, tableModel := range m.Tables {
		if tableModel == nil {
			continue
		}
		localTableModel := tableModel // 为每个任务创建一个本地变量
		tasks = append(tasks, func() error {
			return localTableModel.Run()
		})
	}

	err = tools.RunConcurrentTasks(tasks)
	if err != nil {
		return err
	}

	err = tools.ExecGoimports(config.C.Model.Dir)
	if err != nil {
		return err
	}
	_, err = tools.ExecShell("go get -u github.com/Masterminds/squirrel")
	if err != nil {
		tools.Warning("go get github.com/Masterminds/squirrel failed: %v", err)
	}
	return nil
}

func (m *ModelModule) initTplContent() (err error) {
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
