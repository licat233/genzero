package api

import (
	"bytes"
	"path/filepath"
	"text/template"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/global"
	"github.com/licat233/genzero/modules/api/conf"
	"github.com/licat233/genzero/modules/api/internal"
	"github.com/licat233/genzero/modules/utils"
	"github.com/licat233/genzero/sql"
	"github.com/licat233/genzero/tools"
)

type ApiModule struct {
	ProjectAuthor  string
	ProjectName    string
	ProjectAddr    string
	ProjectVersion string

	TplContent        string
	OutFileName       string
	OldContent        string
	Multiple          bool
	CurrentIsCoreFile bool //标记当前是否为核心文件，用于生成唯一内容，避免重复
	ServiceName       string
	ServiceComment    string

	ImportStartMark        string
	ImportEndMark          string
	CustomImportStartMark  string
	CustomImportContent    string
	CustomImportEndMark    string
	StructStartMark        string
	StructEndMark          string
	CustomStructStartMark  string
	CustomStructContent    string
	CustomStructEndMark    string
	ServiceStartMark       string
	ServiceEndMark         string
	CustomServiceStartMark string
	CustomServiceContent   string
	CustomServiceEndMark   string

	DbTables          sql.TableCollection
	DbIgoreFieldsName []string

	Imports     internal.ImportCollection
	BaseStructs internal.StructCollection
	Structs     internal.StructCollection
	Services    internal.ServiceCollection
}

func getOutFilename(name string) string {
	return filepath.Join(config.C.Api.Dir, tools.ToLowerCamel(name)+".api")
}

func New() *ApiModule {
	return &ApiModule{
		ProjectAuthor:          tools.GetCurrentUserName(),
		ProjectName:            config.ProjectName,
		ProjectAddr:            config.ProjectURL,
		ProjectVersion:         config.CurrentVersion,
		TplContent:             conf.TplContent,
		OutFileName:            getOutFilename(config.C.Api.ServiceName),
		OldContent:             "",
		Multiple:               config.C.Api.Multiple,
		CurrentIsCoreFile:      true,
		ServiceName:            config.C.Api.ServiceName,
		ServiceComment:         global.Schema.Comment,
		ImportStartMark:        config.ImportStartMark,
		ImportEndMark:          config.ImportEndMark,
		CustomImportStartMark:  config.CustomImportStartMark,
		CustomImportContent:    "",
		CustomImportEndMark:    config.CustomImportEndMark,
		StructStartMark:        config.StructStartMark,
		StructEndMark:          config.StructEndMark,
		CustomStructStartMark:  config.CustomStructStartMark,
		CustomStructContent:    "",
		CustomStructEndMark:    config.CustomStructEndMark,
		ServiceStartMark:       config.ServiceStartMark,
		ServiceEndMark:         config.ServiceEndMark,
		CustomServiceStartMark: config.CustomServiceStartMark,
		CustomServiceContent:   "",
		CustomServiceEndMark:   config.CustomServiceEndMark,
		DbTables:               utils.FilterTables(global.Schema.Tables, config.C.Api.Tables, utils.MergeSlice(config.C.Api.IgnoreTables, conf.BaseIgnoreTables)),
		DbIgoreFieldsName:      utils.MergeSlice(config.C.Api.IgnoreColumns, conf.BaseIgnoreColumns),
		Imports:                []*internal.Import{},
		BaseStructs:            internal.GenBaseStructCollection(),
		Structs:                []*internal.Struct{},
		Services:               []*internal.Service{},
	}
}

func (s *ApiModule) Run() error {
	//分两种情况，是否为多文件模式
	if !s.Multiple {
		return s.Generate(s.DbTables...)
	}
	//需要根据table生成多个api文件
	//通过控制tables来实现
	imports := make([]*internal.Import, 0)
	s.CurrentIsCoreFile = false
	for _, table := range s.DbTables {
		s.OutFileName = getOutFilename(table.Name)
		imports = append(imports, internal.NewImport(s.OutFileName))
		if err := s.Generate(table); err != nil {
			return err
		}
	}

	s.CurrentIsCoreFile = true
	s.OutFileName = getOutFilename(config.C.Api.ServiceName)
	s.Imports = imports

	err := s.Generate()
	if err != nil {
		tools.Error("generate api file failed.")
		return err
	}
	tools.Success("generate api file success.")
	return nil
}

func (s *ApiModule) Generate(tables ...sql.Table) error {
	if err := utils.BackupUserFile(s.OutFileName); err != nil {
		return err
	}
	s.DbTables = tables
	err := s.Init()
	if err != nil {
		return err
	}
	content, err := s.Render()
	if err != nil {
		return err
	}
	err = s.WriteFile(content)
	if err != nil {
		return err
	}
	return nil
}

