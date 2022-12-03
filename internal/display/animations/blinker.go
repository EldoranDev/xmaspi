package animations

import (
	xmaspi "github.com/EldoranDev/xmaspi/v2/pkg"
	"time"
)

var _ xmaspi.Animation = (*Blinker)(nil)

type Blinker struct {
	active bool
}

func (b *Blinker) FrameDuration() time.Duration {
	return time.Second * 1
}

func (b *Blinker) ApplyFrame(ctrl xmaspi.Controller) {
	b.active = !b.active

	if b.active {
		ctrl.FillRaw(0x000000)
	} else {
		ctrl.FillRaw(0xFFFFFF)
	}
}
