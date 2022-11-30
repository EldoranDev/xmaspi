//go:build pi
// +build pi

package cmd

import (
	"github.com/spf13/cobra"
	"time"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run a short LED test",
	RunE: func(cmd *cobra.Command, args []string) error {
		controller, err := getController()

		if err != nil {
			return err
		}

		controller.Init()
		defer controller.Close()

		for i := 0; ; i++ {
			for led := 0; led < controller.LedCount(); led++ {
				if i%2 == 0 {
					_ = controller.SetLed(led, 0xFFFFFF, true)
				} else {
					_ = controller.SetLed(led, 0x000000, true)
				}

				time.Sleep(time.Millisecond * 50)
			}
		}
	},
}

func init() {
	controllerCmd.AddCommand(testCmd)
}