func (s *ApiModule) WriteFile(content string) error {
	return tools.WriteFile(s.OutFileName, content)
}

func (s *ApiModule) Render() (string, error) {
	tmpl, err := tools.Template("api").Funcs(template.FuncMap{
		"NeedRenderStruct": func(isMultiple, isCoreFile bool) bool {
			return !isMultiple || (isMultiple && !isCoreFile)
		},
		"NeedRenderService": func(isMultiple, isCoreFile bool) bool {
			return !isMultiple || (isMultiple && !isCoreFile)
		},
	}).Parse(s.TplContent)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, *s)
	if err != nil {
		return "", err
	}
	content := utils.FormatContent(buf.String())
	return content, nil
}

func (s *ApiModule) Init() (err error) {

	if err = s.initTplContent(); err != nil {
		return
	}
	if err = s.initOldContent(); err != nil {
		return
	}
	if err = s.initCustomImportContent(); err != nil {
		return
	}
	if err = s.initCustomStructContent(); err != nil {
		return
	}
	if err = s.initCustomServiceContent(); err != nil {
		return
	}
	if err = s.initImports(); err != nil {
		return
	}
	if err = s.initStructs(); err != nil {
		return
	}
	if err = s.initServices(); err != nil {
		return
	}
	return nil
}

func (s *ApiModule) initTplContent() error {
	// 判断当前目录下是否存在./template/api.tpl文件
	protoTplPath := "./template/api.tpl"
	exist, err := tools.PathExists(protoTplPath)
	if err != nil {
		return err
	}
	if exist {
		//如果存在，则读取该内容作为模板
		s.TplContent, err = tools.ReadFile(protoTplPath)
		return err
	}
	if s.TplContent == "" {
		s.TplContent = conf.TplContent
	}
	return nil
}

func (s *ApiModule) initOldContent() (err error) {
	exists, err := tools.PathExists(s.OutFileName)
	if err != nil {
		return err
	}
	if exists {
		s.OldContent, err = tools.ReadFile(s.OutFileName)
	} else {
		s.OldContent = ""
	}
	return
}

func (s *ApiModule) initCustomImportContent() error {
	customImportContent, err := utils.GetMarkContent(s.CustomImportStartMark, s.CustomImportEndMark, s.OldContent)
	if err != nil {
		return err
	}
	s.CustomImportContent = customImportContent
	return nil
}

func (s *ApiModule) initCustomStructContent() error {
	customStructContent, err := utils.GetMarkContent(s.CustomStructStartMark, s.CustomStructEndMark, s.OldContent)
	if err != nil {
		return err
	}
	s.CustomStructContent = customStructContent
	return nil
}

func (s *ApiModule) initCustomServiceContent() error {
	customServiceContent, err := utils.GetMarkContent(s.CustomServiceStartMark, s.CustomServiceEndMark, s.OldContent)
	if err != nil {
		return err
	}
	s.CustomServiceContent = customServiceContent
	return nil
}

func (s *ApiModule) initImports() (err error) {
	//没有需要导入的包
	return
}

func (s *ApiModule) initStructs() (err error) {
	//整理出所有的消息
	structs := []*internal.Struct{}
	structsMap := map[string]bool{}
	for _, table := range s.DbTables {
		tableName := tools.ToCamel(table.Name)
		if _, ok := structsMap[tableName]; ok {
			continue
		}

		cols := utils.FilterIgnoreFields(table.GetFields(), s.DbIgoreFieldsName)

		fields := []*internal.StructField{}
		for _, field := range cols {
			fields = append(fields, internal.NewStructField(field.UpperCamelCaseName, field.Type, "json", field.Name, "", field.Comment))
		}
		structItem := internal.NewStruct(tableName, "json", table.Comment, fields)
		commonStructs := structItem.GenCommonStructs()
		structs = append(structs, commonStructs...)
	}
	s.Structs = structs
	// sort.Sort(s.Structs)
	return
}

func (s *ApiModule) initServices() (err error) {
	//整理出所有的服务
	services := []*internal.Service{}
	servicesMap := map[string]bool{}
	for _, table := range s.DbTables {
		tableName := tools.ToCamel(table.Name)
		if _, ok := servicesMap[tableName]; ok {
			continue
		}
		service := internal.NewService(tableName, table.Comment)
		services = append(services, service)
	}

	s.Services = services
	// sort.Sort(s.Services)
	return
}
