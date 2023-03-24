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
	for _, logic := range l.Logics {
		if err := logic.Run(); err != nil {
			return err
		}
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
	dirname := path.Join(config.C.Logic.Api.Dir, "dataconv")
	err := tools.MakeDir(dirname)
	if err != nil {
		return err
	}
	filename := path.Join(dirname, "dataconv.go")
	err = tools.WriteFile(filename, buf.String())
	if err != nil {
		return err
	}

	if err := tools.FormatGoFile(dirname); err != nil {
		tools.Error("[logic api core] format go content error, in dir: %s", dirname)
	}
	return nil
}
