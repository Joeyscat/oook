package cmd

import (
	"fmt"
	"github.com/joeyscat/oook/internal/version"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "oook",
		Short: "OOOK is a very useful tool written in golang",
		Long:  `something more...`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	staticServerCmd.PersistentFlags().StringVarP(&StaticServerDirectory, "directory", "d", ".", "Directory for Static Server")
	staticServerCmd.PersistentFlags().UintVarP(&StaticServerPort, "port", "p", 8000, "Port for Static Server")

	proxyCmd.PersistentFlags().UintVarP(&ProxyPort, "port", "p", 1080, "Port for Proxy Server")

	genGoCmd.PersistentFlags().StringVarP(&ModuleName, "module", "m", "", "Generate Golang Project")
	err := genGoCmd.MarkPersistentFlagRequired("module")
	if err != nil {
		panic(err)
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(staticServerCmd)
	rootCmd.AddCommand(proxyCmd)
	rootCmd.AddCommand(genGoCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of OOOK",
	Long:  `Print the version number of OOOK`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.FullVersion())
	},
}
