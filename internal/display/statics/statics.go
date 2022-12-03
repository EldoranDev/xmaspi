package statics

import (
	"errors"
	xmaspi "github.com/EldoranDev/xmaspi/v2/pkg"
)

var statics = map[string]func() xmaspi.Static{
	"BlueWhite": func() xmaspi.Static { return &BlueWhite{} },
}

func GetStatic(name string) (xmaspi.Static, error) {
	anim, ok := statics[name]

	if !ok {
		return nil, errors.New("static does not exist")
	}

	return anim(), nil
}
