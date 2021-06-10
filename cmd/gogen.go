package cmd

import (
	"github.com/joeyscat/oook/internal/gen"
	"github.com/spf13/cobra"
)

var ModuleName string

var genGoCmd = &cobra.Command{
	Use:   "gengo",
	Short: "Generate Golang Project",
	Long:  `Generate Golang Project`,
	RunE: func(cmd *cobra.Command, args []string) error {
		g := gen.NewGoGenerator(ModuleName)

		return g.Generate()
	},
}
