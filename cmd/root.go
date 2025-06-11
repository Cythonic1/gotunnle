package cmd

import (
	"os"

	"github.com/spf13/cobra"
	_ "github.com/spf13/viper"
)

// TODO: Implement the root command
var rootCmd = &cobra.Command{
	Use:   "gotunnl",
	Short: "Tunnling tool to tunnl your traffic DUH",
	Long:  `This tool is design to tunnl your traffic with multiple ways`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
