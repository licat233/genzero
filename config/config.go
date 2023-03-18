package config

import (
	"github.com/licat233/genzero/tools"
	"gopkg.in/yaml.v3"
)

var C *Config

type Config struct {
	DatabaseConfig *DatabaseConfig `yaml:"database"`
	ApiConfig      *ApiConfig      `yaml:"api"`
	PbConfig       *PbConfig       `yaml:"proto"`
	ModelConfig    *ModelConfig    `yaml:"model"`
	LogicConfig    *LogicConfig    `yaml:"logic"`
}

func New() *Config {
	return &Config{
		DatabaseConfig: &DatabaseConfig{},
		ApiConfig:      &ApiConfig{},
		PbConfig:       &PbConfig{},
		ModelConfig:    &ModelConfig{},
		LogicConfig:    &LogicConfig{},
	}
}

func (c *Config) CreateYaml() error {
	//先判断是否已经存在
	exists, err := tools.PathExists(DefaultConfigFileName)
	if err != nil {
		return err
	}
	if exists {
		//判断内容是否为空
		content, err := tools.ReadFile(DefaultConfigFileName)
		if err != nil {
			return err
		}
		if content != "" {
			tools.Warning("%s already exists, please empty the file content or delete the file first.", DefaultConfigFileName)
			return nil
		}
	}
	yamlBytes, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return tools.WriteFile(DefaultConfigFileName, string(yamlBytes))
}

func (c *Config) ConfigureByYaml() error {
	err := ConfigureByYaml(DefaultConfigFileName, c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) Validate() error {
	if c.DatabaseConfig != nil {
		if err := C.DatabaseConfig.Validate(); err != nil {
			return err
		}
	}
	if c.ApiConfig != nil && C.ApiConfig.Status {
		if err := C.ApiConfig.Validate(); err != nil {
			return err
		}
	}
	if c.PbConfig != nil && C.PbConfig.Status {
		if err := C.PbConfig.Validate(); err != nil {
			return err
		}
	}
	if c.ModelConfig != nil && c.ModelConfig.Status {
		if err := C.ModelConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
