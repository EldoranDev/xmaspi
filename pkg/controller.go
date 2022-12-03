package pkg

import (
	"github.com/EldoranDev/xmaspi/v2/internal/leds"
)

type Controller interface {
	Init() error
	Close()

	Render() error

	StartAnimation(animation Animation)
	StopAnimation()

	RenderStatic(static Static)

	SetLedRaw(led int, color uint32)
	SetLed(led int, color *Color)

	FillRaw(color uint32)
	Fill(color *Color)

	LedCount() int
	Leds() leds.Leds
}
