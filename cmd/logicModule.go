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
