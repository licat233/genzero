package pb

import (
	"bytes"
	"fmt"
	"path"
	"sort"
	"text/template"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/global"
	"github.com/licat233/genzero/modules/pb/conf"
	"github.com/licat233/genzero/modules/pb/internal"
	"github.com/licat233/genzero/modules/utils"
	"github.com/licat233/genzero/sql"
	"github.com/licat233/genzero/tools"
)

type PbModule struct {
	ProjectAuthor  string
	ProjectName    string
	ProjectAddr    string
	ProjectVersion string

	TplContent     string
	OutFileName    string
	OldContent     string
	Multiple       bool
	ServiceName    string
	ServiceComment string
	Package        string
	GoPackage      string

	ImportStartMark        string
	ImportEndMark          string
	CustomImportStartMark  string
	CustomImportContent    string
	CustomImportEndMark    string
	EnumStartMark          string
	EnumEndMark            string
	CustomEnumStartMark    string
	CustomEnumContent      string
	CustomEnumEndMark      string
	MessageStartMark       string
	MessageEndMark         string
	CustomMessageStartMark string
	CustomMessageContent   string
	CustomMessageEndMark   string
	ServiceStartMark       string
	ServiceEndMark         string
	CustomServiceStartMark string
	CustomServiceContent   string
	CustomServiceEndMark   string

	DbTables          sql.TableCollection
	DbIgoreFieldsName []string

	Imports      internal.ImportCollection
	Enums        internal.EnumCollection
	BaseMessages internal.MessageCollection
	Messages     internal.MessageCollection
	Services     internal.ServiceCollection
}

func getOutFilename(name string) string {
	filename := utils.ConvertStringStyle(config.C.Pb.FileStyle, name)
	return path.Join(config.C.Pb.Dir, filename+".proto")
}

func New() *PbModule {
	return &PbModule{
		ProjectAuthor:          tools.GetCurrentUserName(),
		ProjectName:            config.ProjectName,
		ProjectAddr:            config.ProjectURL,
		ProjectVersion:         config.CurrentVersion,
		TplContent:             conf.TplContent,
		OutFileName:            getOutFilename(config.C.Pb.ServiceName),
		OldContent:             "",
		Multiple:               config.C.Pb.Multiple,
		ServiceName:            config.C.Pb.ServiceName,
		ServiceComment:         global.Schema.Comment,
		Package:                config.C.Pb.Package,
		GoPackage:              config.C.Pb.GoPackage,
		ImportStartMark:        config.ImportStartMark,
		ImportEndMark:          config.ImportEndMark,
		CustomImportStartMark:  config.CustomImportStartMark,
		CustomImportContent:    "",
		CustomImportEndMark:    config.CustomImportEndMark,
		EnumStartMark:          config.EnumStartMark,
		EnumEndMark:            config.EnumEndMark,
		CustomEnumStartMark:    config.CustomEnumStartMark,
		CustomEnumContent:      "",
		CustomEnumEndMark:      config.CustomEnumEndMark,
		MessageStartMark:       config.MessageStartMark,
		MessageEndMark:         config.MessageEndMark,
		CustomMessageStartMark: config.CustomMessageStartMark,
		CustomMessageContent:   "",
		CustomMessageEndMark:   config.CustomMessageEndMark,
		ServiceStartMark:       config.ServiceStartMark,
		ServiceEndMark:         config.ServiceEndMark,
		CustomServiceStartMark: config.CustomServiceStartMark,
		CustomServiceContent:   "",
		CustomServiceEndMark:   config.CustomServiceEndMark,
		DbTables:               utils.FilterTables(global.Schema.Tables, config.C.Pb.Tables, utils.MergeSlice(config.C.Pb.IgnoreTables, conf.BaseIgnoreTables)),
		DbIgoreFieldsName:      utils.MergeSlice(config.C.Pb.IgnoreColumns, conf.BaseIgnoreColumns),
		Imports:                []*internal.Import{},
		Enums:                  []*internal.Enum{},
		BaseMessages:           []*internal.Message{},
		Messages:               []*internal.Message{},
		Services:               []*internal.Service{},
	}
}

func (s *PbModule) Run() error {
	err := s.Generate()
	if err != nil {
		tools.Error("generate proto file faild.")
		return err
	}
	tools.Success("generate proto file success.")
	return nil
}

