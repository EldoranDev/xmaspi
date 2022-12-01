//go:build scanner
// +build scanner

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	client2 "github.com/EldoranDev/xmaspi/v2/internal/client"
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
	Short: "Run the scanner",
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

		_, _ = pi.SetStatic(context.Background(), &proto.SetStaticRequest{
			Color: 0x000000,
		})

		info, err := pi.GetControllerInfo(context.Background(), &proto.ControllerInfoRequest{})

		if err != nil {
			return err
		}

		var device interface{}

		switch cameraDevice {
		case "remote":
			device = viper.GetString("scanner.camera.remote")
			break
		case "local":
			device = viper.GetInt("scanner.camera.local")
			break
		}
		webcam, err := gocv.OpenVideoCapture(device)

		if err != nil {
			return err
		}

		window := gocv.NewWindow("Scanner")

		defer webcam.Close()

		img := gocv.NewMat()
		gray := gocv.NewMat()

		blurKernel := image.Point{X: 41, Y: 41}

		detectColor := color.RGBA{R: 255}
		posColor := color.RGBA{G: 255}

		var pos *image.Point

		// Import exisint scans
		scan, err := client2.LoadScan(
			scanFile,
			ledCount,
		)

		if err != nil {
			return err
		}

		side := 0
		led := 0

		pi.SetLed(context.Background(), &proto.SetLedRequest{Led: int64(led), Color: 0xFFFFFF})

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

					log.Println("Tree rotation")
				}

				if side == 4 {
					break
				}

				pi.SetStatic(context.Background(), &proto.SetStaticRequest{Color: 0x0})
				pi.SetLed(context.Background(), &proto.SetLedRequest{
					Led:   int64(led),
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
				os.WriteFile(scanFile, out, 0644)
				break
			default:
				print(key)
			}
		}

		out, _ := json.Marshal(scan)
		os.WriteFile(scanFile, out, 0644)

		return nil
	},
}

func init() {
	clientCmd.AddCommand(scanCmd)

	scanCmd.PersistentFlags().String("camera", "", "")
	scanCmd.PersistentFlags().Int("device", 0, "")

	scanCmd.PersistentFlags().StringVarP(&cameraDevice, "camera", "c", "local", "")

	viper.BindPFlag("scanner.camera.remote", scanCmd.PersistentFlags().Lookup("camera"))
	viper.BindPFlag("scanner.camera.device", scanCmd.PersistentFlags().Lookup("device"))
}
