package rendering

import (
	"github.com/EldoranDev/xmaspi/v3/internal/led"
	"math/rand"
	"time"
)

const BlueWhiteStatic = "blue-white-static"

var _ Renderer = (*blueWhiteStatic)(nil)

type blueWhiteStatic struct {
	white led.Color
	blue  led.Color
}

func (b *blueWhiteStatic) Init(ctrl led.Controller) {
	b.blue = led.Color{
		B: 244,
		R: 3,
		G: 252,
	}

	b.white = led.Color{
		B: 180,
		R: 180,
		G: 180,
	}
}

func (*blueWhiteStatic) FrameDuration() time.Duration {
	return time.Hour * 99
}

func (b *blueWhiteStatic) ApplyFrame(ctrl led.Controller, color *led.Color) {
	for i := range ctrl.GetLedCount() {
		if rand.Float32() > 0.5 {
			ctrl.SetLed(int(i), &b.white)
		} else {
			ctrl.SetLed(int(i), &b.blue)
		}
	}
}

func (*blueWhiteStatic) GetIdentifier() string {
	return BlueWhiteStatic
}
