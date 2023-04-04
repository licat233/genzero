package config

import (
	"errors"
	"strings"

	"github.com/licat233/genzero/tools"
)

type DatabaseConfig struct {
	DSN           string   `yaml:"dsn"`           // data source name (DSN) to use when connecting to the database
	Src           string   `yaml:"src"`           // sql file to use when connecting to the database
	Tables        []string `yaml:"tables"`        // need to generate tables, default is all tables，split multiple value by ","
	IgnoreTables  []string `yaml:"ignoreTables"`  // ignore table string, default is none，split multiple value by ","
	IgnoreColumns []string `yaml:"ignoreColumns"` // ignore column string, default is none，split multiple value by ","
}

func (c *DatabaseConfig) Validate() error {
	if c.DSN == "" && c.Src == "" {
		return errors.New("database dsn or src must be set")
	}
	return nil
}

// api配置
type ApiConfig struct {
	Status        bool     `yaml:"status"` // generate api
	Style         string   `yaml:"style"`
	Jwt           string   `yaml:"jwt"`
	Middleware    []string `yaml:"middleware"`
	Prefix        string   `yaml:"prefix"`
	Multiple      bool     `yaml:"multiple"`
	Dir           string   `yaml:"dir"` // api output directory
	ServiceName   string   `yaml:"serviceName"`
	Tables        []string `yaml:"tables"`        // need to generate tables, default is all tables，split multiple value by ","
	IgnoreTables  []string `yaml:"ignoreTables"`  // ignore table string, default is none，split multiple value by ","
	IgnoreColumns []string `yaml:"ignoreColumns"` // ignore column string, default is none，split multiple value by ","
}

func (c *ApiConfig) Validate() error {
	if c.ServiceName == "" {
		return errors.New("serviceName must be set")
	}
	if len(c.Middleware) != 0 {
		for i := range c.Middleware {
			c.Middleware[i] = tools.ToCamel(c.Middleware[i])
		}
	}
	c.ServiceName = strings.TrimRight(tools.ToLowerCamel(c.ServiceName), "Api")
	return nil
}

// pb配置
type PbConfig struct {
	Status        bool     `yaml:"status"` // generate proto
	Style         string   `yaml:"style"`
	Package       string   `yaml:"package"`
	GoPackage     string   `yaml:"goPackage"`
	Multiple      bool     `yaml:"multiple"`
	Dir           string   `yaml:"dir"` // proto output directory
	ServiceName   string   `yaml:"serviceName"`
	Tables        []string `yaml:"tables"`        // need to generate tables, default is all tables，split multiple value by ","
	IgnoreTables  []string `yaml:"ignoreTables"`  // ignore table string, default is none，split multiple value by ","
	IgnoreColumns []string `yaml:"ignoreColumns"` // ignore column string, default is none，split multiple value by ","
}

func (c *PbConfig) Validate() error {
	if c.ServiceName == "" {
		return errors.New("serviceName must be set")
	}
	if c.Package == "" {
		c.Package = tools.ToLowerCamel(c.ServiceName) + "_proto"
	}
	if c.GoPackage == "" {
		c.GoPackage = "./" + tools.ToLowerCamel(c.ServiceName) + "_pb"
	} else if !strings.HasSuffix(c.GoPackage, "/") && !strings.HasSuffix(c.GoPackage, "./") {
		c.GoPackage = "./" + c.GoPackage
	}
	c.ServiceName = tools.ToCamel(c.ServiceName)
	return nil
}

// model配置global.Config
type ModelConfig struct {
	Status        bool     `yaml:"status"`        // generate model
	Dir           string   `yaml:"dir"`           // model output directory
	Tables        []string `yaml:"tables"`        // need to generate tables, default is all tables，split multiple value by ","
	IgnoreTables  []string `yaml:"ignoreTables"`  // ignore table string, default is none，split multiple value by ","
	IgnoreColumns []string `yaml:"ignoreColumns"` // ignore column string, default is none，split multiple value by ","
}

func (c *ModelConfig) Validate() error {
	return nil
}

// logic配置global.Config
type LogicConfig struct {
	Status bool `yaml:"status"` // modify logic
	Api    struct {
		Status bool `yaml:"status"` // generate api
		UseRpc bool `yaml:"useRpc"` // use rpc
		// RpcMultiple  bool     `yaml:"rpcMultiple"`  // is multiple rpc
		Style        string   `yaml:"style"`        // style
		Dir          string   `yaml:"dir"`          // api logic directory
		Tables       []string `yaml:"tables"`       // need to generate tables, default is all tables，split multiple value by ","
		IgnoreTables []string `yaml:"ignoreTables"` // ignore table string, default is none，split multiple value by ","
		// IgnoreColumns []string `yaml:"ignoreColumns"` // ignore column string, default is none，split multiple value by ","
	} `yaml:"api"`
	Rpc struct {
		Status       bool     `yaml:"status"`       // generate rpc
		Multiple     bool     `yaml:"multiple"`     // is multiple
		Style        string   `yaml:"style"`        // style
		Dir          string   `yaml:"dir"`          // rpc logic directory
		Tables       []string `yaml:"tables"`       // need to generate tables, default is all tables，split multiple value by ","
		IgnoreTables []string `yaml:"ignoreTables"` // ignore table string, default is none，split multiple value by ","
		// IgnoreColumns []string `yaml:"ignoreColumns"` // ignore column string, default is none，split multiple value by ","
	} `yaml:"rpc"`
}

func (c *LogicConfig) Validate() error {
	return nil
}
