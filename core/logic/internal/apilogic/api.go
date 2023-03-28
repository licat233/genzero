package apilogic

import (
	"bytes"
	"path"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/core/utils"
	"github.com/licat233/genzero/global"
	"github.com/licat233/genzero/sql"
	"github.com/licat233/genzero/tools"
)

type ApiLogic struct {
	Logics   LogicCollection
	DbTables sql.TableCollection
}

func New() *ApiLogic {
	ignoreTables := utils.MergeSlice(config.C.Logic.Api.IgnoreTables, baseIgnoreTables)
	ignoreTables = append(ignoreTables, config.C.Api.IgnoreTables...)
	dbTables := utils.FilterTables(global.Schema.Tables, config.C.Logic.Api.Tables, ignoreTables)
	// dbIgoreFieldsName := utils.MergeSlice(config.C.Api.IgnoreColumns, baseIgnoreColumns)
	logics := make(LogicCollection, 0, len(dbTables))
	for _, t := range dbTables {
		logics = append(logics, NewLogic(t.Copy()))
	}
	return &ApiLogic{
		Logics:   logics,
		DbTables: dbTables,
	}
}

func (l *ApiLogic) Run() error {
	var buf bytes.Buffer
	buf.WriteString("package dataconv\n\n")
	tasks := make([]tools.TaskFunc, 0, len(l.Logics))
	for _, logic := range l.Logics {
		localLogic := logic // 为每个任务创建一个本地变量
		tasks = append(tasks, func() error {
			return localLogic.Run()
		})
		// if err := logic.Run(); err != nil {
		// 	return err
		// }
		buf.WriteString(logic.MdToApi())
		buf.WriteString(logic.PbToApi())
		if s, err := logic.MdList2ApiList(); err != nil {
			return err
		} else {
			buf.WriteString(s)
		}
		if s, err := logic.PbList2ApiList(); err != nil {
			return err
		} else {
			buf.WriteString(s)
		}
	}
	filename := path.Join(config.C.Logic.Api.Dir, "dataconv/dataconv.go")
	err := tools.WriteFile(filename, buf.String())
	if err != nil {
		return err
	}

	if err := tools.FormatGoFile(filename); err != nil {
		tools.Error("[logic api core] format go content error\n in file: %s\n error: %v", filename, err)
	}

	err = tools.RunConcurrentTasks(tasks)
	if err != nil {
		return err
	}

	return nil
}