func (s *PbModule) Generate() error {
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

func (s *PbModule) WriteFile(content string) error {
	return tools.WriteFile(s.OutFileName, content)
}

func (s *PbModule) Render() (string, error) {
	tmpl, err := tools.Template("proto").Funcs(template.FuncMap{
		"isEmpty": func(str string) bool {
			return tools.TrimSpace(str) == ""
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

func (s *PbModule) Init() (err error) {
	if global.Schema == nil {
		return fmt.Errorf("schema is nil")
	}
	if err = s.initTplContent(); err != nil {
		return
	}
	if err = s.initOldContent(); err != nil {
		return
	}
	if err = s.initCustomImportContent(); err != nil {
		return
	}
	if err = s.initCustomEnumContent(); err != nil {
		return
	}
	if err = s.initCustomMessageContent(); err != nil {
		return
	}
	if err = s.initCustomServiceContent(); err != nil {
		return
	}
	if err = s.initImports(); err != nil {
		return
	}
	if err = s.initEnums(); err != nil {
		return
	}
	if err = s.initBaseMessages(); err != nil {
		return
	}
	if err = s.initMessages(); err != nil {
		return
	}
	if err = s.initServices(); err != nil {
		return
	}
	return nil
}

func (s *PbModule) initTplContent() error {
	// 判断当前目录下是否存在./template/proto.tpl文件
	protoTplPath := "./template/proto.tpl"
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

func (s *PbModule) initOldContent() (err error) {
	exists, err := tools.PathExists(s.OutFileName)
	if err != nil {
		return err
	}
	if exists {
		s.OldContent, err = tools.ReadFile(s.OutFileName)
	}
	return
}

func (s *PbModule) initCustomImportContent() error {
	customImportContent, err := utils.GetMarkContent(s.CustomImportStartMark, s.CustomImportEndMark, s.OldContent)
	if err != nil {
		return err
	}
	s.CustomImportContent = customImportContent
	return nil
}

func (s *PbModule) initCustomEnumContent() error {
	customEnumContent, err := utils.GetMarkContent(s.CustomEnumStartMark, s.CustomEnumEndMark, s.OldContent)
	if err != nil {
		return err
	}
	s.CustomEnumContent = customEnumContent
	return nil
}

func (s *PbModule) initCustomMessageContent() error {
	customMessageContent, err := utils.GetMarkContent(s.CustomMessageStartMark, s.CustomMessageEndMark, s.OldContent)
	if err != nil {
		return err
	}
	s.CustomMessageContent = customMessageContent
	return nil
}

func (s *PbModule) initCustomServiceContent() error {
	customServiceContent, err := utils.GetMarkContent(s.CustomServiceStartMark, s.CustomServiceEndMark, s.OldContent)
	if err != nil {
		return err
	}
	s.CustomServiceContent = customServiceContent
	return nil
}

func (s *PbModule) initImports() (err error) {
	//没有需要导入的包
	return
}

func (s *PbModule) initEnums() (err error) {
	//整理出所有的枚举
	enumsMap := map[string]bool{}
	enums := []*internal.Enum{}
	for _, table := range s.DbTables {
		for _, enum := range table.Enums {
			enumName := tools.ToCamel(enum.Name)
			if _, ok := enumsMap[enumName]; ok {
				continue
			}

			fields := []*internal.EnumField{}
			for _, field := range enum.Fields {
				fields = append(fields, internal.NewEnumField(field.Name, field.Tag))
			}
			enums = append(enums, &internal.Enum{
				Name:    enumName,
				Comment: enum.Comment,
				Fields:  fields,
			})
			enumsMap[enumName] = true
		}
	}
	for _, enum := range s.Enums {
		if _, ok := enumsMap[enum.Name]; ok {
			continue
		}
		enums = append(enums, enum)
		enumsMap[enum.Name] = true
	}
	s.Enums = enums
	sort.Sort(s.Enums)
	return
}

func (s *PbModule) initBaseMessages() (err error) {
	s.BaseMessages = internal.GenBaseMessages()
	return nil
}

func (s *PbModule) initMessages() (err error) {
	//整理出所有的消息
	messages := []*internal.Message{}
	messagesMap := map[string]bool{}
	for _, table := range s.DbTables {
		tableName := tools.ToCamel(table.Name)
		if _, ok := messagesMap[tableName]; ok {
			continue
		}

		cols := utils.FilterIgnoreFields(table.GetFields(), s.DbIgoreFieldsName)

		fields := []*internal.MessageField{}

		tag := 1
		for _, field := range cols {
			t, err := GoTypeToProtoType(field.Type)
			if err != nil {
				return err
			}
			fields = append(fields, internal.NewMessageField(t, field.Name, tag, field.Comment))
			tag++
		}

		message := internal.NewMessage(tableName, table.Comment, fields)
		commonMessages := message.GenCommonMessages()
		messages = append(messages, commonMessages...)
	}
	for _, message := range s.Messages {
		if _, ok := messagesMap[message.Name]; ok {
			continue
		}
		messages = append(messages, message)
		messagesMap[message.Name] = true
	}

	s.Messages = messages
	sort.Sort(s.Messages)
	return
}

func (s *PbModule) initServices() (err error) {
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

	for _, service := range s.Services {
		if _, ok := servicesMap[service.Name]; ok {
			continue
		}
		services = append(services, service)
		servicesMap[service.Name] = true
	}
	s.Services = services
	sort.Sort(s.Services)
	return
}

func GoTypeToProtoType(typeName string) (string, error) {
	switch typeName {
	case "bool":
		return "bool", nil
	case "int", "int8", "int16", "int32":
		return "int32", nil
	case "int64":
		return "int64", nil
	case "uint", "uint8", "uint16", "uint32":
		return "uint32", nil
	case "uint64":
		return "uint64", nil
	case "float32":
		return "float", nil
	case "float64":
		return "double", nil
	case "string":
		return "string", nil
	case "bytes", "[]byte":
		return "bytes", nil
	case "map":
		return "map", nil
	case "struct":
		return "message", nil
	case "interface{}", "any":
		return "any", nil
	case "time.Time":
		return "int64", nil
	default:
		return "", fmt.Errorf("unsupported type: %T", typeName)
	}
}
