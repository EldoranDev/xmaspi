//go:build client
// +build client

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var scanFile string

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
}

func init() {
	rootCmd.AddCommand(clientCmd)

	clientCmd.PersistentFlags().StringVar(&scanFile, "scan-file", "scan.json", "")

	viper.BindPFlag("scanner.result", clientCmd.PersistentFlags().Lookup("scan-file"))
}
