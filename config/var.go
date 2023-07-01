package config

import "strings"

var (
	UseConf     bool
	ConfSrc     string
	InitConfSrc string
)

const (
	// CurrentVersion 当前项目版本
	CurrentVersion = "v1.1.3-alpine.4"

	// ProjectName 当前项目名称
	ProjectName = "genzero"

	// ProjectURL 当前项目地址
	ProjectURL = "https://github.com/licat233/" + ProjectName
	// ProjectInfoURL 当前项目的信息接口
	ProjectInfoURL = "https://api.github.com/repos/licat233/" + ProjectName + "/releases/latest"

	// DefaultConfigFileName 配置文件名称
	DefaultConfigFileName = ProjectName + ".yaml"
	// OldConfigFileName 旧的配置文件名称
	OldConfigFileName = ProjectName + "Config.yaml"

	CamelCase      = "GenZero"
	LowerCamelCase = "genZero"
	SnakeCase      = "gen_zero"

	UpdatedFileMsg = "已更新文件"
	CreatedFileMsg = "已创建文件"
)

var StyleList = strings.Join([]string{
	LowerCamelCase,
	CamelCase,
	SnakeCase,
}, ",")

var (
	InfoStartMark, InfoEndMark       = GetMark("Info")
	ImportStartMark, ImportEndMark   = GetMark("Import")
	StructStartMark, StructEndMark   = GetMark("Struct")
	EnumStartMark, EnumEndMark       = GetMark("Enum")
	MessageStartMark, MessageEndMark = GetMark("Message")
	ServiceStartMark, ServiceEndMark = GetMark("Service")

	BaseFuncsStartMark, BaseFuncsEndMark = GetBaseMark("Funcs")

	CustomImportStartMark, CustomImportEndMark   = GetCustomMark("import")
	CustomStructStartMark, CustomStructEndMark   = GetCustomMark("struct")
	CustomEnumStartMark, CustomEnumEndMark       = GetCustomMark("enum")
	CustomMessageStartMark, CustomMessageEndMark = GetCustomMark("message")
	CustomServiceStartMark, CustomServiceEndMark = GetCustomMark("service")
)
