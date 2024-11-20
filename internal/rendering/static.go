package rendering

import (
	"github.com/EldoranDev/xmaspi/v3/internal/led"
	"time"
)

const StaticRenderer = "static"

var _ Renderer = (*static)(nil)

type static struct {
}

func (*static) FrameDuration() time.Duration {
	return time.Hour * 99
}

func (*static) ApplyFrame(ctrl led.Controller, color *led.Color) {
	ctrl.Fill(color)
}

func (*static) GetIdentifier() string {
	return StaticRenderer
}
