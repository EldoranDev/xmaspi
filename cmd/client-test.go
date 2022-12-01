//go:build client
// +build client

package cmd

import (
	"context"
	"fmt"
	"github.com/EldoranDev/xmaspi/v2/internal/leds"
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
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		ledList, err := leds.LoadLeds("leds.json")

		if err != nil {
			return err
		}

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

		_, _ = pi.SetStatic(context.Background(), &proto.SetStaticRequest{
			Color: 0x00FF00,
		})

		for _, led := range ledList {
			if led.Pos.X > 50 {
				pi.SetLed(
					context.Background(),
					&proto.SetLedRequest{Led: int64(led.Index), Color: 0x0000FF, Render: false},
				)
			}
		}

		pi.Render(context.Background(), &proto.RenderRequest{})

		for _, led := range ledList {
			if led.Pos.X <= 50 {
				pi.SetLed(
					context.Background(),
					&proto.SetLedRequest{Led: int64(led.Index), Color: 0xFF0000, Render: false},
				)
			}
		}

		pi.Render(context.Background(), &proto.RenderRequest{})

		return nil
	},
}

func init() {
	clientCmd.AddCommand(clientTestCmd)
}
