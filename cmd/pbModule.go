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

func init() {
	pbCmd.PersistentFlags().StringVar(&config.C.Pb.FileStyle, "file_style", config.LowerCamelCase, "proto file naming style: "+config.StyleList)
	pbCmd.PersistentFlags().StringVar(&config.C.Pb.Package, "pkg", "", "proto package")
	pbCmd.PersistentFlags().StringVar(&config.C.Pb.GoPackage, "gopkg", "", "proto go package")
	pbCmd.PersistentFlags().BoolVar(&config.C.Pb.Multiple, "multiple", false, "proto multiple")
	pbCmd.PersistentFlags().StringVar(&config.C.Pb.Dir, "dir", "", "proto output directory")
	pbCmd.PersistentFlags().StringVar(&config.C.Pb.ServiceName, "service_name", "", "service name, default value is database name")
	pbCmd.PersistentFlags().StringSliceVar(&config.C.Pb.Tables, "tables", []string{}, "need to generate tables, default is all tables，split multiple value by ','")
	pbCmd.PersistentFlags().StringSliceVar(&config.C.Pb.IgnoreTables, "ignore_tables", []string{}, "ignore table string, default is none，split multiple value by ','")
	pbCmd.PersistentFlags().StringSliceVar(&config.C.Pb.IgnoreColumns, "ignore_columns", []string{}, "ignore column string, default is none，split multiple value by ','")

	pbCmd.SetHelpTemplate(setColorizeHelp(pbCmd.HelpTemplate()))
	rootCmd.AddCommand(pbCmd)
}
