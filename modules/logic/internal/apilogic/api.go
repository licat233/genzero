package apilogic

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/global"
	"github.com/licat233/genzero/modules/utils"
	"github.com/licat233/genzero/sql"
	"github.com/licat233/genzero/tools"
)

type ApiLogic struct {
	Logics   LogicCollection
	DbTables sql.TableCollection

	RpcGoPkgName string
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
		Logics:       logics,
		DbTables:     dbTables,
		RpcGoPkgName: tools.PickGoPkgName(config.C.Pb.GoPackage),
	}
}

func (l *ApiLogic) Run() error {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s\n", config.DoNotEdit))
	buf.WriteString("package dataconv\n\n")
	if config.C.Logic.Api.UseRpc {
		buf.WriteString(l.commonPbToApi())
	}
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
		if s, err := logic.MdList2ApiList(); err != nil {
			return err
		} else {
			buf.WriteString(s)
		}
		if config.C.Logic.Api.UseRpc {
			buf.WriteString(logic.PbToApi())
			if s, err := logic.PbList2ApiList(); err != nil {
				return err
			} else {
				buf.WriteString(s)
			}
		}
	}
	filename := filepath.Join(config.C.Logic.Api.Dir, "dataconv/dataconv_gen.go")
	err := tools.WriteFile(filename, buf.String())
	if err != nil {
		return err
	}

	if err := tools.FormatGoFile(filename); err != nil {
		tools.Error("[logic api] format go content error\n in file: %s\n error: %v", filename, err)
	}

	err = tools.RunConcurrentTasks(tasks)
	if err != nil {
		return err
	}

	//如果没有使用rpc服务，而是单体服务，则修改svc文件，添加model字段
	if !config.C.Logic.Api.UseRpc {
		var tableModelsDefin []string
		for _, t := range l.DbTables {
			name := tools.ToCamel(t.Name)
			tableModelsDefin = append(tableModelsDefin, fmt.Sprintf("model.%sModel", name))
		}

		svcFilePath := fmt.Sprintf("%s/../svc/serviceContext.go", config.C.Logic.Api.Dir)

		err = tools.InsertFieldsToSvcFile(svcFilePath, tableModelsDefin)
		if err != nil {
			return err
		}

		if err := tools.FormatGoFile(svcFilePath); err != nil {
			tools.Error("[logic api] format go content error\n in file: %s\n error: %v", svcFilePath, err)
		}
	}

	return nil
}

func (l *ApiLogic) commonPbToApi() string {
	tpl := `
	func PbEnumToApiEnum(in *__GORPCNAME__.Enum) *types.Enum {
		if in == nil {
			return nil
		}
		return &types.Enum{
			Label: in.Label,
			Value: in.Value,
		}
	}

	func PbEnumsToApiEnums(list []*__GORPCNAME__.Enum) []*types.Enum {
		res := []*types.Enum{}
		for _, v := range list {
			res = append(res, PbEnumToApiEnum(v))
		}
		return res
	}

	func PbOptionToApiOption(in *__GORPCNAME__.Option) *types.Option {
		if in == nil {
			return nil
		}
		return &types.Option{
			Title: in.Title,
			Value: in.Value,
		}
	}

	func PbOptionsToApiOptions(list []*__GORPCNAME__.Option) []*types.Option {
		res := []*types.Option{}
		for _, v := range list {
			res = append(res, PbOptionToApiOption(v))
		}
		return res
	}

	func PbTreeOptionToApiTreeOption(in *__GORPCNAME__.TreeOption) *types.TreeOption {
		if in == nil {
			return nil
		}
		return &types.TreeOption{
			Title:    in.Title,
			Value:    in.Value,
			Children: PbTreeOptionsToApiTreeOptions(in.Children),
		}
	}

	func PbTreeOptionsToApiTreeOptions(list []*__GORPCNAME__.TreeOption) []types.TreeOption {
		res := []types.TreeOption{}
		for _, v := range list {
			res = append(res, *PbTreeOptionToApiTreeOption(v))
		}
		return res
	}
	`

	res := strings.ReplaceAll(tpl, "__GORPCNAME__", l.RpcGoPkgName)

	return res
}
