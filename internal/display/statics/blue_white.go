package statics

import (
	xmaspi "github.com/EldoranDev/xmaspi/v2/pkg"
	"math/rand"
)

var _ xmaspi.Static = (*BlueWhite)(nil)

type BlueWhite struct {
}

func (b *BlueWhite) Apply(ctrl xmaspi.Controller) {
	blue := xmaspi.Color{
		B: 244,
		R: 3,
		G: 252,
	}

	white := xmaspi.Color{
		B: 180,
		R: 180,
		G: 180,
	}

	for _, led := range ctrl.Leds() {
		if rand.Float32() > 0.5 {
			ctrl.SetLed(led.Index, &white)
		} else {
			ctrl.SetLed(led.Index, &blue)
		}
	}
}
