package cmd

import (
	"github.com/Cythonic1/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var normalServer = &cobra.Command{
	Use:   "server",
	Short: "Start the tunnel server",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.InitTunnling().SetLocalPort(localPort).SetBindPort(bindPort).RunTun()
	},
}

func init() {
	normalServer.Flags().StringVarP(&localPort, "localPort", "l", "", "Port to listen for client connections")
	normalServer.Flags().StringVarP(&bindPort, "bindPort", "b", "", "Port to bind for incoming forwarded requests")
	viper.BindPFlag("localPort", normalServer.Flags().Lookup("localPort"))
	viper.BindPFlag("bindPort", normalServer.Flags().Lookup("bindPort"))

	normalMod.AddCommand(normalServer)
}
