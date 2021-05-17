package cmd

import (
	"github.com/joeyscat/oook/internal/staticserver"
	"github.com/spf13/cobra"
)

var StaticServerPort uint
var StaticServerDirectory string

var staticServerCmd = &cobra.Command{
	Use:   "static-server",
	Short: "Running a Static Server",
	Long:  `Running a Static Server on a special port`,
	RunE: func(cmd *cobra.Command, args []string) error {
		server := staticserver.NewStaticServer(StaticServerDirectory, StaticServerPort)

		return server.Run()
	},
}
