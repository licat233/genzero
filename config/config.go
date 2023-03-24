package config

import (
	"github.com/licat233/genzero/tools"
	"gopkg.in/yaml.v3"
)

var C *Config

type Config struct {
	DB    *DatabaseConfig `yaml:"database"`
	Api   *ApiConfig      `yaml:"api"`
	Pb    *PbConfig       `yaml:"proto"`
	Model *ModelConfig    `yaml:"model"`
	Logic *LogicConfig    `yaml:"logic"`
}

func New() *Config {
	return &Config{
		DB:    &DatabaseConfig{},
		Api:   &ApiConfig{},
		Pb:    &PbConfig{},
		Model: &ModelConfig{},
		Logic: &LogicConfig{},
	}
}

func (c *Config) CreateYaml() error {
	yamlBytes, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	//先判断是否已经存在，存在则合并内容
	exists, err := tools.PathExists(InitConfSrc)
	if err != nil {
		return err
	}
	if exists {
		//判断内容是否为空
		content, err := tools.ReadFile(InitConfSrc)
		if err != nil {
			return err
		}
		if content != "" {
			// tools.Warning("%s already exists, please empty the file content or delete the file first.", DefaultConfigFileName)
			return tools.MergeYamlContent(InitConfSrc, yamlBytes)
		}
	}

	content, err := tools.SortYamlContent(string(yamlBytes))
	if err != nil {
		return err
	}

	return tools.WriteFile(InitConfSrc, content)
}

func (c *Config) ConfigureByYaml() error {
	err := ConfigureByYaml(ConfSrc, c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) Validate() error {
	if c.DB != nil {
		if err := C.DB.Validate(); err != nil {
			return err
		}
	}
	if c.Api != nil && C.Api.Status {
		if err := C.Api.Validate(); err != nil {
			return err
		}
	}
	if c.Pb != nil && C.Pb.Status {
		if err := C.Pb.Validate(); err != nil {
			return err
		}
	}
	if c.Model != nil && c.Model.Status {
		if err := C.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}
