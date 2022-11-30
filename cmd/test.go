//go:build pi
// +build pi

package cmd

import (
	"github.com/EldoranDev/xmaspi/v2/internal/led"
	"github.com/spf13/cobra"
	"time"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run a short LED test",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		controller, err := led.NewController(
			led.WithLedCount(50),
			led.WithBrightness(128),
			led.WithDataPin(18),
		)

		if err != nil {
			return err
		}

		for i := 0; ; i++ {
			for led := 0; led < 50; led++ {
				if i%2 == 0 {
					_ = controller.SetLed(led, 0xFFFFFF, true)
				} else {
					_ = controller.SetLed(led, 0x000000, true)
				}
			}
			_ = controller.Fill(0x000000, true)

			time.Sleep(time.Millisecond * 100)
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
