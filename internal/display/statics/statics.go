package statics

import (
	"errors"
	xmaspi "github.com/EldoranDev/xmaspi/v2/pkg"
)

var statics = map[string]func() xmaspi.Static{
	"BlueWhite": func() xmaspi.Static { return &BlueWhite{} },
}

func GetStatic(name string) (xmaspi.Static, error) {
	static, ok := statics[name]

	if !ok {
		return nil, errors.New("static does not exist")
	}

	return static(), nil
}

func GetAll() map[string]xmaspi.Static {
	stcs := make(map[string]xmaspi.Static)

	for s, static := range statics {
		stcs[s] = static()
	}

	return stcs
}
