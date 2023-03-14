package config

import (
	"github.com/licat233/genzero/tools"
	"gopkg.in/yaml.v3"
)

var C *Config

type Config struct {
	DatabaseConfig *DatabaseConfig `yaml:"database"`
	ApiConfig      *ApiConfig      `yaml:"api"`
	PbConfig       *PbConfig       `yaml:"protobuf"`
	ModelConfig    *ModelConfig    `yaml:"model"`
}

func New() *Config {
	return &Config{
		DatabaseConfig: &DatabaseConfig{},
		ApiConfig:      &ApiConfig{},
		PbConfig:       &PbConfig{},
		ModelConfig:    &ModelConfig{},
	}
}

func (c *Config) CreateYaml() error {
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
