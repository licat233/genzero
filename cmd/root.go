package cmd

import (
	"fmt"
	"log"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/core/api"
	"github.com/licat233/genzero/core/model"
	"github.com/licat233/genzero/core/pb"
	"github.com/licat233/genzero/global"
	"github.com/licat233/genzero/tools"
	"github.com/licat233/genzero/update"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Print the version number of " + config.ProjectName,
	Run: func(cmd *cobra.Command, args []string) {
		tools.Println("current version: " + config.CurrentVersion)
	},
}

var upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Aliases: []string{"up", "u"},
	Short:   "Upgrade " + config.ProjectName + " to latest version",
	Run: func(cmd *cobra.Command, args []string) {
		update.Update()
	},
}

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i"},
	Short:   "Create the default " + config.ProjectName + " configuration file in the current directory",
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.C.CreateYaml(); err != nil {
			log.Fatalln(err)
		}
		tools.Println("Done.")
	},
}

var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "Generate model code",
	Run: func(cmd *cobra.Command, args []string) {
		config.C.ModelConfig.Status = true
		Initialize()
		if err := model.New().Run(); err != nil {
			log.Fatal(err)
		}
		tools.Println("Done.")
	},
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Generate .api files",
	Run: func(cmd *cobra.Command, args []string) {
		config.C.ApiConfig.Status = true
		Initialize()
		if err := api.New().Run(); err != nil {
			log.Fatal(err)
		}
		tools.Println("Done.")
	},
}

var pbCmd = &cobra.Command{
	Use:   "pb",
	Short: "Generate .proto files",
	Run: func(cmd *cobra.Command, args []string) {
		config.C.PbConfig.Status = true
		Initialize()
		if err := pb.New().Run(); err != nil {
			log.Fatal(err)
		}
		tools.Println("Done.")
	},
}

var yamlCmd = &cobra.Command{
	Use:   "yaml",
	Short: "Use yaml file configuration",
	Run: func(cmd *cobra.Command, args []string) {
		err := config.C.ConfigureByYaml()
		if err != nil {
			log.Fatal(err)
		}
		Initialize()
		if config.C.ModelConfig.Status {
			if err := model.New().Run(); err != nil {
				log.Fatal(err)
			}
		}
		if config.C.ApiConfig.Status {
			if err := api.New().Run(); err != nil {
				log.Fatal(err)
			}
		}
		if config.C.PbConfig.Status {
			if err := pb.New().Run(); err != nil {
				log.Fatal(err)
			}
		}
		tools.Println("Done.")
	},
}

