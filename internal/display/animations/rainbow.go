package animations

import (
	xmaspi "github.com/EldoranDev/xmaspi/v2/pkg"
	"time"
)

type Rainbow struct {
	marker int

	speed int
	color int

	colors []xmaspi.Color
}

func (s *Rainbow) Init(ctrl xmaspi.Controller) {
	s.marker = 50
	s.speed = 5
	s.color = 0

	s.colors = []xmaspi.Color{
		{R: 255, B: 0, G: 0},
		{R: 200, B: 0, G: 100},
		{R: 150, B: 0, G: 150},
		{R: 0, B: 0, G: 255},
		{R: 0, B: 200, G: 0},
		{R: 0, B: 200, G: 100},
		{R: 100, B: 200, G: 0},
	}
}

func (s *Rainbow) ApplyFrame(ctrl xmaspi.Controller) {
	s.marker += s.speed

	if s.marker > 110 {
		s.marker = 0
		s.color = (s.color + 1) % len(s.colors)
	}

	for _, led := range ctrl.Leds() {
		if led.Pos.X < s.marker {
			ctrl.SetLed(led.Index, &s.colors[s.color])
		}
	}
}

func (s *Rainbow) FrameDuration() time.Duration {
	return time.Millisecond * 100
}
