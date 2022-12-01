//go:build client
// +build client

package cmd

import (
	"github.com/EldoranDev/xmaspi/v2/internal/client"
	"github.com/EldoranDev/xmaspi/v2/internal/leds"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// calcCmd represents the calc command
var calcCmd = &cobra.Command{
	Use:   "calc",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		scan, err := client.LoadScan(
			scanFile,
			viper.GetInt("led.count"),
		)

		if err != nil {
			return err
		}

		ledList := client.ConvertScanToLeds(scan, viper.GetInt("led.count"))

		return leds.SaveLeds(coordinateFile, ledList)
	},
}

func init() {
	clientCmd.AddCommand(calcCmd)

	calcCmd.PersistentFlags().String("out-file", "leds.json", "")

	viper.BindPFlag("scan.coordinates", calcCmd.PersistentFlags().Lookup("out-file"))
}
