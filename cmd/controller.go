//go:build pi
// +build pi

package cmd

import (
	"encoding/json"
	"github.com/EldoranDev/xmaspi/v2/internal/controller"
	"github.com/EldoranDev/xmaspi/v2/internal/leds"
	xmaspi "github.com/EldoranDev/xmaspi/v2/pkg"
	"github.com/spf13/viper"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// controllerCmd represents the controller command
var controllerCmd = &cobra.Command{
	Use: "controller",
}

func init() {
	rootCmd.AddCommand(controllerCmd)

	testCmd.PersistentFlags().Int("brightness", 128, "")

	viper.BindPFlag("leds.brightness", testCmd.PersistentFlags().Lookup("brightness"))
}

func getController() (xmaspi.Controller, error) {
	log.Printf("Using leds from %s", ledsFile)
	file, err := os.ReadFile(ledsFile)

	if err != nil {
		return nil, err
	}

	var list leds.Leds

	err = json.Unmarshal(file, &list)
	if err != nil {
		return nil, err
	}

	log.Printf("Starting with %d connected leds", len(list))

	return controller.NewController(
		list,
		controller.WithBrightness(viper.GetInt("leds.brightness")),
		controller.WithDataPin(18),
	)
}
