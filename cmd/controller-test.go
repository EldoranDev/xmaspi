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
	Short: "Run a short controller test",
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
					controller.SetLedRaw(led, 0x00FF00)
				} else {
					controller.SetLedRaw(led, 0x000000)
				}

				_ = controller.Render()

				time.Sleep(time.Millisecond * 50)
			}
		}
	},
}

func init() {
	controllerCmd.AddCommand(testCmd)
}
