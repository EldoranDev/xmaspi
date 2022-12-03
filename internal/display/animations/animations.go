package animations

import (
	"errors"
	xmaspi "github.com/EldoranDev/xmaspi/v2/pkg"
)

var animations = map[string]func() xmaspi.Animation{
	"Blinker": func() xmaspi.Animation { return &Blinker{} },
	"Stars":   func() xmaspi.Animation { return &Stars{Color: xmaspi.Color{R: 255, G: 0, B: 0}} },
	"UpDown": func() xmaspi.Animation {
		return &UpDown{
			Top:    xmaspi.Color{R: 255, G: 0, B: 0},
			Bottom: xmaspi.Color{R: 0, G: 0, B: 255},
		}
	},
	"LeftRight": func() xmaspi.Animation {
		return &LeftRight{
			Left:  xmaspi.Color{R: 255, G: 0, B: 0},
			Right: xmaspi.Color{R: 0, G: 0, B: 0},
		}
	},
	"Rainbow": func() xmaspi.Animation {
		return &Rainbow{}
	},
}

func GetAnimation(name string) (xmaspi.Animation, error) {
	anim, ok := animations[name]

	if !ok {
		return nil, errors.New("animation does not exist")
	}

	return anim(), nil
}
