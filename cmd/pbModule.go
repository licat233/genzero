package cmd

import (
	"os"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/modules/pb"
	"github.com/licat233/genzero/tools"
	"github.com/spf13/cobra"
)

var pbCmd = &cobra.Command{
	Use:     "pb",
	GroupID: "modules",
	Aliases: []string{"proto"},
	Short:   "Generate .proto files",
	Run: func(cmd *cobra.Command, args []string) {
		config.C.Pb.Status = true
		if err := Initialize(); err != nil {
			tools.Error("Failed to initialize: " + err.Error())
			os.Exit(1)
		}
		if err := pb.New().Run(); err != nil {
			tools.Warning(err.Error())
			os.Exit(1)
		}
		tools.Success("Done.")
	},
}
