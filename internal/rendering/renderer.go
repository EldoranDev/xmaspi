package rendering

import (
	"github.com/EldoranDev/xmaspi/v3/internal/led"
	"time"
)

type Renderer interface {
	FrameDuration() time.Duration
	ApplyFrame(ctrl led.Controller)
}

type RendererWithInitializer interface {
	Init(ctrl led.Controller)
}

type RendererWithDescription interface {
	DisplayName() string
	Description() string
}
