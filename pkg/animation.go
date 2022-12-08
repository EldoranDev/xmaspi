package pkg

import (
	"time"
)

type Animation interface {
	FrameDuration() time.Duration
	ApplyFrame(ctrl Controller)
}

type AnimationWithInitializer interface {
	Init(ctrl Controller)
}

type AnimationWithDescription interface {
	DisplayName() string
	Description() string
}
