package led

import (
	"encoding/json"
	"os"
)

type Settings struct {
	MaxBrightness uint8 `json:"max-brightness"`
	DataPin       uint8 `json:"data-pin"`
	Leds          []Led `json:"leds"`
}

func LoadSettings(path string) *Settings {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	var settings Settings

	if err = json.Unmarshal(data, &settings); err != nil {
		panic(err)
	}

	return &settings
}
