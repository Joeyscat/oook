package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "oook",
		Short: "OOOK is a very useful toolbox written in golang",
		Long:  `something more...`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")

	staticServerCmd.PersistentFlags().StringVarP(&StaticServerPath, "path", "", ".", "Path for Static Server")
	staticServerCmd.PersistentFlags().UintVarP(&StaticServerPort, "port", "p", 8000, "Port for Static Server")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(staticServerCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of OOOK",
	Long:  `Print the version number of OOOK`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("OOOK Toolbox v0.9 -- HEAD")
	},
}
