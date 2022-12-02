package animations

import (
	"github.com/EldoranDev/xmaspi/v2/internal/controller"
	"github.com/EldoranDev/xmaspi/v2/internal/leds"
	"math/rand"
	"time"
)

type Stars struct {
	Color leds.Color

	factors []float32
}

func (s *Stars) Init(ctrl controller.Controller) {
	s.factors = make([]float32, ctrl.LedCount())

	for i := 0; i < ctrl.LedCount(); i++ {
		s.factors[i] = rand.Float32()
	}
}

func (s *Stars) ApplyFrame(ctrl controller.Controller) {
	for i := 0; i < ctrl.LedCount(); i++ {
		s.factors[i] = rand.Float32()
	}

	for i := 0; i < ctrl.LedCount(); i++ {
		ctrl.SetLed(i, s.Color.DimmedInt(s.factors[i]))
	}
}

func (s *Stars) FrameDuration() time.Duration {
	return time.Second * 1
}
