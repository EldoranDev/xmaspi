package rendering

import (
	"github.com/EldoranDev/xmaspi/v3/internal/led"
	"time"
)

type Renderer interface {
	FrameDuration() time.Duration
	ApplyFrame(ctrl led.Controller, color *led.Color)
	GetIdentifier() string
}

type RendererWithInitializer interface {
	Init(ctrl led.Controller)
}
