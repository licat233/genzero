package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/global"
	"github.com/licat233/genzero/modules/api"
	"github.com/licat233/genzero/modules/logic"
	"github.com/licat233/genzero/modules/model"
	"github.com/licat233/genzero/modules/pb"
	"github.com/licat233/genzero/tools"
	"github.com/licat233/genzero/upgrade"
	"github.com/spf13/cobra"
)

var modulesGroup = &cobra.Group{
	ID:    "modules",
	Title: "modules:",
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Print the version number of " + config.ProjectName,
	Run: func(cmd *cobra.Command, args []string) {
		tools.Success("current version: " + config.ProjectVersion)
	},
}

var upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Aliases: []string{"up", "u"},
	Short:   "Upgrade " + config.ProjectName + " to latest version",
	Run: func(cmd *cobra.Command, args []string) {
		upgrade.Upgrade()
	},
}

var startCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"run"},
	Short:   "Use yaml file configuration to start " + config.ProjectName,
	Run: func(cmd *cobra.Command, args []string) {
		config.UseConf = true
		runByConf()
	},
}

var rootCmd = &cobra.Command{
	Use:   config.ProjectName,
	Short: "This is a tool to generate gozero service based on mysql",
	Long:  fmt.Sprintf("This is a tool to generate gozero service based on mysql.\nThe goctl tool must be installed before use.\ncurrent version: %s\nGithub: https://github.com/licat233/genzero", config.ProjectVersion),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("%s requires at least one argument", cmd.CommandPath())
		}
		return nil
	},
	Version: config.ProjectVersion,
	Run: func(cmd *cobra.Command, args []string) {
		// run()
	},
}

func runByConf() {
	if err := Initialize(); err != nil {
		tools.Error("Failed to initialize: " + err.Error())
		os.Exit(1)
	}
	tasks := make([]tools.TaskFunc, 0)
	if config.C.Model.Status {
		tasks = append(tasks, func() error {
			return model.New().Run()
		})
	}

	if config.C.Pb.Status {
		tasks = append(tasks, func() error {
			return pb.New().Run()
		})
	}

	if config.C.Api.Status {
		tasks = append(tasks, func() error {
			return api.New().Run()
		})
	}
	if config.C.Logic.Status {
		tasks = append(tasks, func() error {
			return logic.New().Run()
		})
	}

	err := tools.RunConcurrentTasks(tasks)
	if err != nil {
		tools.Error(err.Error())
		os.Exit(1)
	}

	tools.Success("Done.")
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&IsDev, "dev", false, "dev mode, print error message")
	rootCmd.PersistentFlags().StringVar(&config.ConfSrc, "conf", config.DefaultConfigFileName, "file location for yaml configuration")
	rootCmd.PersistentFlags().StringVar(&config.C.DB.DSN, "dsn", "", "data source name (DSN) to use when connecting to the database")
	rootCmd.PersistentFlags().StringVar(&config.C.DB.Src, "src", "", "sql file to use when connecting to the database")
	rootCmd.PersistentFlags().StringSliceVar(&config.C.DB.Tables, "tables", []string{}, "need to generate tables, default is all tables，split multiple value by ','")
	rootCmd.PersistentFlags().StringSliceVar(&config.C.DB.IgnoreTables, "ignore_tables", []string{}, "ignore table string, default is none，split multiple value by ','")
	rootCmd.PersistentFlags().StringSliceVar(&config.C.DB.IgnoreColumns, "ignore_columns", []string{}, "ignore column string, default is none，split multiple value by ','")

	rootCmd.AddGroup(modulesGroup)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(upgradeCmd)
	rootCmd.AddCommand(startCmd)

	rootCmd.SetHelpTemplate(greenColorizeHelp(rootCmd.HelpTemplate()))
}

func greenColorizeHelp(template string) string {
	// 在这里使用ANSI转义码添加颜色
	// 参考ANSI转义码文档：https://en.wikipedia.org/wiki/ANSI_escape_code
	// 将标题文本设置为绿色
	template = "\033[32m" + template + "\033[0m"
	return template
}

func Initialize() error {
	if config.ConfSrc == "." {
		config.ConfSrc = config.DefaultConfigFileName
		config.UseConf = true
	}
	if !config.UseConf {
		config.UseConf = strings.HasSuffix(config.ConfSrc, ".yaml") || strings.HasSuffix(config.ConfSrc, ".yam")
	}
	if config.UseConf {
		if err := config.C.ConfigureByYaml(); err != nil {
			return fmt.Errorf("read config faild: %v", err)
		}
	}

	if err := config.C.Validate(); err != nil {
		return err
	}

	if err := global.InitSchema(); err != nil {
		return err
	}
	return nil
}

func Execute() {
	defer func() {
		if !IsDev {
			if err := recover(); err != nil {
				tools.Warning(fmt.Sprintf("%v", err))
			}
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		tools.Warning(err.Error())
		os.Exit(1)
	}
}
