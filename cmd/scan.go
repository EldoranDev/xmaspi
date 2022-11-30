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

var cameraDevice string

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Run the scan",
	Long:  ``,
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

		client := proto.NewXmasPIClient(conn)

		resp, err := client.SetStatic(context.Background(), &proto.SetStaticRequest{
			Color: 0x00FF00,
		})

		if err != nil {
			log.Fatalf("Error sending message: %v\n", err)
		}

		fmt.Printf("Successfull request: %v\n", resp)

		return nil
		/*

			var device interface{}

			switch cameraDevice {
			case "remote":
				device = viper.GetString("cam.remote")
				break
			case "local":
				device = viper.GetInt("cam.local")
				break
			}
			webcam, err := gocv.OpenVideoCapture(device)

			if err != nil {
				return err
			}

			window := gocv.NewWindow("Hello")

			defer webcam.Close()

			img := gocv.NewMat()
			gray := gocv.NewMat()

			blurKernel := image.Point{
				X: 41,
				Y: 41,
			}

			detectColor := color.RGBA{
				R: 255,
			}

			posColor := color.RGBA{
				G: 255,
			}

			var pos *image.Point

			for {
				webcam.Read(&img)

				gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

				gocv.GaussianBlur(gray, &gray, blurKernel, 0, 0, gocv.BorderDefault)

				_, _, _, loc := gocv.MinMaxLoc(gray)

				if pos == nil {
					pos = &image.Point{
						X: loc.X,
						Y: loc.Y,
					}
				}

				gocv.Circle(&img, loc, 5, detectColor, 1)
				gocv.Circle(&img, *pos, 5, posColor, 1)

				window.IMShow(img)

				key := window.WaitKey(1)

				if key == -1 {
					continue
				}

				if key == 32 {
					continue
					// go next
				}

				switch key {
				case 0:
					pos.Y -= 10
					break
				case 3:
					pos.X += 10
					break
				case 1:
					pos.Y += 10
					break
				case 2:
					pos.X -= 10
					break
				case 114:
					pos = nil
					break
				default:
					print(key)
				}
			}
		*/
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.PersistentFlags().String("remote-camera", "", "")
	scanCmd.PersistentFlags().Int("device", 0, "")

	scanCmd.PersistentFlags().StringVarP(&cameraDevice, "camera", "c", "local", "")

	viper.BindPFlag("cam.remote", scanCmd.PersistentFlags().Lookup("remote-camera"))
	viper.BindPFlag("cam.device", scanCmd.PersistentFlags().Lookup("device"))
}
