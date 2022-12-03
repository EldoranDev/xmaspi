package animations

import (
	xmaspi "github.com/EldoranDev/xmaspi/v2/pkg"
	"time"
)

type LeftRight struct {
	Left  xmaspi.Color
	Right xmaspi.Color

	buffer    int
	speed     int
	direction int
	position  int
}

func (s *LeftRight) Init(ctrl xmaspi.Controller) {
	s.position = 50
	s.direction = -1
	s.speed = 5
	s.buffer = 10
}

func (s *LeftRight) ApplyFrame(ctrl xmaspi.Controller) {
	s.position += s.direction * s.speed

	if s.position > 100-s.buffer || s.position < s.buffer {
		s.direction *= -1
	}

	for _, led := range ctrl.Leds() {
		if led.Pos.X > s.position {
			ctrl.SetLed(led.Index, &s.Right)
		} else {
			ctrl.SetLed(led.Index, &s.Left)
		}
	}
}

func (s *LeftRight) FrameDuration() time.Duration {
	return time.Millisecond * 100
}
