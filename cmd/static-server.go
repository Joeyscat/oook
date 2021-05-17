package cmd

import (
	"fmt"

	"github.com/joeyscat/oook/internal/staticserver"
	"github.com/spf13/cobra"
)

var StaticServerPort uint
var StaticServerPath string

var staticServerCmd = &cobra.Command{
	Use:   "static-server",
	Short: "Running a Static Server",
	Long:  `Running a Static Server on a special port`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Server Running on http://127.0.0.1:%d\n", StaticServerPort)
		fmt.Printf("path: %s port: %d", StaticServerPath, StaticServerPort)

		server := staticserver.NewStaticServer(StaticServerPath, StaticServerPort)

		return server.Run()
	},
}
