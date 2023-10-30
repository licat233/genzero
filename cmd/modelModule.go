package cmd

import (
	"os"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/modules/model"
	"github.com/licat233/genzero/tools"
	"github.com/spf13/cobra"
)

var modelCmd = &cobra.Command{
	Use:     "model",
	GroupID: "modules",
	Short:   "Generate model code",
	Run: func(cmd *cobra.Command, args []string) {
		config.C.Model.Status = true
		if err := Initialize(); err != nil {
			tools.Error("Failed to initialize: " + err.Error())
			os.Exit(1)
		}
		if err := model.New().Run(); err != nil {
			tools.Warning(err.Error())
			os.Exit(1)
		}
		tools.Success("Done.")
	},
}

func init() {
	modelCmd.PersistentFlags().StringVar(&config.C.Model.Dir, "dir", "", "model output directory")
	modelCmd.PersistentFlags().StringVar(&config.C.Api.ServiceName, "service_name", "", "service name, default value is database name")
	modelCmd.PersistentFlags().StringSliceVar(&config.C.Api.Tables, "tables", []string{}, "need to generate tables, default is all tables，split multiple value by ','")
	modelCmd.PersistentFlags().StringSliceVar(&config.C.Api.IgnoreTables, "ignore_tables", []string{}, "ignore table string, default is none，split multiple value by ','")
	modelCmd.PersistentFlags().StringSliceVar(&config.C.Api.IgnoreColumns, "ignore_columns", []string{}, "ignore column string, default is none，split multiple value by ','")

	modelCmd.SetHelpTemplate(setColorizeHelp(modelCmd.HelpTemplate()))
	rootCmd.AddCommand(modelCmd)
}
