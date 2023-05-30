package rpclogic

import (
	"bytes"
	"path"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/global"
	"github.com/licat233/genzero/modules/utils"
	"github.com/licat233/genzero/sql"
	"github.com/licat233/genzero/tools"
)

type RpcLogic struct {
	Multiple bool
	Logics   LogicCollection
	DbTables sql.TableCollection
}

func New() *RpcLogic {
	ignoreTables := utils.MergeSlice(config.C.Logic.Rpc.IgnoreTables, baseIgnoreTables)
	ignoreTables = append(ignoreTables, config.C.Pb.IgnoreTables...)
	dbTables := utils.FilterTables(global.Schema.Tables, config.C.Logic.Rpc.Tables, ignoreTables)
	// dbIgoreFieldsName := utils.MergeSlice(config.C.Pb.IgnoreColumns, baseIgnoreColumns)
	logics := make(LogicCollection, 0, len(dbTables))
	for _, t := range dbTables {
		logics = append(logics, NewLogic(t.Copy()))
	}
	return &RpcLogic{
		Multiple: config.C.Logic.Rpc.Multiple,
		Logics:   logics,
		DbTables: dbTables,
	}
}

func (l *RpcLogic) Run() (err error) {
	var buf bytes.Buffer
	var goPkgName string
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
		if goPkgName == "" {
			goPkgName = logic.RpcGoPkgName
		}
		buf.WriteString(logic.PbToMd())
		buf.WriteString(logic.MdToPb())
		if s, err := logic.PbList2MdList(); err != nil {
			return err
		} else {
			buf.WriteString(s)
		}
		if s, err := logic.MdList2PbList(); err != nil {
			return err
		} else {
			buf.WriteString(s)
		}
	}
	buf.WriteString(ListReqParams(goPkgName))

	filename := path.Join(config.C.Logic.Rpc.Dir, "dataconv/dataconv.go")
	err = tools.WriteFile(filename, buf.String())
	if err != nil {
		return err
	}

	if err := tools.FormatGoFile(filename); err != nil {
		tools.Error("[logic rpc core] format go content error\n in file: %s\n error: %v", filename, err)
	}

	err = tools.RunConcurrentTasks(tasks)
	if err != nil {
		return err
	}

	if l.Multiple {
		filename = path.Join(config.C.Logic.Rpc.Dir, "base")
	} else {
		filename = config.C.Logic.Rpc.Dir
	}

	if err := tools.FormatGoFile(filename); err != nil {
		tools.Error("[logic rpc core] format go content error\n in file: %s\n error: %v", filename, err)
	}

	return nil
}
