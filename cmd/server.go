//go:build pi
// +build pi

package cmd

import (
	"fmt"
	"github.com/EldoranDev/xmaspi/v2/internal/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen(
			"tcp",
			fmt.Sprintf("%s:%d", viper.GetString("grpc.host"), viper.GetInt("grpc.port")),
		)

		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}

		var opts []grpc.ServerOption

		server := grpc.NewServer(opts...)

		proto.RegisterXmasPIServer(server, proto.NewServer())

		_ = server.Serve(lis)

	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

}
