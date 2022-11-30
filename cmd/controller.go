//go:build pi
// +build pi

package cmd

import (
	"github.com/EldoranDev/xmaspi/v2/internal/led"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// controllerCmd represents the controller command
var controllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "A brief description of your command",
}

func init() {
	rootCmd.AddCommand(controllerCmd)

	testCmd.PersistentFlags().Int("leds", 50, "")
	testCmd.PersistentFlags().Int("brightness", 128, "")

	viper.BindPFlag("led.count", testCmd.PersistentFlags().Lookup("leds"))
	viper.BindPFlag("led.brightness", testCmd.PersistentFlags().Lookup("brightness"))
}

func getController() (led.Controller, error) {
	return led.NewController(
		led.WithLedCount(viper.GetInt("led.count")),
		led.WithBrightness(viper.GetInt("led.brightness")),
		led.WithDataPin(18),
	)
}
