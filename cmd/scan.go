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

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Run the scan",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

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

		client := proto.NewXmasPIClient(conn)

		resp, err := client.SetLed(context.Background(), &proto.SetLedRequest{
			Led:   0,
			Color: 0xFFFFFF,
		})

		if err != nil {
			log.Fatalf("Error sending message: %d\n", err)
		}

		fmt.Printf("Successfull request: %v\n", resp)

		/*
			webcam, _ := gocv.OpenVideoCapture("http://192.168.0.40:8080/video")
			window := gocv.NewWindow("Hello")

			defer webcam.Close()

			img := gocv.NewMat()

			for {
				webcam.Read(&img)
				window.IMShow(img)
				window.WaitKey(1)
			}*/
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
