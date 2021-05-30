package cmd

import (
	"github.com/joeyscat/oook/internal/proxy"
	"github.com/spf13/cobra"
)

var ProxyPort uint
var ProxyVerbose bool

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Running a Proxy Server",
	Long:  `Running a Proxy Server on a special port`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p := proxy.NewProxyServer(ProxyPort, ProxyVerbose)

		return p.Serve()
	},
}
