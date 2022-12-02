//go:build pi
// +build pi

package cmd

import (
	"github.com/EldoranDev/xmaspi/v2/internal/animations"
	"github.com/EldoranDev/xmaspi/v2/internal/leds"
	"github.com/spf13/cobra"
	"time"
)

// testCmd represents the test command
var controllerTestAnimationCmd = &cobra.Command{
	Use:   "animation",
	Short: "Run a short controller test",
	RunE: func(cmd *cobra.Command, args []string) error {
		controller, err := getController()

		if err != nil {
			return err
		}

		controller.Init()
		defer controller.Close()

		controller.StartAnimation(&animations.Stars{
			Color: leds.Color{
				B: 255,
				R: 3,
				G: 200,
			},
		})

		time.Sleep(30 * time.Second)

		return nil
	},
}

func init() {
	testCmd.AddCommand(controllerTestAnimationCmd)
}
