//go:build client
// +build client

package cmd

import (
	"context"
	"fmt"
	"github.com/EldoranDev/xmaspi/v2/internal/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// clientTestCmd represents the clientTest command
var clientTestCmd = &cobra.Command{
	Use:   "test",
	Short: "Run a short test of the client",
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}

		conn, err := grpc.Dial(
			fmt.Sprintf("%s:%d", viper.GetString("grpc.host"), viper.GetInt("grpc.port")),
			opts...,
		)

		if err != nil {
			log.Fatalf("Error connecting: %d", err)
		}

		defer conn.Close()

		pi := proto.NewXmasPIClient(conn)

		pi.SetAnimation(context.Background(), &proto.SetAnimationRequest{Name: "Test"})

		return nil
	},
}

func init() {
	clientCmd.AddCommand(clientTestCmd)
}
