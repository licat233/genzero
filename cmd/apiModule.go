package cmd

import (
	"os"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/modules/api"
	"github.com/licat233/genzero/tools"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:     "api",
	GroupID: "modules",
	Short:   "Generate .api files",
	Run: func(cmd *cobra.Command, args []string) {
		config.C.Api.Status = true
		if err := Initialize(); err != nil {
			tools.Error("Failed to initialize: " + err.Error())
			os.Exit(1)
		}
		if err := api.New().Run(); err != nil {
			tools.Warning(err.Error())
			os.Exit(1)
		}
		tools.Success("Done.")
	},
}

func init() {
	apiCmd.PersistentFlags().StringVar(&config.C.Api.JsonStyle, "json_style", config.SnakeCase, "JSON field naming style: "+config.StyleList)
	apiCmd.PersistentFlags().StringVar(&config.C.Api.Jwt, "jwt", "", "api jwt")
	apiCmd.PersistentFlags().StringSliceVar(&config.C.Api.Middleware, "middleware", []string{}, "api middleware")
	apiCmd.PersistentFlags().StringVar(&config.C.Api.Prefix, "prefix", "", "api prefix")
	apiCmd.PersistentFlags().BoolVar(&config.C.Api.Multiple, "multiple", false, "api multiple")
	apiCmd.PersistentFlags().StringVar(&config.C.Api.Dir, "dir", "", "api output directory")
	apiCmd.PersistentFlags().StringVar(&config.C.Api.ServiceName, "service_name", "", "service name, default value is database name")
	apiCmd.PersistentFlags().StringSliceVar(&config.C.Api.Tables, "tables", []string{}, "need to generate tables, default is all tables，split multiple value by ','")
	apiCmd.PersistentFlags().StringSliceVar(&config.C.Api.IgnoreTables, "ignore_tables", []string{}, "ignore table string, default is none，split multiple value by ','")
	apiCmd.PersistentFlags().StringSliceVar(&config.C.Api.IgnoreColumns, "ignore_columns", []string{}, "ignore column string, default is none，split multiple value by ','")

	apiCmd.SetHelpTemplate(setColorizeHelp(apiCmd.HelpTemplate()))
	rootCmd.AddCommand(apiCmd)
}
