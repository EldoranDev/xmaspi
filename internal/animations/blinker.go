package animations

import (
	"github.com/EldoranDev/xmaspi/v2/internal/controller"
	"time"
)

type Blinker struct {
	active bool
}

func (b *Blinker) FrameDuration() time.Duration {
	return time.Second * 1
}

func (b *Blinker) ApplyFrame(ctrl controller.Controller) {
	b.active = !b.active

	if b.active {
		ctrl.Fill(0x000000)
	} else {
		ctrl.Fill(0xFFFFFF)
	}
}
