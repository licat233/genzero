package config

import (
	"errors"
	"fmt"
	"os"

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

func RenameOldConfigFile() error {
	// 为了兼容旧的配置文件
	// 检测旧配置文件是否存在
	exists, err := tools.PathExists(OldConfigFileName)
	if err != nil {
		return err
	}
	if exists {
		// 同时也要检测新的配置文件是否存在
		exists, err := tools.PathExists(DefaultConfigFileName)
		if err != nil {
			return err
		}
		if !exists {
			// 如果不存在新的配置文件，则更改旧配置文件名为新的配置文件名
			err := os.Rename(OldConfigFileName, DefaultConfigFileName)
			if err != nil {
				return err
			}
		}
		// 如果同时存在新旧配置文件
		// 则优先使用新的配置文件，旧配置文件留着不动

		// 此时，如果 ConfSrc 等于 OldConfigFileName，则应该赋值为 DefaultConfigFileName
		if ConfSrc == OldConfigFileName {
			ConfSrc = DefaultConfigFileName
		}
		// 同理
		if InitConfSrc == OldConfigFileName {
			ConfSrc = DefaultConfigFileName
		}
	}
	return nil
}