var rootCmd = &cobra.Command{
	Use:        config.ProjectName,
	Aliases:    []string{},
	SuggestFor: []string{},
	Short:      "This is a tool to generate gozero service based on mysql",
	GroupID:    "",
	Long:       "This is a tool to generate gozero service based on mysql.\nThe goctl tool must be installed before use.\n\nGithub: https://github.com/licat233/genzero",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("%s requires at least one argument", cmd.CommandPath())
		}
		return nil
	},
	Version: config.CurrentVersion,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	config.C = config.New()

	rootCmd.PersistentFlags().StringVar(&config.C.DatabaseConfig.DSN, "dsn", "", "data source name (DSN) to use when connecting to the database")
	rootCmd.PersistentFlags().StringVar(&config.C.DatabaseConfig.Src, "src", "", "sql file to use when connecting to the database")
	rootCmd.PersistentFlags().StringSliceVar(&config.C.DatabaseConfig.Tables, "tables", []string{}, "need to generate tables, default is all tables，split multiple value by ','")
	rootCmd.PersistentFlags().StringSliceVar(&config.C.DatabaseConfig.IgnoreTables, "ignore_tables", []string{}, "ignore table string, default is none，split multiple value by ','")
	rootCmd.PersistentFlags().StringSliceVar(&config.C.DatabaseConfig.IgnoreColumns, "ignore_columns", []string{}, "ignore column string, default is none，split multiple value by ','")

	pbCmd.PersistentFlags().StringVar(&config.C.PbConfig.Package, "pkg", "", "proto package")
	pbCmd.PersistentFlags().StringVar(&config.C.PbConfig.GoPackage, "gopkg", "", "proto go package")
	pbCmd.PersistentFlags().BoolVar(&config.C.PbConfig.Multiple, "multiple", false, "proto multiple")
	pbCmd.PersistentFlags().StringVar(&config.C.PbConfig.Dir, "dir", "", "proto output directory")
	pbCmd.PersistentFlags().StringVar(&config.C.PbConfig.ServiceName, "service_name", "", "service name")
	pbCmd.PersistentFlags().StringSliceVar(&config.C.PbConfig.Tables, "tables", []string{}, "need to generate tables, default is all tables，split multiple value by ','")
	pbCmd.PersistentFlags().StringSliceVar(&config.C.PbConfig.IgnoreTables, "ignore_tables", []string{}, "ignore table string, default is none，split multiple value by ','")
	pbCmd.PersistentFlags().StringSliceVar(&config.C.PbConfig.IgnoreColumns, "ignore_columns", []string{}, "ignore column string, default is none，split multiple value by ','")

	apiCmd.PersistentFlags().StringVar(&config.C.ApiConfig.Style, "style", "", "api style")
	apiCmd.PersistentFlags().StringVar(&config.C.ApiConfig.Jwt, "jwt", "", "api jwt")
	apiCmd.PersistentFlags().StringVar(&config.C.ApiConfig.Middleware, "middleware", "", "api middleware")
	apiCmd.PersistentFlags().StringVar(&config.C.ApiConfig.Prefix, "prefix", "", "api prefix")
	apiCmd.PersistentFlags().BoolVar(&config.C.ApiConfig.Multiple, "multiple", false, "api multiple")
	apiCmd.PersistentFlags().StringVar(&config.C.ApiConfig.Dir, "dir", "", "api output directory")
	apiCmd.PersistentFlags().StringVar(&config.C.ApiConfig.ServiceName, "service_name", "", "service name")
	apiCmd.PersistentFlags().StringSliceVar(&config.C.ApiConfig.Tables, "tables", []string{}, "need to generate tables, default is all tables，split multiple value by ','")
	apiCmd.PersistentFlags().StringSliceVar(&config.C.ApiConfig.IgnoreTables, "ignore_tables", []string{}, "ignore table string, default is none，split multiple value by ','")
	apiCmd.PersistentFlags().StringSliceVar(&config.C.ApiConfig.IgnoreColumns, "ignore_columns", []string{}, "ignore column string, default is none，split multiple value by ','")

	modelCmd.PersistentFlags().StringVar(&config.C.ModelConfig.Dir, "dir", "", "model output directory")
	modelCmd.PersistentFlags().StringVar(&config.C.ApiConfig.ServiceName, "service_name", "", "service name")
	modelCmd.PersistentFlags().StringSliceVar(&config.C.ApiConfig.Tables, "tables", []string{}, "need to generate tables, default is all tables，split multiple value by ','")
	modelCmd.PersistentFlags().StringSliceVar(&config.C.ApiConfig.IgnoreTables, "ignore_tables", []string{}, "ignore table string, default is none，split multiple value by ','")
	modelCmd.PersistentFlags().StringSliceVar(&config.C.ApiConfig.IgnoreColumns, "ignore_columns", []string{}, "ignore column string, default is none，split multiple value by ','")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(upgradeCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(pbCmd)
	rootCmd.AddCommand(apiCmd)
	rootCmd.AddCommand(modelCmd)
	rootCmd.AddCommand(yamlCmd)
}

func Initialize() {
	if err := config.C.Validate(); err != nil {
		log.Fatal(err)
	}

	if err := global.InitSchema(); err != nil {
		log.Fatal(err)
	}
}

func Execute() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Fatalf("\nerror: %v", err)
	// 	}
	// }()
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
