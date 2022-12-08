//go:build pi
// +build pi

package cmd

import (
	"fmt"
	"github.com/EldoranDev/xmaspi/v2/internal/display/statics"
	"github.com/EldoranDev/xmaspi/v2/internal/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

// serverCmd represents the api command
var controllerServerCmd = &cobra.Command{
	Use:   "api",
	Short: "start the xmaspi api",
	RunE: func(cmd *cobra.Command, args []string) error {
		lis, err := net.Listen(
			"tcp",
			fmt.Sprintf("%s:%d", viper.GetString("grpc.host"), viper.GetInt("grpc.port")),
		)

		if err != nil {
			log.Fatalf("Failed to start api: %v", err)
		}

		var opts []grpc.ServerOption

		server := grpc.NewServer(opts...)
		controller, err := getController()

		if err != nil {
			return err
		}

		proto.RegisterXmasPIServer(server, proto.NewServer(controller))

		controller.Init()
		defer controller.Close()

		static, err := statics.GetStatic("BlueWhite")

		controller.RenderStatic(static)

		_ = server.Serve(lis)

		return nil
	},
}

func init() {
	controllerCmd.AddCommand(controllerServerCmd)
}
