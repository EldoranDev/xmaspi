package animations

import (
	"github.com/EldoranDev/xmaspi/v2/internal/controller"
	"github.com/EldoranDev/xmaspi/v2/internal/leds"
	"time"
)

type UpDown struct {
	Top    leds.Color
	Bottom leds.Color

	buffer    int
	speed     int
	direction int
	position  int
}

func (s *UpDown) Init(ctrl controller.Controller) {
	s.position = 90
	s.direction = -1
	s.speed = 5
	s.buffer = 20
}

func (s *UpDown) ApplyFrame(ctrl controller.Controller) {
	s.position += s.direction * s.speed

	if s.position > 180-s.buffer || s.position < s.buffer {
		s.direction *= -1
	}

	for _, led := range ctrl.Leds() {
		if led.Pos.Y > s.position {
			ctrl.SetLed(led.Index, 0xFF0000)
		} else {
			ctrl.SetLed(led.Index, 0x0000FF)
		}
	}
}

func (s *UpDown) FrameDuration() time.Duration {
	return time.Millisecond * 100
}
