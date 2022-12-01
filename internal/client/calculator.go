package client

import (
	"github.com/EldoranDev/xmaspi/v2/internal/leds"
	"github.com/EldoranDev/xmaspi/v2/internal/math"
)

func ConvertScanToLeds(
	scan ScanResult,
	ledCount int,
) []leds.Led {
	list := make([]leds.Led, ledCount)

	// Go through all LEDs
	for led := 0; led < ledCount; led++ {
		// Y  should be up for everyone from every direction
		y := (scan[0][led].Y + scan[1][led].Y + scan[2][led].Y + scan[3][led].Y) / 4
		x := (scan[1][led].X + (1080 - scan[3][led].X)) / 2
		z := (scan[0][led].X + (1080 - scan[2][led].X)) / 2

		list[led] = leds.Led{
			Index: led,
			Pos: leds.Coordinate{
				X: math.Map(x, 0, 1080, 0, 100),
				Y: math.Map(1920-y, 0, 1920, 0, 180),
				Z: math.Map(z, 0, 1080, 0, 100),
			},
		}
	}

	return list
}
