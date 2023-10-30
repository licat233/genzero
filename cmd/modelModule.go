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
