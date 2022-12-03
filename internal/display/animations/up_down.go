package animations

import (
	xmaspi "github.com/EldoranDev/xmaspi/v2/pkg"
	"time"
)

type UpDown struct {
	Top    xmaspi.Color
	Bottom xmaspi.Color

	buffer    int
	speed     int
	direction int
	position  int
}

func (s *UpDown) Init(ctrl xmaspi.Controller) {
	s.position = 90
	s.direction = -1
	s.speed = 5
	s.buffer = 20
}

func (s *UpDown) ApplyFrame(ctrl xmaspi.Controller) {
	s.position += s.direction * s.speed

	if s.position > 180-s.buffer || s.position < s.buffer {
		s.direction *= -1
	}

	for _, led := range ctrl.Leds() {
		if led.Pos.Y > s.position {
			ctrl.SetLed(led.Index, &s.Top)
		} else {
			ctrl.SetLed(led.Index, &s.Bottom)
		}
	}
}

func (s *UpDown) FrameDuration() time.Duration {
	return time.Millisecond * 100
}
