package config

import (
	"errors"
	"fmt"

	"github.com/licat233/genzero/tools"
	"gopkg.in/yaml.v3"
)

func GetMark(name string) (startMark string, endMark string) {
	startMark = "// ------------------------------ " + name + " Start ------------------------------"
	endMark = "// ------------------------------ " + name + " End ------------------------------"
	return
}

func GetCustomMark(name string) (startMark string, endMark string) {
	startMark = "//[custom " + name + " start]"
	endMark = "//[custom " + name + " end]"
	return
}

func GetBaseMark(name string) (startMark string, endMark string) {
	startMark = "//[base " + name + " start]"
	endMark = "//[base " + name + " end]"
	return
}

func ConfigureByYaml(filename string, config *Config) error {
	exists, err := tools.PathExists(ConfSrc)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("config file (%s) not exists, please create it first, command: %s init config", ConfSrc, ProjectName)
	}
	if filename == "" {
		filename = ConfSrc
	}
	data, err := tools.ReadFile(filename)
	if err != nil {
		return err
	}
	if config == nil {
		return errors.New("config is nil")
	}
	err = yaml.Unmarshal([]byte(data), config)
	if err != nil {
		tools.Error("parse config file (%s) failed", filename)
		return err
	}
	return nil
}
