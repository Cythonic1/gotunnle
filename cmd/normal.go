/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Cythonic1/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var src string
var dest string


// filesCmd represents the files command
var normalMod = &cobra.Command{
	Use:   "normal --local-port",
	Short: "tunnle using socket",
	Long:  `Quickly scan a directory and find large files. . Use the flags below to target the output.`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.InitTunnling().SetSrc(src).SetDest(dest).RunTun()
	},
}

func init() {
	rootCmd.AddCommand(normalMod)

	normalMod.PersistentFlags().StringVarP(&src, "src", "s", "", "Src is from where the traffic gonna be generated ('localhost:2502')")
	viper.BindPFlag("src", normalMod.PersistentFlags().Lookup("src"))

	normalMod.PersistentFlags().StringVarP(&dest, "dest", "d", "", "Connect to where ('localhost:2502')")
	viper.BindPFlag("dest", normalMod.PersistentFlags().Lookup("src"))
}
