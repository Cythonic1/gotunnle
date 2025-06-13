/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var localPort string
var bindPort string

var normalMod = &cobra.Command{
	Use:   "normal",
	Short: "Normal socket tunneling mode",
	Long:  "Normal TCP-based tunnel — supports server and client roles",
	// No Run here — we expect subcommands
}

// filesCmd represents the files command
// var normalMod = &cobra.Command{
// 	Use:   "normal --local-port",
// 	Short: "tunnle using socket",
// 	Long:  `Quickly scan a directory and find large files. . Use the flags below to target the output.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		pkg.InitTunnling().SetLocalPort(localPort).SetBindPort(bindPort).RunTun()
// 	},
// }

func init() {
	rootCmd.AddCommand(normalMod)

	normalMod.PersistentFlags().StringVarP(&localPort, "localPort", "l", "", "Listen port")
	viper.BindPFlag("localPort", normalMod.PersistentFlags().Lookup("localPort"))

	normalMod.PersistentFlags().StringVarP(&bindPort, "bindPort", "b", "", "Which port forward traffic from")
	viper.BindPFlag("bindPort", normalMod.PersistentFlags().Lookup("bindPort"))
}
