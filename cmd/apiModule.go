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
