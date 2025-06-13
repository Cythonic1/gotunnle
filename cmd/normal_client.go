package cmd

import (
	"github.com/Cythonic1/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var target string
var serverAddr string

var normalClient = &cobra.Command{
	Use:   "client",
	Short: "Start the tunnel client (connects to server)",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.ClientInternal(target, serverAddr)
	},
}

func init() {
	normalClient.Flags().StringVarP(&target, "target", "t", "", "target service ip:port (localhost:8080)")
	normalClient.Flags().StringVarP(&serverAddr, "serveraddr", "s", "", "server address (localhost:8080)")
	viper.BindPFlag("target", normalClient.Flags().Lookup("target"))
	viper.BindPFlag("serveraddr", normalClient.Flags().Lookup("serveraddr"))

	normalMod.AddCommand(normalClient)
}
