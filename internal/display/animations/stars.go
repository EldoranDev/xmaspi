package animations

import (
	xmaspi "github.com/EldoranDev/xmaspi/v2/pkg"
	"math/rand"
	"time"
)

type Stars struct {
	Color xmaspi.Color

	factors []float32
}

func (s *Stars) Init(ctrl xmaspi.Controller) {
	s.factors = make([]float32, ctrl.LedCount())

	for i := 0; i < ctrl.LedCount(); i++ {
		s.factors[i] = rand.Float32()
	}
}

func (s *Stars) ApplyFrame(ctrl xmaspi.Controller) {
	for i := 0; i < ctrl.LedCount(); i++ {
		s.factors[i] = rand.Float32()
	}

	for i := 0; i < ctrl.LedCount(); i++ {
		ctrl.SetLedRaw(i, s.Color.DimmedInt(s.factors[i]))
	}
}

func (s *Stars) FrameDuration() time.Duration {
	return time.Second * 1
}
