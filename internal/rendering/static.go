package rendering

import (
	"github.com/EldoranDev/xmaspi/v3/internal/led"
	"time"
)

type static struct {
}

func (*static) FrameDuration() time.Duration {
	return time.Hour * 99
}

func (*static) ApplyFrame(ctrl led.Controller) {

}
