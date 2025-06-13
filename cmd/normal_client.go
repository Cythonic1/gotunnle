package cmd

import (
	"github.com/Cythonic1/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var target string
var port string

var normalClient = &cobra.Command{
	Use:   "client",
	Short: "Start the tunnel client (connects to server)",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.ClientInternal(target, port)
	},
}

func init() {
	normalClient.Flags().StringVarP(&target, "target", "t", "", "Target server IP")
	normalClient.Flags().StringVarP(&port, "port", "p", "", "Target server port")
	viper.BindPFlag("target", normalClient.Flags().Lookup("target"))
	viper.BindPFlag("port", normalClient.Flags().Lookup("port"))

	normalMod.AddCommand(normalClient)
}
