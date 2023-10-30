package cmd

import (
	"fmt"
	"os"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/tools"
	"github.com/spf13/cobra"
)

var IsDev bool

var colorId int = 30

func getCmdColor() int {
	colorId += 1
	if colorId < 30 || (colorId > 37 && colorId < 90) || colorId > 97 {
		return 30
	}
	return colorId
}

func setColorizeHelp(template string) string {
	// 在这里使用ANSI转义码添加颜色
	// 参考ANSI转义码文档：https://en.wikipedia.org/wiki/ANSI_escape_code
	return fmt.Sprintf("\033[%dm%s\033[0m", getCmdColor(), template)
}

func setColorPart(part string) string {
	return fmt.Sprintf("\033[%dm%s", getCmdColor(), part)
}

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i"},
	Short:   "Create the default " + config.ProjectName + " configuration file in the current directory",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("%s requires at least one argument", cmd.CommandPath())
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var initConfigCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"c"},
	Short:   "Create the default " + config.ProjectName + " configuration file in the current directory, or specified directory",
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.C.CreateYaml(); err != nil {
			tools.Warning(err.Error())
			os.Exit(1)
		}
		tools.Success("Done.")
	},
}

func init() {
	initConfigCmd.PersistentFlags().StringVar(&config.InitConfSrc, "dir", config.DefaultConfigFileName, "file location for yaml configuration")
	initCmd.AddCommand(initConfigCmd)
	initCmd.SetHelpTemplate(setColorizeHelp(initCmd.HelpTemplate()))

	rootCmd.AddCommand(initCmd)
}
