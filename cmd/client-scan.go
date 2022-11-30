//go:build client
// +build client

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/EldoranDev/xmaspi/v2/internal/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gocv.io/x/gocv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"image"
	"image/color"
	"log"
	"os"
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

		_, _ = client.SetStatic(context.Background(), &proto.SetStaticRequest{
			Color: 0x000000,
		})

		info, err := client.GetControllerInfo(context.Background(), &proto.ControllerInfoRequest{})

		if err != nil {
			return err
		}

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

		blurKernel := image.Point{X: 41, Y: 41}

		detectColor := color.RGBA{R: 255}
		posColor := color.RGBA{G: 255}

		var pos *image.Point

		// Parse Raw Data

		file, err := os.ReadFile("./scan.json")

		scan := make([][]image.Point, 3)

		for s := range scan {
			scan[s] = make([]image.Point, info.LedCount)
		}

		if err == nil {
			json.Unmarshal(file, &scan)
		}

		side := 0
		led := 0

		client.SetLed(context.Background(), &proto.SetLedRequest{Led: int32(led), Color: 0xFFFFFF})

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

			gocv.Circle(&img, loc, 20, detectColor, 5)
			gocv.Circle(&img, *pos, 20, posColor, 4)

			window.IMShow(img)

			key := window.WaitKey(1)

			if key == -1 {
				continue
			}

			if key == 32 {
				scan[side][led] = image.Point{
					X: pos.X,
					Y: pos.Y,
				}

				led++

				if int64(led) == info.LedCount {
					led = 0
					side++
				}

				client.SetStatic(context.Background(), &proto.SetStaticRequest{Color: 0x0})
				client.SetLed(context.Background(), &proto.SetLedRequest{
					Led:   int32(led),
					Color: 0xFFFFFF,
				})

				if scan[side][led].X != 0 {
					pos.X = scan[side][led].X
					pos.Y = scan[side][led].Y
				}

				continue
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
			case 115:
				out, _ := json.Marshal(scan)
				os.WriteFile("./scan.json", out, 0644)
				break
			default:
				print(key)
			}
		}
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
