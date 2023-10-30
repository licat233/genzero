package cmd

import (
	"os"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/modules/logic"
	"github.com/licat233/genzero/tools"
	"github.com/spf13/cobra"
)

var logicCmd = &cobra.Command{
	Use:     "logic",
	GroupID: "modules",
	Short:   "Modify logic files, this feature has not been developed yet",
	Run: func(cmd *cobra.Command, args []string) {
		config.C.Logic.Status = true
		if err := Initialize(); err != nil {
			tools.Error("Failed to initialize: " + err.Error())
			os.Exit(1)
		}
		if err := logic.New().Run(); err != nil {
			tools.Warning(err.Error())
			os.Exit(1)
		}
		tools.Success("Done.")
	},
}

var apilogicCmd = &cobra.Command{
	Use:   "api",
	Short: "Modify api logic files",
	Run: func(cmd *cobra.Command, args []string) {
		config.C.Logic.Api.Status = true
		if err := Initialize(); err != nil {
			tools.Error("Failed to initialize: " + err.Error())
			os.Exit(1)
		}
		if err := logic.New().Run(); err != nil {
			tools.Warning(err.Error())
			os.Exit(1)
		}
		tools.Success("Done.")
	},
}

var rpclogicCmd = &cobra.Command{
	Use:   "rpc",
	Short: "Modify rpc logic files",
	Run: func(cmd *cobra.Command, args []string) {
		config.C.Logic.Rpc.Status = true
		if err := Initialize(); err != nil {
			tools.Error("Failed to initialize: " + err.Error())
			os.Exit(1)
		}
		if err := logic.New().Run(); err != nil {
			tools.Warning(err.Error())
			os.Exit(1)
		}
		tools.Success("Done.")
	},
}

func init() {
	apilogicCmd.PersistentFlags().BoolVar(&config.C.Logic.Api.UseRpc, "use_rpc", false, "use rpc for api")
	// apilogicCmd.PersistentFlags().BoolVar(&config.C.Logic.Api.RpcMultiple, "rpc_multiple", false, "is multiple rpc ?")
	apilogicCmd.PersistentFlags().StringVar(&config.C.Logic.Api.FileStyle, "file_style", config.LowerCamelCase, "file naming style: "+config.StyleList)
	apilogicCmd.PersistentFlags().StringVar(&config.C.Logic.Api.Dir, "dir", "", "api logic directory")
	apilogicCmd.PersistentFlags().StringSliceVar(&config.C.Logic.Api.Tables, "tables", []string{}, "need to generate tables, default is all tables，split multiple value by ','")
	apilogicCmd.PersistentFlags().StringSliceVar(&config.C.Logic.Api.IgnoreTables, "ignore_tables", []string{}, "ignore table string, default is none，split multiple value by ','")
	// apilogicCmd.PersistentFlags().StringSliceVar(&config.C.Logic.Api.IgnoreColumns, "ignore_columns", []string{}, "ignore column string, default is none，split multiple value by ','")

	logicCmd.AddCommand(apilogicCmd)

	rpclogicCmd.PersistentFlags().BoolVar(&config.C.Logic.Rpc.Multiple, "multiple", false, "is multiple ?")
	rpclogicCmd.PersistentFlags().StringVar(&config.C.Logic.Rpc.FileStyle, "file_style", config.LowerCamelCase, "file naming style: "+config.StyleList)
	rpclogicCmd.PersistentFlags().StringVar(&config.C.Logic.Rpc.Dir, "dir", "", "rpc logic directory")
	rpclogicCmd.PersistentFlags().StringSliceVar(&config.C.Logic.Rpc.Tables, "tables", []string{}, "need to generate tables, default is all tables，split multiple value by ','")
	// rpclogicCmd.PersistentFlags().StringSliceVar(&config.C.Logic.Rpc.IgnoreColumns, "ignore_columns", []string{}, "ignore column string, default is none，split multiple value by ','")

	logicCmd.AddCommand(rpclogicCmd)

	logicCmd.SetHelpTemplate(setColorizeHelp(logicCmd.HelpTemplate()))
	rootCmd.AddCommand(logicCmd)
}
