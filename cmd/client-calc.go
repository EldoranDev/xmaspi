//go:build client
// +build client

package cmd

import (
	"github.com/EldoranDev/xmaspi/v2/internal/client"
	"github.com/EldoranDev/xmaspi/v2/internal/leds"
	"github.com/spf13/viper"
	"image"

	"github.com/spf13/cobra"
)

var scanWidth int
var scanHeight int

// calcCmd represents the calc command
var calcCmd = &cobra.Command{
	Use:   "calc",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		scan, err := client.LoadScan(
			scanFile,
			viper.GetInt("leds.count"),
		)

		if err != nil {
			return err
		}

		ledList := client.ConvertScanToLeds(
			scan,
			viper.GetInt("leds.count"),
			image.Point{X: scanWidth, Y: scanHeight},
			image.Point{X: treeWidth, Y: treeHeight},
		)

		return leds.SaveLeds(ledsFile, ledList)
	},
}

func init() {
	clientCmd.AddCommand(calcCmd)

	calcCmd.PersistentFlags().IntVar(&scanWidth, "scan-width", 1080, "")
	calcCmd.PersistentFlags().IntVar(&scanHeight, "scan-height", 1920, "")

	viper.BindPFlag("scanner.camera.resolution.width", rootCmd.PersistentFlags().Lookup("scan-width"))
	viper.BindPFlag("scanner.camera.resolution.height", rootCmd.PersistentFlags().Lookup("scan-height"))

}
